package newmath

import "testing"

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
