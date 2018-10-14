package main

import (
	"io/ioutil"
	"fmt"
)

func fileIo(){
	const filename = "abc.txt"
	contents, err :=ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	}else {
		fmt.Printf("%s\n", contents)
	}
}

func main()  {
	fileIo()
}
