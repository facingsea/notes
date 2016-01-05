package main	

import (
	"fmt"
)

//map 用于映射 key 到 value（值）。map 可以通过 make 来创建（而非 new）

type MM struct {
	name string
	age int
}

func testMap(){
	m := make(map[string]MM)
	m["Beijing"] = MM{
		"zhangsan", 20, 
	}
	fmt.Println(m)
	fmt.Println(m["Beijing"])
	fmt.Println(m["beijing"])
	
	n := map[string]MM{
		// Type(MM)可以省略
		"shanghai":MM{
			"lisi", 23,
		},
		"guangzhou":{
			"wangwu", 34,
		},
	}
	fmt.Println(n)
	fmt.Println(n["shanghai"])
	fmt.Println(n["guangzhou"])
	fmt.Println(len(n))
}

//delete 
func testMap2(){
	m := map[string]MM{
		"chongqing":MM{
			"ha", 23,
		}, 
		"xian":{
			"he", 20,
		},
	}
	fmt.Println(m)
	delete(m, "chongqing")
	fmt.Println(m)
}

// 检查map中是否有某个key
//	elem, ok := m[key]
//  elem表示该key在m中所对应的值
//	ok表示该m中是否有该key对应的值
func testMap3(){
	m := make(map[string]MM)
	m["beijing"] = MM{
		"hello", 34,
	}
	m["zhengzhou"] = MM{
		"hi", 23,
	}
	fmt.Println(m)
	a1, b1 := m["shanghai"]
	fmt.Println(a1) //out is { 0}
	fmt.Println(b1) //out is false
	
	a2, b2 := m["beijing"]
	fmt.Println(a2) //out is {hello 34}
	fmt.Println(b2) //out is true
}

var aMap map[string]MM

func testNilMap(){
	if aMap == nil { // aMap is nil
		fmt.Println("aMap is nil")
	}else{
		fmt.Println("aMap is not nil")
	}
	
	bMap := make(map[string]MM)
	if bMap == nil { // bMap is not nil
		fmt.Println("bMap is nil")
	}else {
		fmt.Println("bMap is not nil")
	}
}

/*
func main(){
	//testMap()
	//testMap2()
	//testMap3()
	testNilMap()
}
*/
