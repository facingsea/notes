package main

import (
	"fmt"
)

//区分符号&与*
func testIdentify() {
	var a int = 1
	var b *int = &a
	var c **int = &b
	var x int = *b
	fmt.Println("a = ", a)
	fmt.Println("&a = ", &a)
	fmt.Println("*&a = ", *&a)
	fmt.Println("b = ", b)
	fmt.Println("&b = ", &b)
	fmt.Println("*&b = ", *&b)
	fmt.Println("*b = ", *b)
	fmt.Println("c = ", c)
	fmt.Println("*c = ", *c)
	fmt.Println("&c = ", &c)
	fmt.Println("*&c = ", *&c)
	fmt.Println("**c = ", **c)
	fmt.Println("***&*&*&*&c = ", ***&*&*&*&*&c)
	fmt.Println("x = ", x)
}

//如果是取了一次地址，变量类型为*int，取两次，变量类型是**int
func testIdentify2() {
	var a int = 1
	var b *int = &a
	var c **int = &b //c的类型是**int
	fmt.Println("a: ", a)
	fmt.Println("b: ", b)
	fmt.Println("c: ", c)
	var d *int = &a
	fmt.Println("d: ", d)
}

/*
func main() {
	//testIdentify()
	testIdentify2()
}
*/
