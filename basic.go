package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

var aa = 3
var bb = "Jiaqi in package"
var ss = 0.07

func variableZeroValue() {
	var a int
	var s string
	fmt.Println(a, s)
	fmt.Printf("%d %q\n", a, s)
}
func variableInitialValue() {
	var a, b int = 10, 4
	var s string = "Jiaqi"
	fmt.Println(a, b, s)
}

func variableTypeDeduction() {
	var a, b, s, f = 10, true, "Jiaqi in func", 0.7 // 繁琐
	fmt.Println(a, b, s, f)
}

func variableShorter() {
	a, b, s, f := 10, true, "Jiaqi in func", 0.7
	fmt.Println(a, b, s, f)
}
func euler() {
	//c := 3 + 4i
	fmt.Printf("%.3f\n",
		cmplx.Exp(1i*math.Pi)+1)
}

func triangle() {
	a, b := 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

func consts() {
	const (
		filename = "abc,txt"
		a, b     = 3, 4
	)
	//const filename = "abc.txt"
	//const a, b = 3, 4
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c)
}

func enums() {
	const (
		cpp        = 1
		python     = 2
		c          = 3
		java       = 4
		javascript = 5
	)
	// b kb mb gb tb pb eb
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
		eb
	)
	fmt.Println(cpp, python, c, java, javascript)
	fmt.Println(b, kb, mb, gb, tb, pb, eb)
}

func main() {
	fmt.Println("Hello world")
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()
	fmt.Println(aa, bb, ss)
	euler()
	triangle()
	consts()
	enums()
}
