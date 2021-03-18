package main

import (
	"fmt"
	"math"
)

func max(num1, num2 int) int {
	var result int

	if num1 >= num2 {
		result = num1
	} else {
		result = num2
	}
	return result

}

//交换值 引用
func SwapQuote(x *int, y *int) {
	var temp int
	temp = *x
	*x = *y
	*y = temp

}

//交换值 传值
func SwapVal(x int, y int) int {
	var temp int
	temp = x
	x = y
	y = temp

	return temp
}

//持匿名函数，可作为闭包。匿名函数是一个"内联"语句或表达式。匿名函数的优越性在于可以直接使用函数内的变量，不必申明。
//以下实例中，我们创建了函数 GetSequence() ，返回另外一个函数。该函数的目的是在闭包中递增 i 变量，代码如下：
func GetSequence() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

type Circle struct {
	radius float64
}

//属于 Circle 类型对象中的方法
func (c Circle) getArea() float64 {
	//c.radius 即为 Circle 类型对象中的属性
	return 3.14 * c.radius * c.radius
}

func main() {
	//函数调用
	var a = 110
	var b = 120
	MaxNum := max(a, b)
	fmt.Printf("最大值%d", MaxNum)
	fmt.Printf("\n")

	//交换a,b的值
	SwapVal(a, b)
	fmt.Printf("\n")
	fmt.Printf("交换后a的值：%d", a)
	fmt.Printf("\n")
	fmt.Printf("交换后b的值：%d", b)

	SwapQuote(&a, &b)
	fmt.Printf("\n")
	fmt.Printf("交换后a的值：%d", a)
	fmt.Printf("\n")
	fmt.Printf("交换后b的值：%d", b)
	fmt.Printf("\n")

	//函数作为另外一个函数的实参

	//声明函数变量
	getSquareRoot := func(x float64) float64 {
		return math.Sqrt(x)
	}
	//使用函数
	fmt.Println(getSquareRoot(25))

	//闭包- 匿名函数
	NextNumber := GetSequence()
	fmt.Println(NextNumber())
	fmt.Println(NextNumber())
	fmt.Println(NextNumber())

	NextNumber1 := GetSequence()
	fmt.Println(NextNumber1())
	fmt.Println(NextNumber1())
	fmt.Println(NextNumber1())

	//方法就是一个包含了接受者的函数
	var c1 Circle
	c1.radius = 7.00
	fmt.Println("圆的面积 = ", c1.getArea())
}
