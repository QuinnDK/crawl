package main

import (
	"crawl/LearnGo-crawl/crawl_distribute/client"
	"crawl/LearnGo-crawl/engine"
	//"crawl/LearnGo-crawl/persist"

	client2 "crawl/LearnGo-crawl/crawl_distribute/work/client"
	//"crawl/parse"
	"crawl/LearnGo-crawl/parse/zhenai"
	"crawl/LearnGo-crawl/scheduler"
)

func main() {
	//itemsave, err := persist.ItemSave()
	itemsave, err := client.ItemSave("1234")
	process, err := client2.CreateProcess()
	if err != nil {
		panic(err)
	}
	e := engine.ConcurrentEngine{
		&scheduler.QueueScheluder{},
		100,
		itemsave,
		process,
	}
	e.Run(engine.Request{
		Url:   "http://www.zhenai.com/zhenghun",
		Parse: engine.NewFuncparse(zhenai.ParseCityList, "ParseCityList"),
	})

	//result, _ := fetcher.Fetch("https://book.douban.com")
	//parse.ParseContent(result)

}
