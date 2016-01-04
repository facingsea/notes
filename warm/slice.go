package main

import (
	"fmt"
)

//创建切换的方式有四种：
//	make([]Type, length, capacity) ：其长度即为切片的容量
//	make([]Type, length) ：其长度即为切片的长度
// 	[]Type{} ：复合语法，其长度为大括号内项的个数
//	[]Type{value1, value2, ..., valueN} ：复合语法，其长度为大括号内项的个数

//slice 是一个数据结构，其指向一个数组某个连续的部分。slice 用起来很像数组。[]T 为 slice 类型，其中元素类型为 T
func testSclice() {
	p := []int{2, 3, 5, 11}
	fmt.Println("p == ", p)

	for i := 0; i < len(p); i++ {
		fmt.Printf("p[%d] == %d\n", i, p[i])
	}
}

func testSclice2() {
	s := []string{"A", "B", "C", "D", "E", "F", "G"}
	t := s[:5]
	u := s[3 : len(s)-1]
	//out is : [A B C D E F G] [A B C D E] [D E F]
	fmt.Println(s, t, u)
	u[1] = "X"
	//out is : [A B C D X F G] [A B C D X] [D X F]
	fmt.Println(s, t, u)

}

func main() {
	//testSclice()
	testSclice2()
}
