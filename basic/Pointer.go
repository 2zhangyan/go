package main

import "fmt"

const MAX int = 3

//获取变量的内存地址
func pointerAdress() {
	var a int = 10

	fmt.Printf("变量的地址：%x\n", &a)
}

func varPointer() {
	var a int = 20
	var ip *int // 指针变量
	ip = &a     // 指针变量的地址

	fmt.Printf("a 变量的地址为%x\n", &a)

	fmt.Printf("ip 变量的指针地址：%x\n", ip)

	//使用指针访问值
	fmt.Printf("ip 变量的值%d\n", *ip)
}

//空指针
func emptyPointer() {
	var ptr *int
	fmt.Printf("ptr 的值为：%x\n", ptr)

	if ptr == nil {
		fmt.Println("ptr 是空指针")
	} else {
		fmt.Println("ptr 不是空指针")
	}

}

//指针数组
func arrPointer() {
	a := []int{1, 2, 3}
	var i int
	var ptr [MAX]*int

	for i = 0; i < MAX; i++ {
		ptr[i] = &a[i]
		fmt.Printf("a[%d] = %d\n", i, a[i])
	}

	for j := 0; j < MAX; j++ {
		fmt.Printf("a[%d] = %d\n", j, *ptr[j])
	}
}

//指向指针的指针
func ptrPointer() {
	var a int
	var ptr *int
	var pptr **int

	a = 7
	ptr = &a
	pptr = &ptr

	fmt.Printf("变量 a = %d\n", a)
	fmt.Printf("指针变量 *ptr = %d\n", *ptr)
	fmt.Printf("指向指针的指针变量 **pptr = %d\n", **pptr)
}

//向函数传递指针参数
func funPointer() {
	/* 定义局部变量 */
	var a int = 100
	var b int = 200

	fmt.Printf("交换前 a 的值 : %d\n", a)
	fmt.Printf("交换前 b 的值 : %d\n", b)

	/* 调用函数用于交换值
	 * &a 指向 a 变量的地址
	 * &b 指向 b 变量的地址
	 */
	swap(&a, &b)

	fmt.Printf("交换后 a 的值 : %d\n", a)
	fmt.Printf("交换后 b 的值 : %d\n", b)
}

func swap(x *int, y *int) {
	var temp int
	temp = *x
	*x = *y
	*y = temp
}

func main() {
	pointerAdress()
	varPointer()
	emptyPointer()
	arrPointer()
	ptrPointer()
	funPointer()
}
