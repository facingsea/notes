// defer.go
package main

import (
	"io"
	"fmt"
	"os"
)

/*
func main() {
	//testCopyFile("E:/111.sql", "D:/data.sql")
	testDefer()
}
*/


func testCopyFile(dstName, srcName string) (written int64, err error){
	src, err := os.OpenFile(srcName, os.O_RDONLY, os.ModeAppend)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer src.Close()
	
	dst, err := os.Create(dstName)
	if err != nil {
		fmt.Println(err.Error())
		return 
	}
	defer dst.Close()
	
	return io.Copy(dst, src)
}

//一个 defer 语句能够将一个函数调用加入一个列表中（这个函数调用被叫做 deferred 函数调用），
//在当前函数调用结束时调用列表中的函数。
//类似一个栈：先进后出
func testDefer(){
	for i := 0; i < 5; i ++ {
		defer fmt.Print(i)
	}
	fmt.Println("over")
}

