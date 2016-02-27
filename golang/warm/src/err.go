package main

import (
	"errors"
	"fmt"
	//"os"
)

//golang提供的error接口
//type error interface {
//    Error() string
//}

//创建一个 error 值的最简单方式是使用 errors.New 函数：
func MakeErrorTest(f float64) (float64, error) {
	if f < 0 {
		fmt.Println(f)
		return 0, errors.New("f must be a positive number.")
	}
	return 1, nil
}

//自定义一个error的实现
type NegativeNumberError float64

func (n NegativeNumberError) Error() string {
	return fmt.Sprintf("math: the %g can not be a negative number.", n)
}

/*
func main() {
	//	_, err := os.Open("test.txt")
	//	if err != nil {
	//		fmt.Println(err)
	//	}

	//	val, err := MakeErrorTest(-1)
	//	fmt.Println(val)
	//	fmt.Println(err)

	var a = NegativeNumberError(-1)
	fmt.Println(a.Error())
}
*/
