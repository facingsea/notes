package main

import (
	"fmt"
)

func sum(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	c <- sum
}

func calcSum() {
	a := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)
	//接收两个 goroutine 发送的计算结果
	x, y := <-c, <-c
	fmt.Println(x, y)
}

//channel 可以带有一个缓冲区（buffer）来缓存被传递的值，
//向 channel 中发送时只有缓冲区满的情况下会阻塞，接收 channel 中的值时只有在缓冲区空的情况下阻塞
func testBufferChannel() {
	c := make(chan int, 2)
	// 由于 channel 的缓冲区长度为 2
	// 因此发送不会阻塞
	c <- -1
	c <- -2
	fmt.Println(<-c, <-c)
	close(c)
	//c <- -3 //通道关闭之后不能再发送数据，否则会报panic: send on closed channel
}

/*
func main() {
	//calcSum()
	testBufferChannel()
}
*/
