// structs.go
package main

import (
	"fmt"
)

type Vertex struct {
	X int
	Y int
}

type Ha struct{
	X, Y int
}

func testStruct(){
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v)
}

//指针的形式
func testPointer(){
	p := Vertex{1, 2}
	q := &p
	q.X = 2
	fmt.Println(p)
	fmt.Println(q)
}

func testSetValue(){
	a := Ha{X: 3}
	fmt.Println(a)
	a = Ha{Y:4, X:1}
	fmt.Println(a)
}

/*
func main() {
	//testStruct()
	//testPointer()
	testSetValue()
}
*/
