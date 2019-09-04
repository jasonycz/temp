/* url_table_test.go - test for url_table.go  */
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

func TestUrlTable_case_1(t *testing.T) {
	// create table
	ut := NewUrlTable()

	// add to table
	ut.Add("www.baidu.com")

	// check whether exist
	if !ut.Exist("www.baidu.com") {
		t.Errorf("www.baidu.com should exist")
	}

	if ut.Exist("news.baidu.com") {
		t.Errorf("news.baidu.com should not exist")
	}
}
