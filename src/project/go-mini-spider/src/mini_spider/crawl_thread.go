/* crawl_thread.go - thread for crawl files	*/
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
	"regexp"
	"time"
)

import (
	"www.baidu.com/golang-lib/http_util"
	"www.baidu.com/golang-lib/log"
	"www.baidu.com/golang-lib/queue"
)

import (
	"spider_conf"
)

type crawlMsg struct {
	Url   string // url to crawl
	Depth int    // depth of this url
}

type CrawlThread struct {
	urlTable *UrlTable
	config   *spider_conf.ConfBasic

	queue      *queue.Queue
	urlPattern *regexp.Regexp

	active bool
}

// create new crawl thread
func NewCrawlThread(urlTable *UrlTable, config *spider_conf.SpiderConfig,
	queue *queue.Queue) *CrawlThread {
	ct := new(CrawlThread)
	ct.urlTable = urlTable
	ct.config = &config.Spider
	ct.queue = queue

	// TargetUrl has been check in conf load
	ct.urlPattern, _ = regexp.Compile(ct.config.TargetUrl)

	ct.active = true

	return ct
}

// start crawl thread
func (ct *CrawlThread) Start() {
	for ct.active {
		// get new task from queue
		msg := ct.queue.Remove().(*crawlMsg)
		log.Logger.Debug("from queue: url=%s, depth=%d", msg.Url, msg.Depth)

		// read data from given url
		data, err := http_util.Read(msg.Url, ct.config.CrawlTimeout)
		if err != nil {
			log.Logger.Error("http_util.Read(%s):%s", msg.Url, err.Error())
			continue
		}

		// save data to file
		if ct.urlPattern.MatchString(msg.Url) {
			err = urlFileSave(ct.config.OutputDirectory, msg.Url, data)
			if err != nil {
				log.Logger.Error("urlFileSave(%s):%s", msg.Url, err.Error())
			} else {
				log.Logger.Debug("save to file: %s", msg.Url)
			}
		}
		
		// add to url table
		ct.urlTable.Add(msg.Url)

		if msg.Depth < ct.config.MaxDepth {
			// extract url from web page
			links, err := urlExtract(data, msg.Url)
			if err != nil {
				log.Logger.Error("urlExtract():%s", err.Error())
				continue
			}

			for _, link := range links {
				// check whether url match the pattern, or url exists already
				if ct.urlTable.Exist(link) {
					continue
				}

				msgNew := &crawlMsg{Url: link, Depth: msg.Depth + 1}
				log.Logger.Debug("add to queue: url=%s, depth=%d", msgNew.Url, msgNew.Depth)
				ct.queue.Append(msgNew)
			}
		}

		// sleep for a while
		time.Sleep(time.Duration(ct.config.CrawlInterval) * time.Second)
	}
}

// stop crawl thread
func (ct *CrawlThread) Stop() {
	ct.active = false
}
