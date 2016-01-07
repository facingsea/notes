package main

import (
	"fmt"
)

//发送和接收的语法：
// channel <- value     //阻塞发送
// <- channel           //接收并将其丢弃
// x := <- channel      //接收并用x保存
// x, ok := <- channel  //功能同上，同时检查通道是否关闭或者是否为空

//通道可以使用内置的make()函数通过以下语法来创建：
// make(chan Type)
// make(chan Type, capacity)  //capacity为缓冲容量

func createCounter(start int) chan int {
	//该通道是同步的，因此会阻塞直到发送者准备好发送 或 接收者准备接收
	//由于通道创建时容量为0，因此该发送回阻塞直到收到一个从通道中接收数据的请求
	//该阻塞只会影响到此例中的匿名函数所在的goroutine
	next := make(chan int)
	go func(i int) {
		for {
			fmt.Println("i: ", i)
			next <- i
			fmt.Println("get data from channel.")
			i++
		}
	}(start)
	return next
}

func execute() {
	counterA := createCounter(2)
	counterB := createCounter(102)
	for i := 0; i < 5; i++ {
		a := <-counterA
		fmt.Printf("(A->%d, B->%d)", a, <-counterB)
	}
	fmt.Println()
}

/*
func main() {
	execute()
}
*/
