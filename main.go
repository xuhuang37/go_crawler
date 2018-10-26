package main

import (
	"go_crawler/engine"
	"go_crawler/zhenai/parser"
	"go_crawler/scheduler"
)

func main() {

	e:=&engine.ConcurrentEngine{
		Scheduler:&scheduler.SimpleScheduler{},
		WorkerCount:100,
	}
	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun/",
		ParserFunc: parser.ParserCityList,
	})// 启动并行爬虫

}
