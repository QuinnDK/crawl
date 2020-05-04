package scheduler

import "crawl/engine"

type Scheduler interface {
	Submit(engine.Request)

	configureWorkChan(chan engine.Request)
}

type ConcurrentEngine struct {
	Scheduler Scheduler
	Workcount int
}

type SimpleScheduler struct {
	workerchan chan engine.Request
}
