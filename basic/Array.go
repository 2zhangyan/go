package main

import "fmt"

func Arr() {
	var balance [10]float32
	fmt.Printf("数组：%f\n", balance)

	var price = [5]float32{1, 2, 3.6, 7, 8}
	fmt.Printf("价格：%f\n", price)

	var num = [...]float32{1.1, 2.2, 33.3}
	fmt.Printf("数字:%f\n", num)

	num2 := [5]float32{1: 2.0, 4: 2.70}
	fmt.Printf("结果%f\n", num2)
	num2[0] = 7.211111
	fmt.Printf("结果2%f\n", num2)

	var n [5]int
	var i, j int

	for i = 0; i < 5; i++ {
		n[i] = 7 + i
	}

	for j = 0; j < 5; j++ {
		fmt.Printf("Element[%d] = %d\n", j, n[j])
	}

}

//多维数据
func MArr() {
	var values [][]int
	row1 := []int{1, 2, 3}
	row2 := []int{4, 5, 6}
	values = append(values, row1)
	values = append(values, row2)
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[0][2])

	sites := [2][2]string{
		{"zz", "yy"},
		{"zy", "zy"},
	}
	fmt.Println(sites)

	animals := [][]string{}

	an_row1 := []string{"fish", "shark", "eel"}
	an_row2 := []string{"bird"}
	an_row3 := []string{"lizard", "salamander"}

	animals = append(animals, an_row1)
	animals = append(animals, an_row2)
	animals = append(animals, an_row3)

	for i := range animals {
		fmt.Printf("animals Row%v\n", i)
		fmt.Println(animals[i])
	}

}

//向函数传递数组
func getAvg(arr [5]int, size int) float32 {
	var i, sum int
	var avg float32

	for i = 0; i < size; i++ {
		sum += arr[i]
	}
	avg = float32(sum) / float32(size)
	return avg
}

func SendArr() {
	var balance = [5]int{100, 20, 70, 36, 5}
	var avg float32

	avg = getAvg(balance, 5)
	fmt.Printf("平均值%f", avg)
}

func main() {
	Arr()
	MArr()
	SendArr()
}
