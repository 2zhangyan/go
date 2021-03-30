package main

import "fmt"

//使用range去求一个slice的和
func rangeSum() {
	//声明一个切片
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)
	//在数组上使用range将传入index和值两个变量。上面那个例子不需要使用该元素的序号，所以使用空白符"_"省略了。

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	//range也可以用在map的键值对上。
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for i, c := range "go" {
		fmt.Println(i, c)
	}

}

func main() {
	rangeSum()

}
