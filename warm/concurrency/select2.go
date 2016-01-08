package main

import (
	"fmt"
	"time"
)

func fibonacci2(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func testFibonacci2() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0 //通过quit通道结束fibonacci2方法中的无限循环
	}()
	fibonacci2(c, quit)
}

//select 中的 default 会在没有任何 case 可执行时执行（类似于 switch）
func testSelectDefault() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick")
		case <-boom:
			fmt.Println("Boom")
			return
		default:
			fmt.Println("  .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}

/*
func main() {
	//testFibonacci2()
	testSelectDefault()
}
*/
