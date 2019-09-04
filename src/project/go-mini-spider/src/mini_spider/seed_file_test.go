/* seed_file_test.go - test for seed_file.go  */
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
func TestSeedFileLoad(t *testing.T) {
	filePath := "./test_data/seed_file/url.data"
	seeds, err := SeedFileLoad(filePath)
	if err != nil {
		t.Errorf("err in seedFileLoad(%s):%s", filePath, err.Error())
	}

	if len(seeds) != 2 {
		t.Errorf("len(seeds) should be 2, now it's %d", len(seeds))
	}

	if seeds[0] != "http://www.baidu.com" || seeds[1] != "http://news.baidu.com" {
		t.Errorf("err in seeds, %s", seeds)
	}
}
