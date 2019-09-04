/* conf_load_test.go - test for conf_load.go  */
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
	"testing"
)

// normal case
func TestConfigLoad(t *testing.T) {
	_, err := ConfigLoad("./test_data/mini_spider.conf")
	if err != nil {
		t.Errorf("err in ConfigLoad():%s", err.Error())
		return
	}
}
