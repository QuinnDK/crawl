package main

import (
	"crawl/engine"
	"crawl/parse"
)

func main() {
	e := engine.ConcurrentEngine{
		&engine.SimpleScheduler{},
		100,
	}
	e.Run(engine.Request{
		Url:       "https://book.douban.com",
		ParseFunc: parse.ParseTag,
	})

	//result, _ := fetcher.Fetch("https://book.douban.com")
	//parse.ParseContent(result)

}
