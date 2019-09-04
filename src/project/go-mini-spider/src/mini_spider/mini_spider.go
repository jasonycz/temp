/* mini_spider.go - mini spider	*/
/*
modification history
--------------------
2015/6/14, by Zhang Miao, create
*/
/*
DESCRIPTION
*/
package mini_spider

import (
	"www.baidu.com/golang-lib/queue"
)

import (
	"spider_conf"
)

type MiniSpider struct {
	config *spider_conf.SpiderConfig

	urlTable *UrlTable
	queue    queue.Queue

	crawlers []*CrawlThread
}

// create new mini-spider
func NewMiniSpider(conf *spider_conf.SpiderConfig, seeds []string) *MiniSpider {
	ms := new(MiniSpider)
	ms.config = conf

	// create url-table
	ms.urlTable = NewUrlTable()

	// initialize queue
	ms.queue.Init()
	// add seeds to queue
	for _, seed := range seeds {
		msg := &crawlMsg{Url: seed, Depth: 1}
		ms.queue.Append(msg)
	}

	// create crawlers
	ms.crawlers = make([]*CrawlThread, 0)
	for i := 0; i < conf.Spider.ThreadCount; i++ {
		crawler := NewCrawlThread(ms.urlTable, ms.config, &ms.queue)
		ms.crawlers = append(ms.crawlers, crawler)
	}

	return ms
}

// start mini spider
func (ms *MiniSpider) Start() {
	// start all crawlers
	for _, crawler := range ms.crawlers {
		crawler.Start()
	}
}
