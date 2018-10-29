package Persist

import (
	"testing"
	"go_crawler/model"
	"gopkg.in/olivere/elastic.v5"
	"context"
	"encoding/json"
	"go_crawler/engine"
)

func TestSave(t *testing.T) {
	expected := engine.Item{
		Url:  "http://album.zhenai.com/u/1121157968",
		Type: "zhenai",
		Id:   "1121157968",
		Payload: model.Profile{
			Name:       "雨天晴天",
			Gender:     "女",
			Age:        31,
			Height:     159,
			Weight:     55,
			Income:     "12001-20000元",
			Marriage:   "丧偶",
			Education:  "大专",
			Occupation: "会计",
			Location:   "重庆巴南区",
			HomeTown:   "四川宜宾",
			Zodiac:     "射手座",
			House:      "租房",
			Car:        "未购车",
		},
	}

	err := save(expected)
	if err != nil {
		panic(err)
	}

	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}

	resp, err := client.Get().Index("dating_profile").Type(expected.Type).Id(expected.Id).Do(context.Background())
	if err != nil {
		panic(err)
	}

	var actual engine.Item
	err = json.Unmarshal([]byte(*resp.Source), &actual)
	if err != nil {
		panic(err)
	}
	actualProfile,_:=model.FromJsonObject(actual.Payload)
	actual.Payload = actualProfile

	if expected != actual {
		t.Errorf("got %v expected %v", actual, expected)
	}

}
