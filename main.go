package main

import (
	"crawl/engine"
	"crawl/parse"
)

func main() {
	engine.Run(engine.Request{
		Url:       "https://book.douban.com",
		ParseFunc: parse.ParseBookList,
	})

	//result, _ := fetcher.Fetch("https://book.douban.com")
	//parse.ParseContent(result)

}
