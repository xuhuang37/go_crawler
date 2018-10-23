package main

import (
	"go_crawler/engine"
	"go_crawler/zhenai/parser"
)

func main() {

	engine.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun/",
		ParserFunc: parser.ParserCityList,
	})

}
