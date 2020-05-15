package main

import (
	"crawl/engine"
	"crawl/persist"

	//"crawl/parse"
	"crawl/parse/zhenai"
	"crawl/scheduler"
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
