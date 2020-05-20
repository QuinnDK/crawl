package main

import (
	"crawl/LearnGo-crawl/crawl_distribute/rpcsupport"
	"crawl/LearnGo-crawl/crawl_distribute/work/server"
	"log"
)

func main() {
	log.Fatal(rpcsupport.ServeRpc(":1235", &server.CrawlService{}))
}
