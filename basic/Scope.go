package main

import "fmt"

// 局部变量
func Local() {
	var a, b, c int
	a = 1
	b = 2
	c = a + b
	fmt.Printf("a的值 %d；b的值 %d；c的值 %d", a, b, c)
}

//全局变量,优先考虑局部变量
var zy int
var zz = 1

func OverAll() {
	var z, y, total int
	z = 1
	y = 2
	zz = z + y
	total = Sum(z, zz)
	fmt.Printf("zy的值为 %d;zz的值为%d; total的值%d", zy, zz, total)
}
func Sum(x, y int) int {
	return x + y
}

func main() {
	Local()
	fmt.Println("\n")

	OverAll()
}
