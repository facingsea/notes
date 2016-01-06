package main

import (
	"fmt"
	"math"
)

//定义接口
type Abser interface {
	Abs2() float64
}

//定义结构体
type II struct {
	X, Y float64
}

//实现方法Abs2
func (i *II) Abs2() float64 {
	fmt.Println(i)
	return math.Sqrt(i.X*i.X + i.Y*i.Y)
}

/*
func main() {
	i := II{3, 4}
	fmt.Println(i)
	//正确，在*II实现了Abser接口
	var a Abser = &i
	fmt.Println(a.Abs2())

	//var b Abser = i //错误，II does not implement Abser (Abs2 method has pointer receiver)

}
*/
