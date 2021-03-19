package main

import "fmt"

func Arr() {
	var balance [10]float32
	fmt.Printf("数组：%f", balance)
	fmt.Println("\n")

	var price = [5]float32{1, 2, 3.6, 7, 8}
	fmt.Printf("价格：%f", price)

	fmt.Println("\n")
	var num = [...]float32{1.1, 2.2, 33.3}
	fmt.Printf("数字:%f", num)

}

func main() {
	Arr()
}
