package main	

import (
	"fmt"
)

//range 被用在 for 中来迭代一个 slice 或者一个 map：

var s = []int{1, 3, 5}

func testRange(){
	for i, v := range s{
		fmt.Println(i, v)
	}
	
	//如果只需要值，可以用_忽略索引
	for _, v := range s{
		fmt.Println(v)
	}
	
	//只需要索引
	for i := range s{
		fmt.Println(i)
	}
}

/*
func main(){
	testRange()
}
*/