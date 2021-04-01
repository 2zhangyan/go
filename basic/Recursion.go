package main

import "fmt"

//阶乘
func factorial(n uint64) (result uint64) {
	if n > 0 {
		result = n * factorial(n-1)
		return result
	}
	return 1
}

// 斐波那契数列
func fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return fibonacci(n-2) + fibonacci(n-1)

}

func main() {
	var i = 5
	fmt.Printf("%d 的阶乘是 %d\n", i, factorial(uint64(i)))

	for j := 0; j < 5; j++ {
		fmt.Printf("%d\t", fibonacci(j))
	}
}
