package main

import (
	"fmt"
	"math/cmplx"
	"math"
)

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

func euler() {
	//c := 3 + 4i
	//println(cmplx.Abs(c))
	fmt.Printf("%.3f\n", cmplx.Pow(math.E, 1i*math.Pi)+1)
}

func triangle() {
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	println(c)
}

func consts() {
	const filename = "abc.txt"
	const a, b = 3, 4
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c)

}

func enums() {
	const (
		cpp  = iota
		java
		python
		golang
		javascript
	)

	const(
		b = 1<<(10*iota)
		kb
		mb
		gb
		tb
	)
	//println(cpp,java,python,golang,javascript)
	println(b,kb,mb,gb,tb)


}
func main() {
	fmt.Println("hello world")
	variableZeroValue()
	variableInitialValue()
	euler()
	triangle()
	consts()
	enums()
}
