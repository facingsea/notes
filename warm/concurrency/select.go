package main

import (
	"fmt"
	"math/rand"
	//"time"
)

// select语法如下：
// select {
//	 case sendOrReceive1: block1
//	 ...
//	 case sendOrReceiveN: blockN
//	 default: blockD
// }

func createChannels() {
	channels := make([]chan bool, 6) // 创建一个切片
	fmt.Println(channels)            // out is [<nil> <nil> <nil> <nil> <nil> <nil>]
	for i := range channels {
		channels[i] = make(chan bool)
	}
	fmt.Println(channels)
	go func() {
		fmt.Println("func run.")
		for {
			index := rand.Intn(6)
			fmt.Println("begin index ", index)
			channels[index] <- true
			fmt.Println("end index ", index)
		}
	}()

	for i := 0; i < 3; i++ {
		var x int
		fmt.Println("step", i)
		select {
		case <-channels[0]:
			x = 1
		case <-channels[1]:
			x = 2
		case <-channels[2]:
			x = 3
		case <-channels[3]:
			x = 4
		case <-channels[4]:
			x = 5
		case <-channels[5]:
			x = 6
		}
		fmt.Printf("out channel : %d", x)
		fmt.Println()
	}
	fmt.Println()
}

func expensiveComputation(area int, answer chan int, done chan bool) {
	finished := false
	var i = 0
	for !finished {
		result := rand.Intn(area)
		//time.Sleep(10 * time.Second)
		answer <- result
		i++
		if i > 5 {
			break
		}
	}
	done <- true
}

func calc() {
	const allDone = 2
	doneCount := 0
	answera := make(chan int)
	answerb := make(chan int)
	fmt.Println(answera) // out is 0x11fc2140, 通道返回的也是指针
	defer func() {
		close(answera)
		close(answerb)
	}()
	done := make(chan bool)
	defer func() {
		close(done)
	}()
	go expensiveComputation(10, answera, done)
	go expensiveComputation(5, answerb, done)
	for doneCount != allDone {
		var which, result int
		//在一个select语句中，GO语言会按照顺序从头至尾评估每一个发送和接收语句，
		//如果其中的任意一语句可以继续执行（即没有阻塞），那么就从那些可以执行的语句中 任意 选择
		//一条来使用。
		//如果没有任意一条可以执行（即所有通道都被阻塞），那么有两种可能的情况，如果给出了default
		//语句，那么就会执行default语句，同时程序的执行会从select语句后的语句中恢复；如果没有default语句
		//那么select将会被阻塞，直到至少有一个通信可以继续执行下去
		select {
		case result = <-answera:
			which = 'a'
		case result = <-answerb:
			which = 'b'
		case <-done:
			doneCount++
		}
		if which != 0 {
			fmt.Printf("%c --> %d ", which, result)
			fmt.Println()
		}
	}
	fmt.Println()
}

func testChar() {
	var aa = 'c'
	fmt.Println(aa) // out is 99
}

/*
func main() {
	//	createChannels()
	//	time.Sleep(1000)
	//	fmt.Println("come on")
	//	time.Sleep(1000)

	calc()
	//testChar()
}
*/
