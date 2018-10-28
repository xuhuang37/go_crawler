package main

import (
	"go_crawler/engine"
	"go_crawler/zhenai/parser"
	"go_crawler/scheduler"
	"go_crawler/Persist"
)

func main() {
	e:=&engine.ConcurrentEngine{
		Scheduler:&scheduler.SimpleScheduler{},
		WorkerCount:10,
		ItemChan:Persist.ItemSaver(),
	}
	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun/",
		ParserFunc: parser.ParserCityList,
	})// 启动并行爬虫

}
