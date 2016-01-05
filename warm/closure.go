package main	

import (
	"fmt"
)

//Golang 中函数也是一个值（就像 int 值一样），且函数可以是一个闭包。闭包是一个引用了外部变量的函数。

//返回的是一个函数
func adder() func(int) int{
	sum := 0
	
	// 返回一个闭包，此闭包引用了外部变量 sum
	return func(x int) int{
		sum += x
		fmt.Println(sum)
		return sum
	}
}

/*
func main(){
	a := adder()
	fmt.Println(a)
	fmt.Println(a(2))
}
*/