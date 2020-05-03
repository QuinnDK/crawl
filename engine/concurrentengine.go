package engine

import (
	"crawl/fetcher"
	"fmt"
	"log"
)

type Scheduler interface {
	Submit(Request)

	configureWorkChan(chan Request)
}

type ConcurrentEngine struct {
	Scheduler Scheduler
	Workcount int
}

type SimpleScheduler struct {
	workerchan chan Request
}

func (s *SimpleScheduler) Submit(r Request) {

	s.workerchan <- r
}

func (s *SimpleScheduler) configureWorkChan(c chan Request) {

	s.workerchan = c
}

func (e *ConcurrentEngine) Run(seeds ...Request) {

	in := make(chan Request)
	out := make(chan ParseResult)

	for i := 0; i < e.Workcount; i++ {

		CreateWork(in, out)
	}

	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}

}

func CreateWork(in chan Request, out chan ParseResult) {

	go func() {
		for {
			request := <-in

			result, err := worker(request)

			if err != nil {
				continue
			}
			out <- result
		}
	}()

}

func worker(r Request) (ParseResult, error) {
	fmt.Printf("Fetching url:%s", r.Url)
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetching Error:%s", r.Url)
		return ParseResult{}, err
	}
	return r.ParseFunc(body), nil

}
