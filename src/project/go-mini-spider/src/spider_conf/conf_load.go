/* conf_load.go - load config for mini-spider */
/*
modification history
--------------------
2015/5/17, by Zhang Miao, create
*/
/*
DESCRIPTION
*/
package spider_conf

import (
	"fmt"
	"regexp"
)

import (
	"code.google.com/p/gcfg"
)

type ConfBasic struct {
	UrlListFile     string
	OutputDirectory string

	MaxDepth      int
	CrawlInterval int
	CrawlTimeout  int
	TargetUrl     string

	ThreadCount int
}

type SpiderConfig struct {
	Spider ConfBasic
}

/*
ConfigLoad - load config for mini-spider

Params:
    - filePath: path of config file

Returns:
    (SpiderConfig, error)
*/
func ConfigLoad(filePath string) (SpiderConfig, error) {
	var cfg SpiderConfig
	var err error

	// read config from file
	err = gcfg.ReadFileInto(&cfg, filePath)
	if err != nil {
		return cfg, fmt.Errorf("gcfg.ReadFileInto():%s", err.Error())
	}

	// check Target Url
	_, err = regexp.Compile(cfg.Spider.TargetUrl)
	if err != nil {
		return cfg, fmt.Errorf("regexp.Compile(TargetUrl):%s", err.Error())
	}

	return cfg, nil
}
