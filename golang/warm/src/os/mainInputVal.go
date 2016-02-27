package main

import (
	"fmt"
	"os"
	"strings"
)

//测试main方法获取输入参数

//init方法先于main方法执行
func init() {
	fmt.Println("init method run.")
}

//用命令行的方式执行：go run mainInputVal.go "zhangsan"
func testMain() {
	who := "World!"
	fmt.Println(os.Args[0]) // out is E:\code\coding.net\golang\warm\os\os.exe，为改程序编译的名称
	if len(os.Args) > 1 {
		who = strings.Join(os.Args[1:], " ")
	}
	fmt.Println("Hello ", who)
}

/*
func main() {
	testMain()
}
*/
