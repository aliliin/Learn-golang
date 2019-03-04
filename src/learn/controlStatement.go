package learn

import (
	"fmt"
	"unsafe"
)

// 控制语句学习
func init() {

}

// goto label;
func LearnGoto() {
	var num3 = 3.7E+1 + 5.98E-2i
	fmt.Print(num3)
	fmt.Print(unsafe.Sizeof(num3))
	goto one
	fmt.Print("这里是中间代码块 1 \n")
one:
	fmt.Print("这里是代码块 1 \n")

	var a int = 10
	/* 循环 */
LOOP:
	for a < 20 {
		if a == 15 {
			/* 跳过迭代 */
			a = a + 1
			goto LOOP
		}
		// 跳出循环
		if a == 18 {
			break
		}
		fmt.Printf("a的值为 : %d\n", a)
		a++
	}

}

// for init; condition; post { }
// init： 一般为赋值表达式，给控制变量赋初值；
// condition： 关系表达式或逻辑表达式，循环控制条件；
// post： 一般为赋值表达式，给控制变量增量或减量。
func LearnFor() {
	// 无限循环
	//for {
	//	fmt.Print("无限循环")
	//}

	for i := 1; i < 3; i++ {
		fmt.Print(i)
	}
	//foreach 数组用法 配合range
	array := []string{"高", "永", "立"}
	for key, value := range array {
		fmt.Print(key)
		fmt.Print(value)
	}
	//foreach 数组用法 不需要key 的用法
	for _, value := range array {
		fmt.Print(value)
	}
}

// switch 语句用于基于不同条件执行不同动作，每一个case 分支都是唯一的，从上往下逐一测试，知道匹配位置
// switch 语句执行的过程从上到下
// switch 语句还可以被用于 type-switch 来判断某个 interface 变量中实际存储的变量类型
func LearnSwitch() {

	var a interface{}
	a = 32
	a = "Aliliin"
	a = 3.143
	switch a.(type) {
	case int:
		fmt.Print("类型为int\n")
	case string:
		fmt.Print("类型为字符串\n")
	default:
		fmt.Print("未知类型 \n")
	}

	var grade string = "D"
	switch {
	case grade == "A":
		fmt.Printf("优秀!\n")
	case grade == "B", grade == "C":
		fmt.Printf("良好\n")
	case grade == "D":
		fmt.Printf("及格\n")
	case grade == "F":
		fmt.Printf("不及格\n")
	default:
		fmt.Printf("差\n")
	}
	fmt.Printf("你的成绩是 %s\n", grade)

	switch 4 {
	case 1:
		fmt.Print("判断为1")
	case 2:
		fmt.Print("判断为2")
	case 3:
		fmt.Print("判断为3")
	default:
		fmt.Print("以上都不满足")
	}
}
func LearnIf() {
	a := 100
	b := 10
	if a > b {
		fmt.Print("a大于等于b \n")
		if a > 20 {
			fmt.Print("a大于于20 \n")
		}
	} else {
		fmt.Print("a小于b \n")
	}
	d := 1
	if d == 2 {
		fmt.Print("等于2")
	} else {
		fmt.Print("不等于2")
	}
	age := 28
	if age == 26 {
		fmt.Print("age是26")
	} else if age < 27 {
		fmt.Print("age是27")
	} else {
		fmt.Print("age大于27") // age大于27
	}
}
