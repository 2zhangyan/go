package main

import "fmt"

/**
	[ 类型 ]
		布尔型 		true false
		数字类型 		int float32 float64
		字符串类型	UTF-8 编码标识 Unicode 文本
		派生类型	{
					(1) 指针类型（Pointer）
					(2) 数组类型
					(3) 结构化类型(struct)
					(4) Channel 类型
					(5) 函数类型
					(6) 切片类型
					(7) 接口类型（interface）
					(8) Map 类型
		}
**/

var x, y int
var (
	a int
	b bool
)
var c, d int = 1, 3
var e, f = 123, "hello"

//字符连接
func StringConnect() {
	var age = "1"
	var age_flx = "9"
	fmt.Println(age + age_flx)

	var a, b int = 1, 2
	fmt.Println(a, b)
}

//定义变量
func variable() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Println(i, f, b, s)
	var d = true
	fmt.Println(d)
}

//定义变量,如果var生成过变量 不要用 := 否则编译有错误
func variableTwo() {
	g, h := "z", "y"
	fmt.Println(g, h)

	var intvalOne int
	intvalOne = 1
	intvalTwo := 2
	fmt.Println(intvalOne, intvalTwo)
}

func main() {
	StringConnect()
	variable()
	variableTwo()
	print(x, y, a, b, c, d, e)
}
