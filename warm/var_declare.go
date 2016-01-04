// var_declare.go
package main

import (
	"fmt"
)


var java bool
var i int 
var i8 int8
var i16 int16
var i32 int32
var i64 int64
var b byte
var b8 uint8
var b16 uint16
var b32 uint32
var b64 uint64
var f32 float32
var f64 float64
var c64 complex64 //复数
var c128 complex128 //复数

/*
func main() {
	//defaultValue()
	//testInt()
	//testDeclear()
	testConst()
	//testCastType()
}
*/

//输出默认值
func defaultValue(){
	fmt.Println(java)
	fmt.Println(i)
	fmt.Println(i8)
	fmt.Println(i16)
	fmt.Println(i32)
	fmt.Println(i64)
	fmt.Println(b)
	fmt.Println(b8)
	fmt.Println(b16)
	fmt.Println(b32)
	fmt.Println(b64)
	fmt.Println(f32)
	fmt.Println(f64)
	fmt.Println(c64)
	fmt.Println(c128)
}

func testInt(){
	var x, y int = 1, 2
	var i, j = true, "hello"
	
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(i)
	fmt.Println(j)
}

func testDeclear(){
	x, y := 3, "hello"
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(string(x) + y)
	fmt.Println(string(x))
	//fmt.Println(float64(x))
}

const CODE = 123
const (
	Big int64 = 1 << 10
	Small int64 = Big >> 9
)
func testConst(){
	const Pi = 3.14
	fmt.Println(Pi)
	fmt.Println(CODE)
	fmt.Println(Big)
	fmt.Println(Small)
}

//类型转换
func testCastType(){
	x := 23.7
	y := int8(x)
	z := int16(x)
	fmt.Println(y)
	fmt.Println(z)
	//m := complex64(x)
	//fmt.Println(m)
}
