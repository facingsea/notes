// Hello.go
package main

import (
	"fmt"
)

//func main() {
	/*fmt.Println("Hello World!")
	var result int = add(2, 3)
	fmt.Println(result)
	plus()
	x, y := swap("hello", "world")
	fmt.Println(x)
	fmt.Println(y)
	
	var a, b, c string = swap2("hello", "world", "go")
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)*/
	
	//e, f := split(9)
	//fmt.Println(e)
	//fmt.Println(f)
//}

func add(x int, y int) int {
	return x + y
}


func plus(){
	var i int = 1
	i += 2
	j := i + 3
	fmt.Println(j)
}

//函数返回多个值
func swap(x, y string) (string, string){
	return y, x
}

func swap2(x, y, z string) (string, string, string){
	return z, x, y
}

//返回值可以被指定变量名
func split(sum int) (x, y int){
	x = sum * 4 /9 
	y = sum - x
	return  //此处可以不用指定返回值了
}







