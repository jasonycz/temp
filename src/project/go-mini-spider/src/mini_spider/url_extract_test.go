/* url_extract_test.go - test for url_extract.go  */
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

// test for urlExtract()
func TestUrlExtract(t *testing.T) {
	s := []byte(`<p>Links:</p><ul><li><a href="foo">Foo</a><li><a href="/bar/baz">BarBaz</a></ul>`)

	links, err := urlExtract(s, "")
	if err != nil {
		t.Errorf("err in urlExtract():%s", err.Error())
		return
	}

	if len(links) != 2 {
		t.Errorf("len(links) should be 2, now it's %d", len(links))
		return
	}

	if links[0] != "foo" || links[1] != "/bar/baz" {
		t.Errorf("links:%s", links)
	}
}
