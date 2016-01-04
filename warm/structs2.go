package main	

import (
	"fmt"
)

type DESK struct {
	HEIGHT int
	LENGTH int
	AREA float64
}

func testFmt2() {
	fmt.Println("test")
}

/*
func main(){
	a := DESK{}
	fmt.Println(a.HEIGHT)
	fmt.Println(a.LENGTH)
	fmt.Println(a.AREA)
	
	d := DESK{2, 3, 23.6}
	fmt.Println(d.HEIGHT)
	fmt.Println(d.LENGTH)
	fmt.Println(d.AREA)
	//重新构造一个新的结构体
	v := DESK{}
	fmt.Println(v.AREA)
	fmt.Println(v.LENGTH)
	fmt.Println(v.HEIGHT)
	v.AREA = 45.7
	fmt.Println(v.AREA)
}
*/