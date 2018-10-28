package Persist

import (
	"testing"
	"go_crawler/model"
)

func TestSave(t *testing.T) {
	profile:=model.Profile{
		Name:"雨天晴天",
		Gender:"女",
		Age:31,
		Height:159,
		Weight:55,
		Income:"12001-20000元",
		Marriage:"丧偶",
		 Education:"大专",
		Occupation:"会计" ,
		Location:"重庆巴南区",
		HomeTown:"四川宜宾" ,
		Zodiac:"射手座",
		House:"租房",
		Car:"未购车",
	}
	save(profile)
}
