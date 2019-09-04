/* url_table.go - table for saving url that has been crawled    */
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
	"sync"
)

type UrlTable struct {
	lock sync.Mutex
	// data: url => status
	table map[string]bool
}

// create new UrlTable
func NewUrlTable() *UrlTable {
	ut := new(UrlTable)
	ut.table = make(map[string]bool)
	return ut
}

// add url to table
func (ut *UrlTable) Add(url string) {
	ut.lock.Lock()
	ut.table[url] = true
	ut.lock.Unlock()
}

// whether given url is in table
func (ut *UrlTable) Exist(url string) bool {
	ut.lock.Lock()
	_, ok := ut.table[url]
	ut.lock.Unlock()
	return ok
}
