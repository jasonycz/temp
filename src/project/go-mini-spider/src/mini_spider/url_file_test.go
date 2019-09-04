/* url_file_test.go - test for url_file.go  */
/*
modification history
--------------------
2015/5/17, by Zhang Miao, modify
*/
/*
DESCRIPTION
*/
package mini_spider

import (
	"testing"
)

// test for filePathGen()
func TestFilePathGen(t *testing.T) {
	rootPath := "/var/mini-spider"
	url := "ab"

	filePath := filePathGen(url, rootPath)
	if filePath != "/var/mini-spider/ab" {
		t.Errorf("err in filePathGen(), filePath=%s", filePath)
	}
}

// test for urlFileSave
func TestUrlFileSave(t *testing.T) {
	rootPath := "./crawled"
	url := "www.baidu.com"
	data := []byte("this is a test")

	err := urlFileSave(rootPath, url, data)
	if err != nil {
		t.Errorf("err in urlFileSave(%s, %s):%s", rootPath, url, err.Error())
	}
}
