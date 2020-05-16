package main

import (
	"crawl/LearnGo-crawl/crawl_distribute/persist"
	"crawl/LearnGo-crawl/crawl_distribute/rpcsupport"
	"gopkg.in/olivere/elastic.v5"
)

func main() {

	//itemsave,err:= client.ItemSave(":1234")
	//
	//if err!=nil{
	//	panic(err)
	//}
	//e:= engine.ConcurrentEngine{
	//	&scheduler.QueueScheduler{},
	//	100,
	//	itemsave,
	//}
	//
	//e.Run(engine.Request{
	//	Url:"http://www.zhenai.com/zhenghun",
	//	Parse:engine.NewFuncparse(zhengai.ParseCity,"Parsecity") ,
	//})

	serveRpc("1234")
}

func serveRpc(host string) error {

	client, err := elastic.NewClient(elastic.SetSniff(false))

	if err != nil {
		return nil
	}
	return rpcsupport.ServeRoc(host, &persist.ItemService{
		Client: client,
	})
}
