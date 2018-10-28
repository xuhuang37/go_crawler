package parser

import (
	"go_crawler/engine"
	"regexp"
	)

var profileRe = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[\d]+)"[^>]+>([^<]+)</a>`)
var cityUrlRe = regexp.MustCompile(`href="(http://www.zhenai.com/zhenghun/[^"]+)"`)

func ParseCity(contents []byte) engine.ParseResult {
	matches := profileRe.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, ele := range matches {
		profileUrl := string(ele[1])
		userName := string(ele[2])
		req := engine.Request{
			Url: profileUrl,
			ParserFunc: func(contents []byte) engine.ParseResult {
				result := ParseProfile(contents, userName)
				return result
			},
		}
		result.Requests = append(result.Requests, req)
	}
	matches = cityUrlRe.FindAllSubmatch(contents,-1)

	for _,m :=range matches{
		result.Requests = append(result.Requests,engine.Request{
			Url:string(m[1]),
			ParserFunc: func(content []byte) engine.ParseResult {
				return ParseCity(content)
			},
		})
	}
	return result
}

