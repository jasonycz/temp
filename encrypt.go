package main

import (
	// "flag"
	"bytes"
	"compress/gzip"
	"crypto/md5"
	"encoding/binary"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	// "strconv"
	// "encoding/hex"
	"encoding/base64"
	"log"
	"strings"

	"github.com/spacemonkeygo/openssl"
)

// type smappFile struct {
// 	f           *os.File
// 	output      string
// 	MagicNumber uint32
// 	VersionCode uint32
// 	FileNumber  uint32
// 	IndexLength uint32
// 	DataLength  uint32
// 	Appid       uint64
// 	AESKey      []byte
// 	AESIV       []byte
// 	Files       []*fileInfo
// }
// FILE_FLAG = 0xBD180704;

type FileInfo struct {
	// Content    []byte
	Offset     uint32 //int64
	Size       uint32 //int64
	PathLength uint32 //int
	Path       string
}
type Files []FileInfo

type StructureDeflateData struct {
	MagicNumber uint32
	Appid       uint64
	VersionCode uint32
	Size        uint32 //int64
	Data        []byte
	IndexData   Files
	FileNum     uint32 //int
	AESKey      [16]byte
	AESIV       [16]byte
}

type Crypter struct {
	key    []byte
	iv     []byte
	cipher *openssl.Cipher
}

func NewCrypter(key []byte, iv []byte) (*Crypter, error) {
	cipher, err := openssl.GetCipherByName("AES-128-CBC")
	if err != nil {
		return nil, err
	}

	return &Crypter{key, iv, cipher}, nil
}
func (c *Crypter) Encrypt(input []byte) ([]byte, error) {
	ctx, err := openssl.NewEncryptionCipherCtx(c.cipher, nil, c.key, c.iv)
	if err != nil {
		return nil, err
	}

	cipherbytes, err := ctx.EncryptUpdate(input)
	if err != nil {
		return nil, err
	}

	finalbytes, err := ctx.EncryptFinal()
	if err != nil {
		return nil, err
	}

	cipherbytes = append(cipherbytes, finalbytes...)
	return cipherbytes, nil
}

func GzipEncode(in []byte) ([]byte, error) {
	var (
		buffer bytes.Buffer
		out    []byte
		err    error
	)
	writer := gzip.NewWriter(&buffer)
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

func getFilelist(path string, pStructureDeflateData *StructureDeflateData) error {
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
		pStructureDeflateData.IndexData = append(pStructureDeflateData.IndexData, FileInfo{
			// Content:    content,
			Offset:     offset,
			Size:       uint32(f.Size()),
			PathLength: uint32((strings.Count(path, "") - 1)),
			Path:       path,
		})
		offset += uint32(f.Size())
		buffer.Write(content)
		fileNum++
		return nil
	})

	// fmt.Printf("======%+v===startOffset\n", startOffset)
	if err != nil {
		return err
	}
	pStructureDeflateData.Size = offset
	pStructureDeflateData.Data, err = GzipEncode(buffer.Bytes())
	pStructureDeflateData.FileNum = fileNum
	if err != nil {
		return errors.New("gzip encode error:" + err.Error())
	}
	return nil
}

// $version_code, $fileNum, $appId
// 根据已有数据hash生成aes的key和iv,hash返回16字节二进制
func createAesKeyIv(pStructureDeflateData *StructureDeflateData) error {
	pStructureDeflateData.AESKey = md5.Sum([]byte(fmt.Sprint(pStructureDeflateData.VersionCode) + fmt.Sprint(pStructureDeflateData.FileNum)))
	pStructureDeflateData.AESIV = md5.Sum([]byte(fmt.Sprint(pStructureDeflateData.Appid)))
	fmt.Printf("%x\n", pStructureDeflateData.AESIV)
	return nil
}

// ($thirdContent['indexData'], $version_code, count($files), $appId);
func structureIndex(pStructureDeflateData *StructureDeflateData) error {
	// cipher := "AES-128-CBC"
	buf := new(bytes.Buffer)
	byteOrder := binary.LittleEndian

	// 采用小端模式
	for _, f := range pStructureDeflateData.IndexData {
		binary.Write(buf, byteOrder, f.Offset)
		binary.Write(buf, byteOrder, f.Size)
		binary.Write(buf, byteOrder, f.PathLength)
		binary.Write(buf, byteOrder, f.Path)
	}

	// 生成aes key和iv
	createAesKeyIv(pStructureDeflateData)
	// fmt.Printf("======%+v===pStructureDeflateData\n", pStructureDeflateData.AESIV)
	// fmt.Printf("======%+v===pStructureDeflateData\n", pStructureDeflateData.AESKey)

	// 通过AES算法CBC模式加密
	// Initialize new crypter struct. Errors are ignored.
	crypter, _ := NewCrypter(pStructureDeflateData.AESKey, pStructureDeflateData.AESIV)
	// Lets encode plaintext using the same key and iv.
	// This will produce the very same result: "RanFyUZSP9u/HLZjyI5zXQ=="
	encoded, _ := crypter.Encrypt([]byte("hello world"))
	log.Println(base64.StdEncoding.EncodeToString(encoded))

	return nil
}

func encrypt(sourcePath string, versionCode int, appId int64, destination string) error {
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
	err := getFilelist(sourcePath, pStructureDeflateData)
	if err != nil {
		return errors.New("getFilelist error:" + err.Error())
	}
	// fmt.Printf("======%+v===pStructureDeflateData\n", pStructureDeflateData)

	// 2. $this->_structureIndex($thirdContent['indexData'], $version_code, count($files), $appId);
	err = structureIndex(pStructureDeflateData)
	if err != nil {
		return errors.New("structureIndex error")
	}

	return nil
}

func main() {
	sourcePath := "./zip"
	versionCode := 1
	appId := int64(123)
	destination := "./smapp/test.smapp"
	err := encrypt(sourcePath, versionCode, appId, destination)
	if err != nil {
		panic(err)
	}
	fmt.Print("encrypt ok")
}
