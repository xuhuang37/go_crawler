package engine

import (
	"go_crawler/fetch"
	"log"
)

func Run(seeds ...Request) {
	var requests []Request
	for _, seed := range seeds {
		requests = append(requests, seed)
	}
	for len(requests) > 0 {
		headReq := requests[0]
		requests = requests[1:]
		log.Printf("Fetching %s", headReq.Url)
		parseResult,err:= worker(headReq)
		if err!=nil{

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
		log.Printf("Fetcher error: fetcher url %s :%v", headReq.Url, err)
		return ParseResult{}, err

	} else {
		return headReq.ParserFunc(body), nil
	}

}
