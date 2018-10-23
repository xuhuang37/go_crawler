package fetcher

import (
	"net/http"
	"fmt"
	"golang.org/x/text/transform"
	"golang.org/x/net/html/charset"
	"io/ioutil"
	"golang.org/x/text/encoding"
	"bufio"
			"golang.org/x/text/encoding/unicode"
	"log"
)

func Fetch(url string)([]byte,error)  {
	resp, err := http.Get("http://www.zhenai.com/zhenghun/")
	if err != nil {
		return nil,err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil,fmt.Errorf("wrong status code:%d",resp.StatusCode)
	}
	bodyReader:=bufio.NewReader(resp.Body)
	e := determineEncoding(bodyReader)
	utf8Reader:= transform.NewReader(bodyReader,e.NewDecoder())
	return ioutil.ReadAll(utf8Reader)

}

func determineEncoding(r *bufio.Reader) encoding.Encoding {
	bytes,err:= r.Peek(1024)
	if err!=nil{
		log.Printf("Fetch Error:%v", err)
		return unicode.UTF8
	}
	e,_,_:=charset.DetermineEncoding(bytes,"")
	return e
}




