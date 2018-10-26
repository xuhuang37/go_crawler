package engine

import (
	"fmt"
	)

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
}

type Scheduler interface {
	ReadyNotifier
	Submit(Request)
	WorkerChan() chan Request
	Run()
}

type ReadyNotifier interface {
	WorkerReady(chan Request)
} 

func (e *ConcurrentEngine) Run(seeds ...Request) {

	//in := make(chan Request)                  //创建请求通道
	out := make(chan ParseResult)             // 创建解析结果通道
	//e.Scheduler.ConfigureMasterWorkerChan(in) // 传入请求通道 赋值给workerChan
	e.Scheduler.Run()

	for i := 0; i < e.WorkerCount; i++ {
		createWorker(e.Scheduler.WorkerChan(),out,e.Scheduler) // 创建多个worker协程
	}
	for _, req := range seeds {
		e.Scheduler.Submit(req) //
	}
	for {
		result := <-out //消费out
		for _, item := range result.Items {
			fmt.Printf("Got Item: %v\n", item)
		}
		for _, req := range result.Requests {
			e.Scheduler.Submit(req)
		}
	}
}

func createWorker(in chan Request ,out chan ParseResult,ready ReadyNotifier) {
	go func() {
		for {
			ready.WorkerReady(in)
			request := <-in                //消费请求
			result, err := worker(request) // 解析返回结果
			if err != nil {
				continue
			}
			out <- result // 为out生产结果
		}
	}()
}
