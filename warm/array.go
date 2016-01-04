package main

import (
	"fmt"
)

//[n]T 在 Golang 中是一个类型（就像 *T 一样），表示一个长度为 n 的数组其元素类型为 T。
func testArray() {
	var a [2]string
	a[0] = "hello"
	a[1] = "world"
	fmt.Println(a[0], a[1])
	fmt.Println(a)
	//a[3] = "!" //数组的长度不能改变，此处错误
	//fmt.Println(a)
}

/*
func main() {
	testArray()
}
*/
