package main

import (
	"bytes"
	"compress/gzip"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rsa"
	"crypto/x509"
	"encoding/binary"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"io"
	"io/ioutil"
	"math/big"
	"os"
	"path/filepath"
)

const (
	_PublicKey = "-----BEGIN PUBLIC KEY-----\n" +
		"MEwwDQYJKoZIhvcNAQEBBQADOwAwOAIxAMrOpIWOfuGDG1bjUXV5aPU5UQr0vmOq\n" +
		"Jif4uJC+7/2B9Nm27SEGINei70QIW4x/vwIDAQAB\n" +
		"-----END PUBLIC KEY-----"
)

type fileInfo struct {
	Offset     uint32
	Size       uint32
	PathLength uint32
	Path       string
}
type smappFile struct {
	f           *os.File
	output      string
	MagicNumber uint32
	VersionCode uint32
	FileNumber  uint32
	IndexLength uint32
	DataLength  uint32
	Appid       uint64
	AESKey      []byte
	AESIV       []byte
	Files       []*fileInfo
}

func openSmappFile(path string, output string) (*smappFile, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	return &smappFile{
		f:      f,
		output: output,
	}, nil
}

func getUInt32(f io.Reader, buf []byte) (uint32, error) {
	if buf == nil {
		buf = make([]byte, 4)
	} else if len(buf) != 4 {
		return 0, fmt.Errorf("buf's length should be 4 bytes")
	}
	n, err := f.Read(buf)
	if err != nil || n != 4 {
		panic(err)
	}
	return binary.LittleEndian.Uint32(buf), nil
}

func getUInt64(f io.Reader, buf []byte) (uint64, error) {
	if buf == nil {
		buf = make([]byte, 8)
	} else if len(buf) != 8 {
		return 0, fmt.Errorf("buf's length should be 4 bytes")
	}
	n, err := f.Read(buf)
	if err != nil || n != 8 {
		panic(err)
	}
	return binary.LittleEndian.Uint64(buf), nil
}

func getAESKeyIV(f *os.File) ([]byte, []byte) {
	buf := make([]byte, 48)
	n, err := f.Read(buf)
	if err != nil || n != 48 {
		panic(err)
	}
	pubBlock, _ := pem.Decode([]byte(_PublicKey))
	if pubBlock == nil {
		panic("PublicKey error")
	}
	pki, err := x509.ParsePKIXPublicKey(pubBlock.Bytes)
	if err != nil {
		panic(err)
	}
	pk := pki.(*rsa.PublicKey)
	crypted := new(big.Int).SetBytes(buf)
	decrypted := crypted.Exp(crypted, big.NewInt(int64(pk.E)), pk.N)
	plaintext := decrypted.Bytes()

	key := plaintext[15:31]
	iv := plaintext[31:47]
	return key, iv
}

func (f *smappFile) UnpackStage1() {
	var err error
	fmt.Println("------------- Stage 1 -------------")
	buf := make([]byte, 4)

	// Magic Number
	f.MagicNumber, err = getUInt32(f.f, buf)
	if err != nil {
		panic(err)
	}
	fmt.Printf("MagicNumber: 0x%x\n", f.MagicNumber)

	// Version Code
	f.VersionCode, err = getUInt32(f.f, buf)
	if err != nil {
		panic(err)
	}
	fmt.Printf("VersionCode: %d\n", f.VersionCode)

	// File Number
	f.FileNumber, err = getUInt32(f.f, buf)
	if err != nil {
		panic(err)
	}
	fmt.Printf(" FileNumber: %d\n", f.FileNumber)

	// Index Length
	f.IndexLength, err = getUInt32(f.f, buf)
	if err != nil {
		panic(err)
	}
	fmt.Printf("IndexLength: %d\n", f.IndexLength)

	// Data Length
	f.DataLength, err = getUInt32(f.f, buf)
	if err != nil {
		panic(err)
	}
	fmt.Printf(" DataLength: %d\n", f.DataLength)

	// Appid
	f.Appid, err = getUInt64(f.f, nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf("      Appid: %d\n", f.Appid)

	f.AESKey, f.AESIV = getAESKeyIV(f.f)
	fmt.Printf(" AESKey Hex: 0x%s\n", hex.EncodeToString(f.AESKey))
	fmt.Printf("  AESIV Hex: 0x%s\n", hex.EncodeToString(f.AESIV))
}

func (f *smappFile) UnpackStage2() {
	fmt.Println("------------- Stage 2 -------------")
	key, err := aes.NewCipher(f.AESKey)
	if err != nil {
		panic(key)
	}
	decryptor := cipher.NewCBCDecrypter(key, f.AESIV)

	// indexDataBuf 是长度为 f.IndexLength 的字节数组
	indexDataBuf := make([]byte, f.IndexLength)
	indexData := make([]byte, f.IndexLength)
	// 将长度为f.IndexLength的 f 的数据读入到 indexDataBuf中
	_, err = f.f.Read(indexDataBuf)
	if err != nil {
		panic(err)
	}
	// CryptBlocks(目标数据, 加密后的数据)
	decryptor.CryptBlocks(indexData, indexDataBuf)

	// 将数据转换成一个 bytes 类型 的可读取的内容
	indexDataReader := bytes.NewReader(indexData)

	fileCount := 0
	for indexDataReader.Len() > 12 {
		offset, err := getUInt32(indexDataReader, nil)
		if err != nil {
			panic(err)
		}
		size, err := getUInt32(indexDataReader, nil)
		if err != nil {
			panic(err)
		}
		pathlen, err := getUInt32(indexDataReader, nil)
		if err != nil {
			panic(err)
		}
		if indexDataReader.Len() < int(pathlen) {
			break
		}
		path := make([]byte, pathlen)
		_, err = indexDataReader.Read(path)
		if err != nil {
			panic(err)
		}
		fileCount++
		fmt.Printf("File %d: %s, Offset=%d, Size=%d, UnreadLen: %d\n",
			fileCount, string(path), offset, size, indexDataReader.Len())
		f.Files = append(f.Files, &fileInfo{
			Offset:     offset,
			Size:       size,
			PathLength: pathlen,
			Path:       string(path),
		})
	}
}

func (f *smappFile) UnpackStage3() {
	fmt.Println("------------- Stage 3 -------------")
	fileContents, err := ioutil.ReadAll(f.f)
	if err != nil {
		panic(err)
	}
	fmt.Printf("fileContentHeader: %s\n", hex.EncodeToString(fileContents[:16]))
	fcReader, err := gzip.NewReader(bytes.NewReader(fileContents))
	if err != nil {
		panic(err)
	}
	allContents, err := ioutil.ReadAll(fcReader)
	if err != nil {
		panic(err)
	}

	for _, fileInfo := range f.Files {
		fpath := f.output + fileInfo.Path
		if _, err := os.Stat(filepath.Dir(fpath)); os.IsNotExist(err) {
			os.MkdirAll(filepath.Dir(fpath), 0755)
		}
		if err = ioutil.WriteFile(fpath, allContents[fileInfo.Offset:fileInfo.Offset+fileInfo.Size], 0644); err != nil {
			panic(err)
		}
		fmt.Printf("Wrote file: %s, offset=%d, size=%d\n", fpath, fileInfo.Offset, fileInfo.Size)
	}
}

func main() {
	if len(os.Args) < 3 {
		fmt.Printf("Usage: %s package.smapp output_dir\n\n", os.Args[0])
		fmt.Println("      Date: 2018-06-24")
		fmt.Println("    Author: chenxin<chenxin@baidu.com>")
		return
	}
	sf, err := openSmappFile(os.Args[1], os.Args[2])
	if err != nil {
		panic(err)
	}
	sf.UnpackStage1()
	sf.UnpackStage2()
	sf.UnpackStage3()
}
