package main	

import (
	"fmt"
)

//区分符号&与*
func testIdentify(){
	var a int = 1 
    var b *int = &a
    var c **int = &b
    var x int = *b
    fmt.Println("a = ",a)
    fmt.Println("&a = ",&a)
    fmt.Println("*&a = ",*&a)
    fmt.Println("b = ",b)
    fmt.Println("&b = ",&b)
    fmt.Println("*&b = ",*&b)
    fmt.Println("*b = ",*b)
    fmt.Println("c = ",c)
    fmt.Println("*c = ",*c)
    fmt.Println("&c = ",&c)
    fmt.Println("*&c = ",*&c)
    fmt.Println("**c = ",**c)
    fmt.Println("***&*&*&*&c = ",***&*&*&*&*&c)
    fmt.Println("x = ",x)
}

func main() {
    
}