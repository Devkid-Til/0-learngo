package main

import (
	"fmt"
	"io/ioutil"
)

func grade(score int) string {
	var g string
	switch {
	case score < 60:
		g = "F"
	case score < 70:
		g = "D"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"
	default:
		panic(fmt.Sprint("wrong score: %d\n", score))
	}
	return g
}

func eval(a, b int, op string) int {
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		q, _ := div(a, b)
		return q
	default:
		panic("unsuported operation: " + op)
	}
}

func div(a, b int) (q, r int) {
	//fmt.Println(a, "/", b, "=", a/b, "余", a%b)
	return a / b, a % b
}

func main() {
	const filename = "abc.txt"
	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}
	fmt.Println(
		grade(60),
		grade(75),
		grade(80),
		grade(91),
		grade(100),
	)

	fmt.Println(
		eval(12, 2, "/"),
		//eval(1.4, 2, ")"),
	)

	a, b := div(10, 4)
	fmt.Println(a, b)

}
