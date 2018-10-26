package parser

import (
	"go_crawler/engine"
	"regexp"
	"strconv"
	"go_crawler/model"
	"log"
)

var ageRe = regexp.MustCompile(`<td><span class="label">年龄：</span>([\d]+)岁</td>`)
var heightRe = regexp.MustCompile(`<td><span class="label">身高：</span><span field="">([\d]+)CM</span></td>`)
var weightRe = regexp.MustCompile(`<td><span class="label">体重：</span><span field="">([\d]+)KG</span></td>`)
var genderRe = regexp.MustCompile(`<td><span class="label">月收入：</span>([^<]+)</td>`)
var marriageRe = regexp.MustCompile(`<td><span class="label">婚况：</span>([^<]+)</td>`)
var educationRe = regexp.MustCompile(`<td><span class="label">学历：</span>([^<]+)</td>`)
var LocationRe = regexp.MustCompile(`<td><span class="label">工作地：</span>([^<]+)</td>`)
var OccupationRe = regexp.MustCompile(`<td><span class="label">职业：</span>([^<]+)</td>`)
var HometownRe = regexp.MustCompile(`<td><span class="label">籍贯：</span>([^<]+)</td>`)
var ZodiacRe = regexp.MustCompile(`<td><span class="label">星座：</span><span field="">([^<]+)</span></td>`)
var HouseRe = regexp.MustCompile(`<td><span class="label">住房条件：</span><span field="">([^<]+)</span></td>`)
var CarRe = regexp.MustCompile(`<td><span class="label">是否购车：</span><span field="">([^<]+)</span></td>`)

func ParseProfile(contents []byte, userName string) engine.ParseResult {
	match := ageRe.FindSubmatch(contents)
	profile := model.Profile{}
	profile.Name = userName
	if match != nil {
		age, err := strconv.Atoi(string(match[1]))
		if err != nil {
			log.Println(err)
		}else {
			profile.Age = age
		}
	}
	match = weightRe.FindSubmatch(contents)
	if match != nil {
		weight, err := strconv.Atoi(string(match[1]))
		if err != nil {
			log.Println(err)
		}else {
			profile.Weight = weight
		}
	}
	match = heightRe.FindSubmatch(contents)
	if match != nil {
		height, err := strconv.Atoi(string(match[1]))
		if err != nil {
			log.Println(err)
		}else {
			profile.Height = height
		}
	}
	profile.Gender = extractString(contents,genderRe)
	profile.Marriage = extractString(contents,marriageRe)
	profile.Education = extractString(contents,educationRe)
	profile.Location = extractString(contents,LocationRe)
	profile.Occupation = extractString(contents,OccupationRe)
	profile.HomeTown = extractString(contents,HometownRe)
	profile.Zodiac = extractString(contents,ZodiacRe)
	profile.House = extractString(contents,HouseRe)
	profile.Car = extractString(contents,CarRe)
	result:=engine.ParseResult{
		Items:[]interface{}{profile},
	}
	return result

}

func extractString(contents []byte, reg *regexp.Regexp) string {
	match:=reg.FindSubmatch(contents)
	if len(match)>=2{
		return string(match[1])
	}else{
		return ""
	}
}

