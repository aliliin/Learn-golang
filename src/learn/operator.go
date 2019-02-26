package learn

import "fmt"

func init() {
	
}
// 位运算符
func LearnBitwise(){
	var a uint = 60    /* 60 = 0011 1100 */
	var b uint = 13    /* 13 = 0000 1101 */
	var c uint = 0
	c = a & b       /* 12 = 0000 1100 */
	fmt.Printf("第一行 - c 的值为 %d\n", c )
	c = a | b       /* 61 = 0011 1101 */
	fmt.Printf("第二行 - c 的值为 %d\n", c )

	c = a ^ b       /* 49 = 0011 0001 */
	fmt.Printf("第三行 - c 的值为 %d\n", c )

	c = a << 2     /* 240 = 1111 0000 */
	fmt.Printf("第四行 - c 的值为 %d\n", c )

	c = a >> 2     /* 15 = 0000 1111 */
	fmt.Printf("第五行 - c 的值为 %d\n", c )
}

// 逻辑运算符
func LearnOpearatorLogical(){
	var a bool = true
	var b bool = false
	if ( a && b ) {
		fmt.Printf("第一行 - 条件为 true\n" )
	}
	if ( a || b ) {
		fmt.Printf("第二行 - 条件为 true\n" )
	}
	/* 修改 a 和 b 的值 */
	a = false
	b = true
	if ( a && b ) {
		fmt.Printf("第三行 - 条件为 true\n" )
	} else {
		fmt.Printf("第三行 - 条件为 false\n" )
	}
	if ( !(a && b) ) {
		fmt.Printf("第四行 - 条件为 true\n" )
	}
}

// 关系运算符
func LearnOpearatorRelation(){
	var a int = 21
	var b int = 10

	if( a == b ) {
		fmt.Printf("第一行 - a 等于 b\n" )
	} else {
		fmt.Printf("第一行 - a 不等于 b\n" )
	}
	if ( a < b ) {
		fmt.Printf("第二行 - a 小于 b\n" )
	} else {
		fmt.Printf("第二行 - a 不小于 b\n" )
	}

	if ( a > b ) {
		fmt.Printf("第三行 - a 大于 b\n" )
	} else {
		fmt.Printf("第三行 - a 不大于 b\n" )
	}
	/* Lets change value of a and b */
	a = 5
	b = 20
	if ( a <= b ) {
		fmt.Printf("第四行 - a 小于等于 b\n" )
	}
	if ( b >= a ) {
		fmt.Printf("第五行 - b 大于等于 a\n" )
	}
}




// 算术运算符
func LearnOperatorArithmetic(){
	var a int = 10
	var b int = 10
	var c int
	c = a+b
	fmt.Print("ab相加的值为：") // 20
	fmt.Print(c)
	fmt.Print("\n")
	c = a-b
	fmt.Print("ab相减的值为：") // 0
	fmt.Print(c)
	fmt.Print("\n")
	c = a * b
	fmt.Print("ab相乘的值为：") // 100
	fmt.Print(c)
	fmt.Print("\n")
	c = a/b
	fmt.Print("ab相除的值为：") // 1
	fmt.Print(c)
	fmt.Print("\n")
	c = a%b
	fmt.Print("ab求余的值为：") // 0
	fmt.Print(c)
	fmt.Print("\n")
	a++
	fmt.Print("a自增的值为：") // 11
	fmt.Print(a)
	fmt.Print("\n")
	a--
	fmt.Print("ab自减的值为：") // 10
	fmt.Print(a)
	fmt.Print("\n")
	a = 19 // 方便测试新赋值 a 的值 为19
	a--
	fmt.Print("ab自减的值为：") // 18
	fmt.Print(a)
}
