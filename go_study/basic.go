package study

import "fmt"

func variableZeroValue() {
	var a int
	var s string
	fmt.Printf("%d %q", a, s)
}

func variableInitialValue() {
	var a, b int = 1, 4
	var s string = "huangxu"
	println(a, b, s)
}

func variableShorter() {
	a, b, c, s := 2, 3, "def", true
	println(a, b, c, s)
}

func main() {
	fmt.Print("hello world")
	variableZeroValue()
	variableInitialValue()
}
