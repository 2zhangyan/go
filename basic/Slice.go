package main

import "fmt"

func definitionSlice() {
	//定义切片
	//var silce []type = make([]type,len)
	//slice := make([]type,len)
	//make([]T, length, capacity) // capacity 指定容量， capacity 为可选参数。

	//初始化
	//s := []int {1,2,3} //[] 表示是切片类型，{1,2,3} 初始化值依次是 1,2,3，其 cap=len=3。

	//len() 和 cap() 函数
	//切片是可索引的，并且可以由 len() 方法获取长度。
	//切片提供了计算容量的方法 cap() 可以测量切片最长可以达到多少。

	var numbers = make([]int, 3, 4)
	printSlice(numbers)

	var numbers2 []int //定义空切片
	printSlice(numbers2)

}

//切片截取
func interceptSlice() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7}
	printSlice(numbers)

	fmt.Println("numbers == ", numbers)

	fmt.Println("numbers[1:4] == ", numbers[1:4])

	fmt.Println("numbers[:2] == ", numbers[:2])

	fmt.Println("numbers[2:] == ", numbers[2:])

	numbers1 := make([]int, 0, 5)
	printSlice(numbers1)

	numbers2 := numbers[:2]
	printSlice(numbers2)

	// 不包含索引4
	numbers3 := numbers[1:4]
	printSlice(numbers3)

}

//append() 和 copy() 函数
func acSlice() {
	var numbers []int
	printSlice(numbers)

	//允许追加空切片
	numbers = append(numbers, 0)
	printSlice(numbers)

	numbers = append(numbers, 1)
	printSlice(numbers)

	numbers = append(numbers, 1, 2, 3, 4)
	printSlice(numbers)

	numbers1 := make([]int, len(numbers), (cap(numbers))*3)
	copy(numbers1, numbers)
	printSlice(numbers1)
}

func printSlice(x []int) {
	fmt.Printf("len = %d ;cap = %d; slice=%v\n", len(x), cap(x), x)
}

//切片
func main() {
	//definitionSlice()
	//interceptSlice()
	acSlice()
}
