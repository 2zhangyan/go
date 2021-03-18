package main

import (
	"fmt"
	"reflect"
	"time"
)

//条件
//没有三目运算符，所以不支持 ?: 形式的条件判断。
func Condition() {
	a := 1
	AType := reflect.TypeOf(a)
	b := "aaa"
	Btype := reflect.TypeOf(b)

	fmt.Println(AType)
	fmt.Println(reflect.Int)

	if AType == reflect.TypeOf(int(0)) {
		fmt.Println("a is num")
	}
	if Btype == reflect.TypeOf(string("0")) {
		fmt.Println("b is string")
	}

	switch a {
	case 1:
		fmt.Printf("%d ,i am one", a)
		break
	case 2:
		fmt.Printf("%d ,i am two", a)
		break
	}

	//select 会循环检测条件，如果有满足则执行并退出，否则一直循环检测。
	ch := make(chan int)
	c := 0
	StopCh := make(chan bool)

	fmt.Printf("\n")
	fmt.Println(ch, StopCh)
	fmt.Printf("\n")

	go Chann(ch, StopCh)

	for {
		select {
		case c = <-ch:
			fmt.Println("Receive", c)
			fmt.Println("channe1")
		case s := <-ch:
			fmt.Println("Receive", s)
		case _ = <-StopCh:
			goto end
		}
	}
end:
}

func Chann(ch chan int, StopCh chan bool) {

	i := 10

	for j := 0; j < 10; j++ {
		ch <- i
		time.Sleep(time.Second)
	}
	StopCh <- true
}

// 循环
func Loop() {
	//for 循环
	for i := 1; i < 10; i++ {
		fmt.Println(i)
	}

	//乘法表
	for a := 1; a <= 9; a++ {
		for b := 1; b <= a; b++ {
			fmt.Printf("%d * %d = %d ", a, b, a*b)
		}
		fmt.Printf("\n")
	}
}

func main() {

	fmt.Println("条件--start")
	Condition()

	fmt.Println("循环--start")
	Loop()

}
