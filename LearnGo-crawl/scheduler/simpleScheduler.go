package scheduler

import "crawl/LearnGo-crawl/engine"

type SimpleScheduler struct {
	workerchan chan engine.Request
}

func (s *SimpleScheduler) Run() {
	s.workerchan = make(chan engine.Request)
}

func (s *SimpleScheduler) WorkReady(chan engine.Request) {
	return
}

func (s *SimpleScheduler) WorkChan() chan engine.Request {
	return s.workerchan
}

func (s *SimpleScheduler) Submit(r engine.Request) {
	go func() {
		s.workerchan <- r
	}()
}
