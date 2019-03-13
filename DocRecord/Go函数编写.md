
#Learning
# **Go 函数**
## 函数声明
* 关键字 `func` 用来声明一个函数
* 函数可以有一个或者多个参数，每个参数后面带有不同的类型，通过,分隔
* 函数可以返回多个值
* 如果只有一个返回值且不声明返回值变量，那么你可以省略 包括返回值 的括号
* 如果没有返回值，那么就直接省略最后的返回信息
* 如果有返回值， 那么必须在函数的外层添加return语句
* 函数可以作为变量的值
* 函数可以作为参数的返回值

## 函数基本定义
```golang
// 函数定义
func funcName(inputOne type, inputTwo type) (outputOne type, outputTwo type) {**
			// 逻辑代码...
			// 可以返回多个值
			return valueOne, valueTwo
}
// 返回指定类型
func Testmax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 返回两个值
func SumAndProduct(A, B int) (int, int) {
	return A+B, A*B
}

// 有意思的函数编写
// 计算函数运行时长
func timeSpent(inner func(op int) int) func(op int) int {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("time spent:", time.Since(start).Seconds())
		return ret
	}
}
```