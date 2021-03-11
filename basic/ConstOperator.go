package main

import "fmt"

func one() {
	const TestName string = "Test"
	const TestAge int = 18
	const TestRegion, TestAddress = "中国", "北京"
	fmt.Println(TestName, TestAge, TestRegion, TestAddress)
}

func main() {
	one()
}
