package parser

import (
	"go_crawler/engine"
	"regexp"
	)

var cityRe = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[\d]+)"[^>]+>([^<]+)</a>`)

func ParseCity(contents []byte) engine.ParseResult {
	matches := cityRe.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, ele := range matches {
		userName := string(ele[1])
		req := engine.Request{
			Url: userName,
			ParserFunc: func(contents []byte) engine.ParseResult {
				result := ParseProfile(contents, userName)
				return result
			},
		}
		result.Requests = append(result.Requests, req)
		result.Items = append(result.Items, "Name "+string(ele[2]))
	}
	return result
}
