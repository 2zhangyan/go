package main

import "fmt"

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

}
