// foreach.go
package main

import (
	"fmt"
	"math/rand"
)


/*
func main() {
//	testFor()
	//testIfElse()
	aa := dijkstra(9)
	fmt.Println("The val is ", aa)
}
*/

func dijkstra(ha int ) (aa int) {
	i := 1
	for j := 0; j < 5; j++ {
		ha = ha - i
	}
	return ha
}

func testFor(){
	sum := 0
	for i:=0; i < 4; i++{
		sum += i
	}
	fmt.Println("The result is : ", sum)
	
	for sum < 1000 {
		sum += sum
	}
	fmt.Println("the result is : ", sum)
}

func testIfElse(){
	if n := rand.Intn(6); n <= 2 {
		fmt.Println("[0, 2]", n)
	}else{
		fmt.Println("[3, 5]", n)
	}
	//fmt.Println(n)
	
	//for{
	//	a := rand.Intn(6)
	//	fmt.Println(a)
	//}
}
