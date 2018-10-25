package parser

import (
	"testing"
	"io/ioutil"
	"log"
)

func TestParseCity(t *testing.T) {
	contents, err:=ioutil.ReadFile("city_test.html")
	if err!=nil{
		log.Println(err)
	}else{
		log.Println(ParseCity(contents))
	}
}