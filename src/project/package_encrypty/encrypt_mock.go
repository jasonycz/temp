package main

import (
	"bytes"
	"compress/gzip"
	"crypto"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/binary"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
	"math/big"
	"os"
	"path/filepath"
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
var privateKeyForSign = []byte(`
-----BEGIN RSA PRIVATE KEY-----
MIICXAIBAAKBgQDZuy3GEbahJc292fsyvrGneTJKQnzpdhNsJfDS5csb0MtmW+4J
EvBH5wCZK5j4+nrRfKBF7JuTHe0nSWOZWNxgLU87pwCxozXSNrsiiOjsV+3KwYfd
z5QlvvyCfvmllGObPqL7dWR92V2UYEWMSneBHtwDhCBCzmhAoOxZVsAq2wIDAQAB
AoGAe6Xi7tLsZkOzvZdnIoBNH5CiUK3Fhke56/b39qztVZSTjkOywJByyEKVgmbR
7sNAsGsYEd221CtwnJj+6P0f2Kn2mAXIeLtp76xTZQpMXpyNts+TTHn0AR0t1CPb
YO+UPxMH2MAe0P8LVoitN8nCTZLRwxO5GI1JXOEuyinoPXECQQD7O9Y0nE5Axp5B
Ya6YbpivfxLDPc8AAWhpjeVLBdaGsGIvBFk+w1MPmltcmNujvrH7TzCo1V088pau
jEx663xlAkEA3dyhcAXHflCWLha8oeApeSl1KvFsK2V8dlv9+GcOxLLOQiDUUNpc
NBYjQ8hrp33drAyqif0gtLuME7lSjz52PwJAclVD4YN0tVPud3C95U8Hm25oGoRJ
bQY2/zp9LAxHGdqxhBYli0OWzFYB1/0Uj8tb8jOSwHMBUJhR7LOapqH+ZQJBANbl
GL2/605vUjgML+P/k0xXmrTgqWyTXW3rhuEvOX/OzbgBwoUaLIsLEldabuQKJ0e5
8Qm9VAubSX/lVpGG59MCQEatbgT7lje91d7dGDLM9jVjUqRs35kzPCEQ78gbM5XF
X0/1u9UnH4N/KFTTUPtZ8Y0dlyJdRqjQ9tF37CRTCYY=
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
type SmappFile struct {
	header Header
	index  Index
	data   Data
}
type EncryptFileInfo struct {
	Md5  string
	Sign string
	Size int
}

func gzipEncode(in []byte) ([]byte, error) {
	var (
		buffer bytes.Buffer
		out    []byte
		err    error
	)
	writer, err := gzip.NewWriterLevel(&buffer, 9) //php: gzencode($json_data,9)
	if err != nil {
		return out, err
	}
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
func getFileData(path string, root string, pSmappFile *SmappFile) ([]byte, error) {
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
		pSmappFile.index = append(pSmappFile.index, FileInfo{
			Offset:     offset,
			Size:       uint32(f.Size()),
			PathLength: uint32(len([]byte(path[len(root):]))),
			Path:       path[len(root):],
		})
		offset += uint32(f.Size())
		buffer.Write(content)
		fileNum++
		return nil
	})
	if err != nil {
		return nil, err
	}
	// fmt.Printf("pSmappFile.index:\n%+v\n", pSmappFile.index)
	pSmappFile.header.DataLength = offset
	pSmappFile.header.FileNum = fileNum
	pSmappFile.data, err = gzipEncode(buffer.Bytes())
	// fmt.Printf("gzip data:\n%+v\n", hex.EncodeToString(pSmappFile.Data))
	if err != nil {
		return nil, errors.New("gzip encode error:" + err.Error())
	}
	return pSmappFile.data, nil
}
func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}
func aesEncrypt(origData, key []byte, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	origData = PKCS7Padding(origData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, iv)
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}
func createAesKeyIv(pSmappFile *SmappFile) error {
	key := md5.Sum([]byte(fmt.Sprint(pSmappFile.header.VersionCode) + fmt.Sprint(pSmappFile.header.FileNum)))
	pSmappFile.header.AESKey = key[:]
	iv := md5.Sum([]byte(fmt.Sprint(pSmappFile.header.Appid)))
	pSmappFile.header.AESIV = iv[:]
	fmt.Printf("AESKey Hex: 0x%s\n", hex.EncodeToString(pSmappFile.header.AESKey))
	fmt.Printf("AESIV Hex: 0x%s\n", hex.EncodeToString(pSmappFile.header.AESIV))
	fmt.Printf("versionCode:\n%+v\n", pSmappFile.header.VersionCode)
	fmt.Printf("FileNum:\n%+v\n", pSmappFile.header.FileNum)
	return nil
}
func getIndex(pSmappFile *SmappFile) ([]byte, error) {
	// cipher := "AES-128-CBC"
	buf := new(bytes.Buffer)
	byteOrder := binary.LittleEndian
	// 采用小端模式
	for _, f := range pSmappFile.index {
		binary.Write(buf, byteOrder, f.Offset)
		binary.Write(buf, byteOrder, f.Size)
		binary.Write(buf, byteOrder, f.PathLength)
		buf.Write([]byte(f.Path))
	}
	// 生成aes key和iv
	createAesKeyIv(pSmappFile)
	// 通过AES算法CBC模式加密
	aesEncryptData, err := aesEncrypt(buf.Bytes(), pSmappFile.header.AESKey, pSmappFile.header.AESIV)
	if err != nil {
		return nil, errors.New("AesEncrypt error:" + err.Error())
	}
	// fmt.Printf("aesEncryptData:\n%+v\n", hex.EncodeToString(aesEncryptData))
	pSmappFile.header.IndexLength = uint32(len(aesEncryptData))
	return aesEncryptData, nil
}
func PrivateEncrypt(priv *rsa.PrivateKey, data []byte) (enc []byte, err error) {
	k := (priv.N.BitLen() + 7) / 8
	tLen := len(data)
	// // rfc2313, section 8:
	// // The length of the data D shall not be more than k-11 octets
	if tLen > k-11 {
		return nil, errors.New("input size too large")
	}
	em := make([]byte, k)
	em[1] = 1
	for i := 2; i < k-tLen-1; i++ {
		em[i] = 0xff
	}
	copy(em[k-tLen:k], data)
	c := new(big.Int).SetBytes(em)
	if c.Cmp(priv.N) > 0 {
		return nil, errors.New("encryption error")
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
func rsaEncrypt(origData []byte) ([]byte, error) {
	// 加密pem格式的私钥
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, errors.New("private key error")
	}
	// 解析私钥
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	// rsaKeyIv, err := PrivateEncrypt(priv, origData)

	// rsa 签名
	rsaKeyIv, err := rsa.SignPKCS1v15(nil, priv, crypto.Hash(0), origData)

	fmt.Printf("\n==========\n%+v\n===========\n", rsaKeyIv)
	return rsaKeyIv, nil
}
func getHeader(pSmappFile *SmappFile) ([]byte, error) {
	buf := new(bytes.Buffer)
	// 采用小端模式
	byteOrder := binary.LittleEndian
	binary.Write(buf, byteOrder, pSmappFile.header.MagicNumber)
	binary.Write(buf, byteOrder, pSmappFile.header.VersionCode)
	binary.Write(buf, byteOrder, pSmappFile.header.FileNum)
	binary.Write(buf, byteOrder, pSmappFile.header.IndexLength)
	binary.Write(buf, byteOrder, pSmappFile.header.DataLength)
	binary.Write(buf, byteOrder, pSmappFile.header.Appid)
	// rsa 加密 key iv
	aesKeyIv := new(bytes.Buffer)
	aesKeyIv.Write(pSmappFile.header.AESKey)
	aesKeyIv.Write(pSmappFile.header.AESIV)
	// fmt.Printf("aesKeyIv:\n%+v\n", hex.EncodeToString(aesKeyIv.Bytes()))
	fmt.Printf("aesKeyIv:\n%+v\n", aesKeyIv.Bytes())
	rsaKeyIv, err := rsaEncrypt(aesKeyIv.Bytes())
	if err != nil {
		return nil, errors.New("rsa Kye Iv error :\n" + err.Error())
	}
	// fmt.Printf("Rsa aesKeyIv:\n%+v\n", hex.EncodeToString(rsaKeyIv))
	binary.Write(buf, byteOrder, rsaKeyIv)
	// fmt.Printf("Header:\n%+v\n", pSmappFile.header)
	return buf.Bytes(), nil
}
func createSmapp(data []byte, fileName string) error {
	// 写入文件
	file, err := os.Create(fileName)
	if err != nil {
		return errors.New("create file error :\n" + err.Error())
	}
	// 写入byte的slice数据
	file.Write(data)
	file.Close()
	return nil
}

// Sign本质是使用rsa私钥签名
func Sign(origData []byte) (string, error) {
	// 加密pem格式的私钥
	block, _ := pem.Decode(privateKeyForSign)
	// fmt.Printf("block:\n%+v\n", block)
	if block == nil {
		return "", errors.New("private key error")
	}
	// 解析私钥
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return "", err
	}
	// 类比：php 中的openssl_private_encrypt($data, $signature, $pk, OPENSSL_PKCS1_PADDING);
	// 使用PrivateEncrypt Badcase: origData= []byte(ed925b76ad1874a51683cd9c8bac81da)
	// sign, err := PrivateEncrypt(priv, origData)
	// 参考: https://stackoverflow.com/questions/18011708/encrypt-message-with-rsa-private-key-as-in-openssls-rsa-private-encrypt
	sign, err := rsa.SignPKCS1v15(nil, priv, crypto.Hash(0), origData)
	// fmt.Printf("sign:\n%+v\n", hex.EncodeToString(sign))
	if err != nil {
		return "", errors.New("rsa sign error " + err.Error())
	}
	return base64.URLEncoding.EncodeToString(sign), nil
}
func Encrypt(versionCode int, appID int64, input string, output string) (EncryptFileInfo, error) {
	if versionCode <= 0 {
		return EncryptFileInfo{}, errors.New("versionCode is invalid")
	}
	if appID <= 0 {
		return EncryptFileInfo{}, errors.New("appID is invalid")
	}
	// 0.初始化变量
	pSmappFile := &SmappFile{
		header: Header{
			MagicNumber: 3172468484, //0xBD180704
			VersionCode: uint32(versionCode),
			Appid:       uint64(appID),
		},
	}
	// 1.获取所有的文件的压缩后二进制内容 & 内容大小 & 路径 & 路径大小 (加密文件第三段内容)
	data, err := getFileData(input, input, pSmappFile)
	if err != nil {
		return EncryptFileInfo{}, errors.New("getFileInfo error:\n" + err.Error())
	}
	// 2.获取文件的 Index (加密文件的第二段)
	index, err := getIndex(pSmappFile)
	if err != nil {
		return EncryptFileInfo{}, errors.New("getIndex error:\n" + err.Error())
	}
	// 3.获取文件的Header (加密文件的第一段)
	header, err := getHeader(pSmappFile)
	if err != nil {
		return EncryptFileInfo{}, errors.New(" getHeader error:\n" + err.Error())
	}
	// 4.合并第一段第二段第三段并生成 smapp文件
	buf := new(bytes.Buffer)
	buf.Write(header)
	buf.Write(index)
	buf.Write(data)
	createSmapp(buf.Bytes(), output)
	// 5. 计算签名、计算 md5、计算大小
	md5 := md5.Sum(buf.Bytes())
	// 注意此处sign 的对象不是md5[:] 而是 []byte(hex.EncodeToString(md5[:]))
	sign, err := Sign([]byte(hex.EncodeToString(md5[:])))
	if err != nil {
		return EncryptFileInfo{}, errors.New(" get sign error:\n" + err.Error())
	}
	return EncryptFileInfo{
		Md5:  hex.EncodeToString(md5[:]),
		Sign: sign,
		Size: buf.Len(),
	}, nil
}

func main() {
	versionCode := 413
	appID := int64(14539182)
	input := "/Users/yangchengzhi/yangchengzhi/test/tt/t/"
	// input := "/Users/yangchengzhi/yangchengzhi/test/zip/"
	output := "./smapp/test_mock.smapp"

	e := EncryptFileInfo{}

	e, err := Encrypt(versionCode, appID, input, output)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", e)
}
