package loop_test

import (
	"testing"
)

// 循环 go 只支持 for 循环并且没用（）
//
/*
while 条件循环的写法
	n:=0
	for n<5{
		n++
		fmt.Println(n)
	}
	无限循环的写法
	for {
		....
	}
*/

func TestWhileLoop(t *testing.T) {
	n := 0
	for n < 5 {
		t.Log(n)
		n++
	}
}

// if 条件
// 1.condition 表达式结果必须为布尔值
// 2.支持变量赋值
/**
if var declaration;condition{
	...

}
*/

func TestIfMultiSec(t *testing.T) {
	if a := 1 == 1; a {
		t.Log("1==1")
	}
	// 调用方法的返回值
	// if v, err := someFunc(); err == nil {
	// 	t.Log("dd")
	// } else {
	// 	t.Log("ee")
	// }
}

// switch 条件
// 条件表达式不限制为常量或者整数
// 单个case中可与你出现多个结果选项，使用逗号，分割
// 与c语言等规则相反，Go语言不需要用break来证明推出一个 case
// 可以不设定 switch 之后的条件表达式，在此情况下，整个 switch 结构与 多个if。。。else..的逻辑相同

func TestSwitchCase(t *testing.T) {

	for i := 0; i < 5; i++ {
		switch i {
		case 0, 2:
			t.Log("偶数")
		case 1, 3:
			t.Log("奇数")
		default:
			t.Log("不是偶数不是奇数")
		}
	}

	for i := 0; i < 5; i++ {
		switch {
		case i%2 == 0:
			t.Log("偶数")
		case i%1 == 1:
			t.Log("奇数")
		default:
			t.Log("不是偶数不是奇数")
		}
	}
}
