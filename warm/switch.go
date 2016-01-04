// switch.go
package main

import (
	"fmt"
	"runtime"
	"math/rand"
	"time"
)

/*
func main() {
	//testSwitch()
	testNoConditionSwitch()
}
*/

func testSwitch(){
	fmt.Println("Go runs on ")
	switch os := runtime.GOOS; os{
		case "darwin": 
			fmt.Println("OS X")
		case "linux": 
			fmt.Println("Linux")
		default : 
			fmt.Println("%s.", os)
	}
	
	switch m := rand.Intn(5); m {
		case 1 :
			fmt.Println("first", m)
		case 2 : 
			fmt.Println("second", m)
		case 3 : 
			fmt.Println("third", m)
		default : 
			fmt.Println("Other ", m)
	}
}

//switch可以没有条件
func testNoConditionSwitch(){
	t := time.Now()
	switch{
		case t.Hour() < 12 :
			fmt.Println("Good morning!")
		case t.Hour() < 17 :
			fmt.Println("Good afternoon!")
		default : 
			fmt.Println("Good evening!")
	}
}
