package main

import (
	"io/ioutil"
	"fmt"
	"path/filepath"
	"os"
	"log"
	"strings"
)

func fileIo() {
	const filename = "abc.txt"
	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}
}

func grade(score int) string {
	g := ""
	switch {
	case score > 100 || score < 0:
		panic(fmt.Sprintf("Wrong score %d", score))
	case score < 60:
		g = "E"
	case score < 70:
		g = "D"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score < 100:
		g = "A"
	}
	return g
}

func GetCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

func main() {
	fileIo()
	println(grade(101))

}
