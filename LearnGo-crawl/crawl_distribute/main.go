package main

import (
	"crawl/LearnGo-crawl/crawl_distribute/client"
	"crawl/LearnGo-crawl/engine"
	"crawl/LearnGo-crawl/parse/zhenai"
	"crawl/LearnGo-crawl/scheduler"
)

func main() {

	itemsave, err := client.ItemSave(":1234")

	if err != nil {
		panic(err)
	}
	e := engine.ConcurrentEngine{
		&scheduler.QueueScheluder{},
		100,
		itemsave,
	}

	e.Run(engine.Request{
		Url:   "http://www.zhenai.com/zhenghun",
		Parse: engine.NewFuncparse(zhenai.ParseCity, "Parsecity"),
	})
}
