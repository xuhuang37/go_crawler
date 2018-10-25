package parser

import (
	"log"
	"testing"
			"io/ioutil"
)

func TestParseProfile(t *testing.T) {
	contents,err:=ioutil.ReadFile("profile_test.html")
	if err != nil {
		log.Println(err)
	} else {
		log.Println(ParseProfile(contents))

	}
}
