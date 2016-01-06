package main

import (
	"fmt"
	"time"
)

// goroutine 形式

func say(s string, t int) {
	fmt.Println("begin..", t)
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s, t)
	}
	fmt.Println("end...", t)
}

/*
func main() {
	//开启一个 goroutine 执行 say 函数
	go say("hello", 1)
	//普通形式执行
	say("world", 2)
}
*/
