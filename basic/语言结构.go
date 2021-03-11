package main

import "fmt"

func hi() {
	var age = "1"
	var age_flx = "9"
	fmt.Println(age + age_flx)

	var a, b int = 1, 2
	fmt.Println(a, b)
}

func zy() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Println(i, f, b, s)
}

func valtype() {
	var d = true
	fmt.Println(d)
}

func valtob() {
	var intvalOne int
	intvalOne = 1
	intvalTwo := 2

	fmt.Println(intvalOne, intvalTwo)
}

var x, y int
var (
	a int
	b bool
)
var c, d int = 1, 3
var e, f = 123, "hello"

func vals() {
	g, h := "z", "y"
	fmt.Println(g, h)
}

func main() {
	hi()
	zy()
	valtype()
	valtob()
	vals()
	print(x, y, a, b, c, d, e)
}
