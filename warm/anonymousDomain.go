package main

import (
	"fmt"
)

//匿名域
//结构体中可以存在只有类型而没有名字的域，它们被叫做匿名域。
//struct {
//    T1
//    *T2
//}
//一个结构体的匿名域中的域或者方法可以被此结构体实例直接访问

type Car struct {
	wheelCount int
}

func (c *Car) numberOfWheels() int {
	fmt.Println("method")
	return c.wheelCount
}

type BMW struct {
	Car
}

type Benz struct {
	*Car
}

/*
func main() {
	bmw := BMW{Car{4}}
	//bmw实例可以直接访问匿名域Car里的方法
	fmt.Println(bmw.numberOfWheels()) //out is 4
	bmw2 := &BMW{Car{5}}
	fmt.Println(bmw2.numberOfWheels()) //out is 5

	benz := &Benz{&Car{6}}
	fmt.Println(benz.numberOfWheels())
}
*/
