package main

import (
	"fmt"
	"unsafe"
)

// 常量
func Const() {

	fmt.Println("=======常量start======")

	const TestName string = "Test" //显示类型定义
	const TestAge = 18             //隐式类型定义
	const TestRegion, TestAddress = "中国", "北京"
	fmt.Println(TestName, TestAge, TestRegion, TestAddress)

	const (
		a = "zzz"
		b = len(a) //可以用内置函数
		c = unsafe.Sizeof(b)
		d = iota
		e
		q = "yyy"
		w = iota
		f
	)
	fmt.Println(a, b, c, e, q, w, f)

	//左移 128 64 32 16 8 4 2 1 0
	const (
		p = 1 << iota
		j = 4 << iota
		k
		l
	)
	fmt.Println(p, j, k, l)
}

//运算符 + - * / % ++ --
func OperatorArithmetic() {

	fmt.Println("=======运算符start======")

	//算术运算符 + - * / % ++ --
	var num = 1
	var num2 = 2

	c := num + num2
	fmt.Println(c)

	c = num2 - num
	fmt.Println(c)

	c = num * num2
	fmt.Println(c)

	//正常的两个数字相除应该是0.5 但是int所以是0
	c = num / num2
	fmt.Printf("除法的的值为 %d \n", c)
	//想要正确的值 则先转换float后在做触发
	fmt.Printf("除法的的值为 %.2f \n", float64(num)/float64(num2))

	c = num2 % num
	fmt.Println(c)

	c++
	fmt.Println(c)

	c--
	fmt.Println(c)

}

//关系运算符  == != > < >= <=
func OperatorRelationship() {
	fmt.Println("=======关系运算符start======")

	num := 21
	num2 := 2

	if num == num2 {
		fmt.Println("num 等于 num2")
	} else {
		fmt.Println("num 不等于 num2")
	}

	//都差不多的判断不一一展开了
}

//逻辑运算符 && || !
func OperatorLogic() {
	fmt.Println("=======逻辑运算符start======")

	a := true
	b := true

	if a && b {
		fmt.Println("哎哟喂，为真呢")
	}

	if a || b {
		fmt.Println("is true")
	}
}

func main() {
	Const()
	OperatorArithmetic()
	OperatorRelationship()
	OperatorLogic()
}
