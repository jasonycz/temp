/* mini_spider.go - main program of mini-spider   */
/*
modification history
--------------------
2015/6/14, by Zhang Miao, create
*/
/*
DESCRIPTION
*/
package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"path"
	"syscall"
	"time"
)

import (
	"www.baidu.com/golang-lib/log"
)

import (
	"mini_spider"
	"spider_conf"
)

var (
	help     *bool   = flag.Bool("h", false, "to show help")
	confRoot *string = flag.String("c", "../../conf", "root path of config file")
	logPath  *string = flag.String("l", "./log", "dir path of log")
	stdOut   *bool   = flag.Bool("s", false, "to show log in stdout")
	debugLog *bool   = flag.Bool("d", false, "to show debug log (otherwise >= info)")
)

func Exit(code int) {
	log.Logger.Close()
	/* to overcome bug in log, sleep for a while    */
	time.Sleep(1 * time.Second)
	os.Exit(code)
}

/* the main function. */
func main() {
	var err error
	var logSwitch string

	flag.Parse()
	if *help {
		flag.PrintDefaults()
		return
	}

	// debug switch
	if *debugLog {
		logSwitch = "DEBUG"
	} else {
		logSwitch = "INFO"
	}

	err = log.Init("mini-spider", logSwitch, *logPath, *stdOut, "midnight", 5)
	if err != nil {
		fmt.Printf("mini-spider: err in log.Init():%s\n", err.Error())
		Exit(1)
	}

	log.Logger.Info("mini-spider start")

	// load config
	confPath := path.Join(*confRoot, "mini_spider.conf")
	config, err := spider_conf.ConfigLoad(confPath)
	if err != nil {
		log.Logger.Error("main():err in ConfigLoad(%s):%s", confPath, err.Error())
		Exit(1)
	}

	// load seeds
	seeds, err := mini_spider.SeedFileLoad(config.Spider.UrlListFile)
	if err != nil {
		log.Logger.Error("main():err in SeedFileLoad(%s):%s",
			config.Spider.UrlListFile, err.Error())
		Exit(1)
	}

	// create mini-spider
	miniSpider := mini_spider.NewMiniSpider(&config, seeds)

	// start mini-spider
	miniSpider.Start()

	// Handle SIGINT and SIGTERM.
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch

	// ensure that all logs are export and normal exit
	Exit(0)
}
