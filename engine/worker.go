package engine

import (
	"go_crawler/fetch"
	"log"
)

func worker(headReq Request) (ParseResult, error) {
	body, err := fetcher.Fetch(headReq.Url)
	if err != nil {
		log.Printf("Fetcher error: fetcher url %s :%v \n", headReq.Url, err)
		return ParseResult{}, err
	}
	return headReq.ParserFunc(body), nil
}
