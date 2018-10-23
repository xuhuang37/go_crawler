package parser

import (
	"regexp"
		"go_crawler/engine"
)
const cityListRe = `(http://www.zhenai.com/zhenghun/[a-z0-9]+)[^>]*>([^<]+)`
func ParserCityList(content []byte) engine.ParseResult  {
	regexCtx:=regexp.MustCompile(cityListRe)
	strArray:=regexCtx.FindAllSubmatch(content,-1)
	result:= engine.ParseResult{}
	for _,m := range strArray {
		result.Items = append(result.Items, string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:string(m[1]),
			ParserFunc: engine.NilParser,
		})
		//fmt.Printf("city:%s Url:%s",m[2],m[1])
	}
	return result
}
