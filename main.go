package main

import (
	"crawl/LearnGo-crawl/engine"
	"crawl/LearnGo-crawl/persist"

	//"crawl/parse"
	"crawl/LearnGo-crawl/parse/zhenai"
	"crawl/LearnGo-crawl/scheduler"
)

func main() {
	itemsave, err := persist.ItemSave()
	if err != nil {
		panic(err)
	}
	e := engine.ConcurrentEngine{
		&scheduler.QueueScheluder{},
		100,
		itemsave,
	}
	e.Run(engine.Request{
		Url:       "http://www.zhenai.com/zhenghun",
		ParseFunc: zhenai.ParseCity,
	})

	//result, _ := fetcher.Fetch("https://book.douban.com")
	//parse.ParseContent(result)

}
