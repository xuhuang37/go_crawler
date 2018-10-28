package scheduler

import (
	"go_crawler/engine"
	)

type QueueScheduler struct {
	requestChan chan engine.Request
	workerChan  chan chan engine.Request
}

func (q *QueueScheduler)WorkerChan()chan engine.Request  {
	return make(chan engine.Request)
}

func (q *QueueScheduler) Submit(r engine.Request) {
	q.requestChan <- r
}

func (q *QueueScheduler) WorkerReady(w chan engine.Request) {
		q.workerChan <- w
}

func (q *QueueScheduler) Run() {

	q.requestChan = make(chan engine.Request)
	q.workerChan = make(chan chan engine.Request)

	go func() {
		var requestQ []engine.Request
		var workerQ []chan engine.Request
		for {
			var activeRequest engine.Request
			var activeWorker chan engine.Request
			if len(requestQ) > 0 && len(workerQ) > 0 {
				activeRequest = requestQ[0]
				activeWorker = workerQ[0]
			}
			select {
			case r := <-q.requestChan:
				requestQ = append(requestQ, r)
			case w := <-q.workerChan:
				workerQ = append(workerQ, w)
			case activeWorker <- activeRequest:
				workerQ = workerQ[1:]
				requestQ = requestQ[1:]
			}
		}
	}()
}
