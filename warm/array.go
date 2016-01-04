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
	fmt.Println("len is ", len(a))
}

//多维数组
func testMultidimensionalArray() {
	var buffer [20]byte
	var grid1 [3][3]int
	grid1[1][0], grid1[1][1], grid1[1][2] = 8, 6, 2
	grid2 := [3][3]int{{4, 3}, {5, 8, 10}}
	cities := [...]string{"shanghai", "Mumbai", "New York", "California"}
	cities[len(cities)-1] = "Las Vegas"
	fmt.Println("Type Len Contents ")
	fmt.Printf("%-8T %2d %v\n", buffer, len(buffer), buffer)
	fmt.Printf("%-8T %2d %q\n", cities, len(cities), cities)
	fmt.Printf("%-8T %2d %v\n", grid1, len(grid1), grid1)
	fmt.Printf("%-8T %2d %v\n", grid2, len(grid2), grid2)
}

/*
func main() {
	//testArray()
	testMultidimensionalArray()
}
*/
