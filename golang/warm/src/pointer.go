package main

import (
	"fmt"
)

//refer to sign.go

type HH struct {
	X int
	Y int
	Z float64
}

func testFmt() {
	fmt.Println("test")
}

/*
//所谓的指针
func main(){
	p := HH{2, 3, 34.5}
	q := &p
	fmt.Println(q.X)
	q.X = 1
	fmt.Println(q.X)
}
*/
