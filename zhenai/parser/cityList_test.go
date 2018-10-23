package parser

import (
	"testing"
			"io/ioutil"
)

func TestParserCityList(t *testing.T) {
	contents,err:=ioutil.ReadFile("city_list_test.html")
	if err!=nil{
		panic(err)
	}
	results := ParserCityList(contents)
	expectedUrl := []string{
		"http://www.zhenai.com/zhenghun/aba",
		"http://www.zhenai.com/zhenghun/akesu",
		"http://www.zhenai.com/zhenghun/alashanmeng",
	}
	//expectedCity:= []string{
	//	"阿克苏","阿拉善盟","阿勒泰",
	//}

	const resultSize = 540
	if n:= len(results.Requests);n != resultSize{
		t.Errorf("Result should have %d,but %d",resultSize,n)
	}
	if n:= len(results.Items);n != resultSize{
		t.Errorf("Items should have %d,but %d",resultSize,n)
	}

	for i,ex := range expectedUrl{
		if cur:=results.Requests[i].Url;cur!=ex{
			t.Errorf("expected url #%d: %s,but was %s",i,ex,cur)
		}
	}

}
