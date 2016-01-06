package main

import (
	"fmt"
	"math"
)

type AA struct {
	X, Y float64
}

// 结构体 AA 的方法
// 这里的方法接收者（method receiver）v 的类型为 *AA
func (v *AA) Abs() float64 {
	return math.Abs(v.X*v.X + v.Y*v.Y)
}

type Person struct {
	name string
	age  int
}

func (p *Person) sayHello1() { //p为指针类型，如果修改的话，会影响到原值
	p.name = "lisi"
}

func (p Person) sayHello2() { //p为Person类型的一个拷贝，修改p的值不会影响到原值
	p.name = "wangwu"
}

//自定义的基本类型不能有大括号
type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		fmt.Println("work")
		return float64(-f)
	}
	return float64(f)
}

/*
func main() {
	//v := &AA{3, 4}
	//fmt.Println(v.Abs())

	//	p1 := &Person{"zhangsan1", 20}
	//	fmt.Println(p1.name)
	//	p1.sayHello1()
	//	fmt.Println(p1.name)
	//	p1.sayHello2()
	//	fmt.Println(p1.name)

	//	p2 := Person{"zhangsan2", 34}
	//	fmt.Println(p2.name)
	//	p2.sayHello2()
	//	fmt.Println(p2.name)
	//	p2.sayHello1()
	//	fmt.Println(p2.name)

	f := MyFloat(-math.Sqrt2)
	fmt.Println(-math.Sqrt2)
	fmt.Println(f.Abs())
}
*/
