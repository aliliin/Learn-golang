package loop_test

import "testing"

// 如果case 条件达成就会跳出循环，想继续执行下面的代码可以使用 fallthroug 强制执行后面的case代码
func TestSwitch(t *testing.T) {
	integer := 6
	switch integer {
	case 4:
		t.Log("The integer was <=4")
		fallthrough
	case 5:
		t.Log("The integer was <=5")
		fallthrough
	case 6:
		t.Log("The integer was <=6")
		fallthrough
	case 7:
		t.Log("The integer was <=7")
		fallthrough
	default:
		t.Log("default Case")
	}
}
