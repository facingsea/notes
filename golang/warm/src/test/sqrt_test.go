package newmath

import (
	"fmt"
	"testing"
)

//使用go test 方法执行
//测试文件中包含的所有测试函数 func TestXxx(t *testing.T) 都会被执行
//测试方法
// Test + 方法名（方法名首字母必须大写，不然找不到测试方法，原方法首字母也要大写）
func TestCalcInt(t *testing.T) {
	const inX, inY, out = 1, 2, 3
	if sum := CalcInt(inX, inY); sum != out {
		t.Errorf("%x + %y should be %o, not %s", inX, inY, out, sum)
	} else {
		t.Logf("Right !")
	}

}

//这里 Example 函数结尾可以跟上一个以 “Output:” 字符串开始的注释（被叫做输出注释），
//在测试运行时会将 Example 函数输出注释中字符串和函数标准输出进行比对。
//需要注意的是，如果没有输出注释，Example 函数是不会被运行的（但是会被编译）
func ExampleCalcInt() {
	//fmt.Println("Example func run")
	sum := CalcInt(2, 3)
	fmt.Println(sum)
	// Output: 6  // 会将6与上面fmt输出的结果做对比，如果符合则通过，否则失败
}

//以上方法输入结果为：

//=== RUN   TestCalcInt
//--- PASS: TestCalcInt (0.00s)
//	sqrt_test.go:17: Right !
//=== RUN   ExampleCalcInt
//--- FAIL: ExampleCalcInt (0.00s)
//got:
//5
//want:
//6
//FAIL
//exit status 1
//FAIL	_/E_/code/coding.net/golang/warm/test	0.220s
//错误: 进程退出代码 1.
