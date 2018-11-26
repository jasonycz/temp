package main

import (
	"bytes"
	"compress/gzip"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rsa"
	"crypto/x509"
	"encoding/binary"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
	"math/big"
	"os"
	"path/filepath"
	"strings"
)

var privateKey = []byte(`
-----BEGIN RSA PRIVATE KEY-----
MIHyAgEAAjEAys6khY5+4YMbVuNRdXlo9TlRCvS+Y6omJ/i4kL7v/YH02bbtIQYg
16LvRAhbjH+/AgMBAAECMQCSFyLqVkhN2GyUPXGLZubVZvFbLAwNrR1Ax91zjvq/
DXUWeK11uqhXHrZUGYGFupECGQDsKg0ff/P9T4lkswderTrRMUjoyU7AwRkCGQDb
11sISDIlbZYFLGUBg/4IW51v5vigqpcCGCokqujXHYQzoFH+/K6xfRB9mgQogEw3
aQIYF52KqdkXie27r9KVBxSmtT8iKmMlUh0BAhgkT1vwPJqByItUaAPH/EyBgXdv
z2jjGE0=
-----END RSA PRIVATE KEY-----
`)

type Header struct {
	MagicNumber uint32
	VersionCode uint32
	FileNum     uint32
	IndexLength uint32
	DataLength  uint32
	Appid       uint64
	Cipher      []byte //48 Bytes
	AESKey      []byte
	AESIV       []byte
}

type FileInfo struct {
	Offset     uint32 //int64
	Size       uint32 //int64
	PathLength uint32 //int
	Path       string
}
type Index []FileInfo
type Data []byte

type smappFile struct {
	header Header
	index  Index
	data   Data
}

type Files []FileInfo

type StructureDeflateData struct {
	MagicNumber uint32 // magic num
	Appid       uint64 // appid
	VersionCode uint32 // version code
	Size        uint32 // all file size
	FileNum     uint32 // all file num
	Data        []byte // all really data
	Files       Files  // all files
	IndexData   []byte // the second index data
	IndexLength uint32 // the index length
	AESKey      []byte // aes key
	AESIV       []byte // aes iv
	Header      []byte // header
}

func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func AesEncrypt(origData, key []byte, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	origData = PKCS7Padding(origData, blockSize)
	// origData = ZeroPadding(origData, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, iv)
	crypted := make([]byte, len(origData))
	// 根据CryptBlocks方法的说明，如下方式初始化crypted也可以
	// crypted := origData
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

func GzipEncode(in []byte) ([]byte, error) {
	var (
		buffer bytes.Buffer
		out    []byte
		err    error
	)
	writer, _ := gzip.NewWriterLevel(&buffer, 9) //php: gzencode($json_data,9)

	_, err = writer.Write(in)
	if err != nil {
		writer.Close()
		return out, err
	}
	err = writer.Close()
	if err != nil {
		return out, err
	}

	return buffer.Bytes(), nil
}

func getFilelist(path string, root string, pStructureDeflateData *StructureDeflateData) error {
	offset := uint32(0)
	fileNum := uint32(0)
	var buffer bytes.Buffer
	err := filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			return nil
		}
		content, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}
		pStructureDeflateData.Files = append(pStructureDeflateData.Files, FileInfo{
			Offset:     offset,
			Size:       uint32(f.Size()),
			PathLength: uint32((strings.Count(path[len(root):], "") - 1)),
			Path:       path[len(root):],
		})
		offset += uint32(f.Size())
		buffer.Write(content)
		fileNum++
		return nil
	})

	if err != nil {
		return err
	}
	pStructureDeflateData.Size = offset
	pStructureDeflateData.Data, err = GzipEncode(buffer.Bytes())
	fmt.Printf("gzip data:\n%+v\n", hex.EncodeToString(pStructureDeflateData.Data))

	pStructureDeflateData.FileNum = fileNum
	if err != nil {
		return errors.New("gzip encode error:" + err.Error())
	}
	return nil
}

// 根据已有数据hash生成aes的key和iv,hash返回16字节二进制
func createAesKeyIv(pStructureDeflateData *StructureDeflateData) error {
	key := md5.Sum([]byte(fmt.Sprint(pStructureDeflateData.VersionCode) + fmt.Sprint(pStructureDeflateData.FileNum)))
	pStructureDeflateData.AESKey = key[:]
	iv := md5.Sum([]byte(fmt.Sprint(pStructureDeflateData.Appid)))
	pStructureDeflateData.AESIV = iv[:]
	fmt.Printf("aes key:\n%+v\n", hex.EncodeToString(pStructureDeflateData.AESKey))
	fmt.Printf("aes iv:\n%+v\n", hex.EncodeToString(pStructureDeflateData.AESIV))
	return nil
}

// ($thirdContent['indexData'], $version_code, count($files), $appId);
func structureIndex(pStructureDeflateData *StructureDeflateData) error {
	// cipher := "AES-128-CBC"
	buf := new(bytes.Buffer)
	byteOrder := binary.LittleEndian

	// 采用小端模式
	for _, f := range pStructureDeflateData.Files {
		binary.Write(buf, byteOrder, f.Offset)
		binary.Write(buf, byteOrder, f.Size)
		binary.Write(buf, byteOrder, f.PathLength)
		buf.Write([]byte(f.Path))
	}

	// 生成aes key和iv
	createAesKeyIv(pStructureDeflateData)

	// 通过AES算法CBC模式加密
	aesEncryptData, err := AesEncrypt(buf.Bytes(), pStructureDeflateData.AESKey, pStructureDeflateData.AESIV)
	if err != nil {
		return errors.New("AesEncrypt error:" + err.Error())
	}
	fmt.Printf("aesEncryptData:\n%+v\n", hex.EncodeToString(aesEncryptData))
	pStructureDeflateData.IndexData = aesEncryptData

	pStructureDeflateData.IndexLength = uint32(len(aesEncryptData))
	return nil
}

var (
	ErrInputSize  = errors.New("input size too large")
	ErrEncryption = errors.New("encryption error")
)

func PrivateEncrypt(priv *rsa.PrivateKey, data []byte) (enc []byte, err error) {

	k := (priv.N.BitLen() + 7) / 8
	tLen := len(data)
	// rfc2313, section 8:
	// The length of the data D shall not be more than k-11 octets
	if tLen > k-11 {
		err = ErrInputSize
		return
	}
	em := make([]byte, k)
	em[1] = 1
	for i := 2; i < k-tLen-1; i++ {
		em[i] = 0xff
	}
	copy(em[k-tLen:k], data)
	c := new(big.Int).SetBytes(em)
	if c.Cmp(priv.N) > 0 {
		err = ErrEncryption
		return
	}
	var m *big.Int
	var ir *big.Int
	if priv.Precomputed.Dp == nil {
		m = new(big.Int).Exp(c, priv.D, priv.N)
	} else {
		// We have the precalculated values needed for the CRT.
		m = new(big.Int).Exp(c, priv.Precomputed.Dp, priv.Primes[0])
		m2 := new(big.Int).Exp(c, priv.Precomputed.Dq, priv.Primes[1])
		m.Sub(m, m2)
		if m.Sign() < 0 {
			m.Add(m, priv.Primes[0])
		}
		m.Mul(m, priv.Precomputed.Qinv)
		m.Mod(m, priv.Primes[0])
		m.Mul(m, priv.Primes[1])
		m.Add(m, m2)

		for i, values := range priv.Precomputed.CRTValues {
			prime := priv.Primes[2+i]
			m2.Exp(c, values.Exp, prime)
			m2.Sub(m2, m)
			m2.Mul(m2, values.Coeff)
			m2.Mod(m2, prime)
			if m2.Sign() < 0 {
				m2.Add(m2, prime)
			}
			m2.Mul(m2, values.R)
			m.Add(m, m2)
		}
	}

	if ir != nil {
		// Unblind.
		m.Mul(m, ir)
		m.Mod(m, priv.N)
	}
	enc = m.Bytes()
	return
}

func RsaEncrypt(origData []byte) ([]byte, error) {
	//加密pem格式的私钥
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, errors.New("private key error")
	}
	// 解析私钥
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	rsaKeyIv, err := PrivateEncrypt(priv, origData)
	if err != nil {
		return nil, errors.New("rsa sign error " + err.Error())
	}
	return rsaKeyIv, nil

}

func structureHeader(pStructureDeflateData *StructureDeflateData) error {
	buf := new(bytes.Buffer)
	byteOrder := binary.LittleEndian

	// 采用小端模式
	binary.Write(buf, byteOrder, pStructureDeflateData.MagicNumber)
	binary.Write(buf, byteOrder, pStructureDeflateData.VersionCode)
	binary.Write(buf, byteOrder, pStructureDeflateData.FileNum)
	binary.Write(buf, byteOrder, pStructureDeflateData.IndexLength)
	binary.Write(buf, byteOrder, pStructureDeflateData.Size)
	binary.Write(buf, byteOrder, pStructureDeflateData.Appid)

	// rsa 加密 key iv
	aesKeyIv := new(bytes.Buffer)
	aesKeyIv.Write(pStructureDeflateData.AESKey)
	aesKeyIv.Write(pStructureDeflateData.AESIV)
	fmt.Printf("aesKeyIv:\n%+v\n", hex.EncodeToString(aesKeyIv.Bytes()))

	rsaKeyIv, err := RsaEncrypt(aesKeyIv.Bytes())
	if err != nil {
		return errors.New("rsa Kye Iv error :\n" + err.Error())
	}
	// fmt.Printf("aesEncryptData:\n%+v\n", hex.EncodeToString(aesEncryptData))
	fmt.Printf("Rsa aesKeyIv:\n%+v\n", hex.EncodeToString(rsaKeyIv))
	binary.Write(buf, byteOrder, rsaKeyIv)

	pStructureDeflateData.Header = buf.Bytes()

	fmt.Printf("Rsa aesKeyIv:\n%+v\n", hex.EncodeToString(buf.Bytes()))
	return nil
}

func createSmapp(data []byte, fileName string) error {
	//写入文件
	file, err := os.Create(fileName)
	if err != nil {
		return errors.New("create file error :\n" + err.Error())
	}
	//写入byte的slice数据
	file.Write(data)
	file.Close()
	return nil
}

func encrypt(versionCode int, appId int64, input string, output string) error {
	if versionCode <= 0 {
		return errors.New("versionCode is invalid")
	}
	if appId <= 0 {
		return errors.New("appId is invalid")
	}

	//1. 获取所有的包 二进制内容 & 内容大小 & 路径 & 路径大小
	pStructureDeflateData := &StructureDeflateData{
		MagicNumber: 3172468484, //0xBD180704
		VersionCode: uint32(versionCode),
		Appid:       uint64(appId),
	}
	err := getFilelist(input, input, pStructureDeflateData)
	if err != nil {
		return errors.New("getFilelist error:" + err.Error())
	}
	// fmt.Printf("======%+v===pStructureDeflateData\n", pStructureDeflateData)

	// 2. $this->_structureIndex($thirdContent['indexData'], $version_code, count($files), $appId);
	err = structureIndex(pStructureDeflateData)
	if err != nil {
		return errors.New("structureIndex error" + err.Error())
	}

	// 3. _structureHeader($version_code, count($files), $secondContent['size'], $thirdContent['size'], $appId);
	err = structureHeader(pStructureDeflateData)
	if err != nil {
		return errors.New("structureHeader error " + err.Error())
	}

	// 4. 合并 第一段 第二段 第三段
	buf := new(bytes.Buffer)
	buf.Write(pStructureDeflateData.Header)
	buf.Write(pStructureDeflateData.IndexData)
	buf.Write(pStructureDeflateData.Data)

	// 5. 生成 smapp 文件
	createSmapp(buf.Bytes(), output)

	return nil
}

func main() {
	versionCode := 1
	appId := int64(123)
	input := "/Users/yangchengzhi/yangchengzhi/test/zip"
	output := "./smapp/test.smapp"
	err := encrypt(versionCode, appId, input, output)
	if err != nil {
		panic(err)
	}
	fmt.Print("encrypt ok\n")
}
