package parser

import (
	"regexp"
	"go_crawler/engine"
)

const cityListRe = `(http://www.zhenai.com/zhenghun/[a-z0-9]+)[^>]*>([^<]+)`

func ParserCityList(content []byte) engine.ParseResult {
	regexCtx := regexp.MustCompile(cityListRe)
	matches := regexCtx.FindAllSubmatch(content, -1)
	result := engine.ParseResult{}

	for _, m := range matches {
		result.Items = append(result.Items, "City "+string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			ParserFunc: ParseCity,
		})
	}
	return result
}
