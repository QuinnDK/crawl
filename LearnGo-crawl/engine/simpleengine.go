package engine

import (
	"crawl/LearnGo-crawl/fetcher"
	"fmt"
	"log"
)

type simpleEngine struct {
}

func (s simpleEngine) Run(seeds ...Request) {
	var requests []Request
	for _, e := range seeds {
		requests = append(requests, e)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]
		log.Printf("Fetching url:%s", r.Url)
		body, err := fetcher.Fetch(r.Url)
		if err != nil {
			log.Printf("Fetching error:%s", r.Url)
		}
		parseResult := r.ParseFunc(body)
		requests = append(requests, parseResult.Requesrts...)

		for _, item := range parseResult.Items {
			fmt.Printf("Go Item:%s\n", item)
		}
	}
}
