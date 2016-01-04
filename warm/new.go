// new.go
package main

import (
	"fmt"
)

type NEWC struct {
	X, Y int
}

//我们可以通过表达式 new(T) 分配一个被初始化为 0 且类型为 T 的值，并且返回指向此值的指针
func testNew() {
	v := new(NEWC) //返回的是指针
	fmt.Println(v)
	v.X, v.Y = 11, 9
	fmt.Println(v)
	//fmt.Println(NEWC) //结构体不能直接输出
	fmt.Println(v.X)
	fmt.Println(v.Y)
}

/*
func main() {
	testNew()
}
*/
