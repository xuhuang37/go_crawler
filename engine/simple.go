package engine

import (
	"go_crawler/fetch"
	"log"
	)

type SimpleEngine struct {
}

func (s SimpleEngine) Run(seeds ...Request) {
	var requests []Request
	for _, seed := range seeds {
		requests = append(requests, seed)
	}
	for len(requests) > 0 {
		headReq := requests[0]
		requests = requests[1:]
		parseResult, err := worker(headReq)
		if err != nil {
			continue
		}

		requests = append(requests, parseResult.Requests...)
		for _, item := range parseResult.Items {
			log.Printf("Got item %v", item)
		}
	}
}

func worker(headReq Request) (ParseResult, error) {
	body, err := fetcher.Fetch(headReq.Url)
	if err != nil {
		log.Printf("Fetcher error: fetcher url %s :%v \n", headReq.Url, err)
		return ParseResult{}, err
	}
	return headReq.ParserFunc(body), nil
}
