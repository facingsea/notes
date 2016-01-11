package main

import "fmt"

//在Go语言中，所有其它数据类型都实现了空接口。
func testDeclear() {
	var v1 interface{} = 1
	var v2 interface{} = "hello"
	var v3 interface{} = struct { X int }{ 1 }
	fmt.Println(v1)
	fmt.Println(v2)
	fmt.Println(v3)
}

//如果函数打算接收任何数据类型，则可以将参考声明为interface{}。
//最典型的例子就是标准库fmt包中的Print和Fprint系列的函数
func testInterface(){
	var user struct{Name string; Gender int}
    user.Name = "dotcoo"
	myPrintln(user)
	myPrintln("hello", " ", 25, user)
	//测试匿名结构体
	myPrintln(struct { X int }{ 1 })
}

//
func myPrintln(a ...interface{}){
	fmt.Println(a)
}

/*
func main() {
	//testDeclear()
	testInterface()
}
*/