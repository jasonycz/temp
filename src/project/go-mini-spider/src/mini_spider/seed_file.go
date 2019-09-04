/* seed_file.go - load data from seed file	*/
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
	"encoding/json"
	"fmt"
	"io/ioutil"
)

/*
load data from seed file

Params:
	filePath: full path of seed file

Returns:
	(urls, error)
*/
func SeedFileLoad(filePath string) ([]string, error) {
	// read all data from file
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("ioutil.ReadFile(%s):%s", filePath, err.Error())
	}

	// json decode
	var seeds []string
	err = json.Unmarshal(data, &seeds)
	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal():%s", err.Error())
	}

	return seeds, nil
}
