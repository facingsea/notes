// new.go
package main

import (
	"fmt"
)

type NEWC struct {
	X, Y int
}

func testNew() {
	v := new(NEWC)
	fmt.Println(v)
	v.X, v.Y = 11, 9
	fmt.Println(v)
	//fmt.Println(NEWC)
}


func main() {
	testNew()
}
