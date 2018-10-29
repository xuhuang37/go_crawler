package engine

import "sync"

type Request struct {
	Url        string
	ParserFunc func([]byte) ParseResult
}

type ParseResult struct {
	Requests []Request
	Items    []Item
}

type Item struct {
	Url     string
	Id      string
	Type    string
	Payload interface{}
}

// 去重映射
type UrlMap struct {
	UrlList map[string]int
	Lock    *sync.RWMutex
}

func (u *UrlMap) Valid(k string, v int) bool {
	u.Lock.Lock()
	defer u.Lock.Unlock()
	_, ok := u.UrlList[k]
	if !ok {
		u.UrlList[k] = v
		return false
	} else {
		return true
	}
}

func (u *UrlMap) Set(k string, v int) {
	u.Lock.Lock()
	defer u.Lock.Unlock()
	u.UrlList[k] = v
}
