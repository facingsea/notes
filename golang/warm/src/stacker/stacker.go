package main

import (
	"fmt"
	"stacker/stack"
)

func testMyStack() {
	var myStack stack.Stack
	myStack.Push("hello")
	myStack.Push(-12)
	myStack.Push([]string{"zhangsan", "lisi", "wangwu"})
	myStack.Push(12.25)
	for {
		item, err := myStack.Pop()
		if err != nil {
			break
		}
		fmt.Println(item)
	}
}

/*
func main() {
	testMyStack()
}
*/
