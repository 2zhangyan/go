package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

//以一个不同的、新创建的 goroutine 来执行一个函数
func runSay() {
	go say("word")
	say("hello")
}

/**
通道
注意：默认情况下，通道是不带缓冲区的。发送端发送数据，同时必须有接收端相应的接收数据。
**/
func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		fmt.Println("aaaa", v)
		sum += v
	}
	c <- sum //把sum发送到通道c
	fmt.Println("bbbb", sum)
}

func cSum() {
	s := []int{1, 2, 3, 4, 5, 6, -8}
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // 从通道c中接收
	fmt.Println(x, y, x+y)
}

//通道缓冲区
//如果通道不带缓冲，发送方会阻塞直到接收方从通道中接收了值。如果通道带缓冲，发送方则会阻塞直到发送的值被拷贝到缓冲区内；
//如果缓冲区已满，则意味着需要等待直到某个接收方获取到一个值。接收方在有值可以接收之前会一直阻塞。
func cNum() {
	ch := make(chan int, 2) //缓冲区大小为2

	ch <- 3
	ch <- 2

	fmt.Println(<-ch)
	fmt.Println(<-ch)

}

//遍历通道与关闭通道
func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func cClosNum() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	fmt.Println("cccc")
	// range 函数遍历每个从通道接收到的数据，因为 c 在发送完 10 个
	// 数据之后就关闭了通道，所以这里我们 range 函数在接收到 10 个数据
	// 之后就结束了。如果上面的 c 通道不关闭，那么 range 函数就不
	// 会结束，从而在接收第 11 个数据的时候就阻塞了。
	for i := range c {
		fmt.Println(i)
	}
}

func main() {
	runSay()
	cSum()
	cNum()
	cClosNum()
}
