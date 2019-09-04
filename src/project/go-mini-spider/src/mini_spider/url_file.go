/* url_file.go - for saving crawled data to file	*/
/*
modification history
--------------------
2015/5/17, by Zhang Miao, create
*/
/*
DESCRIPTION
*/
package mini_spider

import (
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"path"
)

/*
generate file path for given url

Params:
	url: url to crawl
	rootPath: root path for saving file

Returns:
	file path
*/
func filePathGen(urlStr, rootPath string) string {
	filePath := url.QueryEscape(urlStr)
	filePath = path.Join(rootPath, filePath)
	return filePath
}

/*
save data for given url to file

Params:
	url: url to crawl
	rootPath: root path for saving file
	data: data to save
*/
func urlFileSave(rootPath string, url string, data []byte) error {
	// create root dir, if not exist
	if _, err := os.Stat(rootPath); os.IsNotExist(err) {
		if os.MkdirAll(rootPath, 0644) != nil {
			return fmt.Errorf("os.MkdirAll(%s):%s", rootPath, err.Error())
		}
	}

	// save data to disk
	// generate full file path
	filePath := filePathGen(url, rootPath)

	// save to file
	err := ioutil.WriteFile(filePath, data, 0777)
	if err != nil {
		return fmt.Errorf("ioutil.WriteFile(%s):%s", filePath, err.Error())
	}

	return nil
}
