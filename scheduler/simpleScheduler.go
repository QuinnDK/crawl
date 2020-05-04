package scheduler

import "crawl/engine"

type SimpleScheduler struct {
	workerchan chan engine.Request
}

func (s *SimpleScheduler) Submit(r engine.Request) {
	go func() {
		s.workerchan <- r
	}()
}

func (s *SimpleScheduler) configureWorkChan(c chan engine.Request) {

	s.workerchan = c
}
