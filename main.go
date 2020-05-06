package main

import (
	"crawl/engine"
	//"crawl/parse"
	"crawl/parse/zhenai"
	"crawl/scheduler"
)

func main() {
	e := engine.ConcurrentEngine{
		&scheduler.QueueScheluder{},
		100,
	}
	e.Run(engine.Request{
		Url:       "http://www.zhenai.com/zhenghun",
		ParseFunc: zhenai.ParseCity,
	})

	//result, _ := fetcher.Fetch("https://book.douban.com")
	//parse.ParseContent(result)

}
