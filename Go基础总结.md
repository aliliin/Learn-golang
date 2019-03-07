## Go 基础总结

## 流程控制和函数
Go 中流程控制有三大类：条件判断、循环控制、无条件跳转
* if
* goto
* for
* switch

### if
* 不需要`()`
* 允许声明变量「作用域只在条件逻辑内」
```go
if a > b {
    //
} else {
    //
}
// 声明变量
if x := computedValue(); x > 10 {
    fmt.Println("x is greater than 10")
} else {
    fmt.Println("x is less than 10")
}
// 多个条件
if a > 4 {
    //
} else if a < 5 {
    //
} else {
    //
}

```

### goto
跳转到必须在当前函数定义的标签
```go
func myFunc() {
    i := 0
Hear:   // 这行的第一个词，以冒号结束作为标签
    fmt.Println(i)
    i++
    goto Hear   // 跳转到 Hear 去「大小写敏感」
}
```
### for
go 中的 for 的三大用处
* 循环读取数据`map` `slice` = php 的 foreach
* while 来控制逻辑 = php 的 while
* 迭代操作 = php 的 for

```go
// foreach
for k, v := range map {
    fmt.Println("map's key is ", k)
    fmt.Println("map's value is ", v)
}
//由于 Go 支持 “多值返回”, 而对于“声明而未被调用”的变量, 编译器会报错, 在这种情况下, 可以使用 _ 来丢弃不需要的返回值
for _, v := range map {
    //
}
// while
sum := 1
for sum < 100 {
    sum += sum
}
// for
sum := 0
for i := 0; i < 10; i++ {
    sum += i
}
```

### switch
与 Php 比较
* case 的数据类型必须一致
* 表达式不必是常量或者整数，可以没有表达式就找 true
* 自带 break
* 支持多值聚合在一个 case 里面
* `fallthrough` 强制执行下一个 case
* `case` 里面不能有重复条件 

```go
i := 10
switch i {
case 10:
	fmt.Println("i match 10")
	fallthrough // 强制执行下一个 case
case i*2 - 10:
	fmt.Println("表达式")
	fallthrough
case 1, 2, 3: // 不能有重复条件 如果有 10 就会报错
	fmt.Println("多值聚合 i is 1 2 or 10")
	fallthrough
case i == 10:
	fmt.Println("bool 类型会报错")
	fallthrough
case 4:
	fmt.Println("默认就有 break 的")
	default:
	fmt.Println("default")
}
```
## 函数
与 Php 不一样的地方
* 关键字 func
* 参数 & 返回值需要指定类型 `name type`
* 多返回值
* 变参

```go
// 多返回值1 参数相同类型可以简写
func sumAndProduct(a, b int) (int, int) {
    return a + b, a * b
}
// 多返回值2
func sumAndProduct(a int, b int) (sum int, product int) {
    sum = a + b
    product = a * b
    return
}
// 变参
func myfunc(arg ...int) {
    for _, n := range arg { // arg 是一个int 的 slice
        
    }
}
```

### 传值与传指针
PHP 一样，传值是传值的 copy, & 传地址。\
不同的是：\
channel，slice，map这三种类型的实现机制类似指针，所以可以直接传递，而不用取地址后传递指针
### defer
Php 中没有

作用：defer 后指定的函数会在函数退出前调用。多个 defer 调用顺序后进后出「栈」
```go
func ReadWrite() bool {
    file.Open("file")
    defer file.close() // 函数退出前这里会被调用执行
    if FailureX {
        return false
    }
    return true
}
```
### 函数作为值类型
函数当做值来传递
* 声明函数类型 = 定义数据类型
* 实现函数类型 = 实现数据类型
* 通过类型跟函数名 = 声明一个变量

```go
type testInt func(int) bool // 声明类型
// 类型的实现
func isTwo(integer int) bool {
    if (integer % 2) {
        return true
    }
    return false
}
// 类型的实现
func isFive(integer int) bool {
    if (integer % 5) {
        return true
    }
    return false
}
// 声明参数函数类型
func filter(slice []int, f testInt) []int {
    var result = []int
    for _, value := range slice {
        if f(value) {
            result = append(result, value)
        }
    }
    return result
}

func main() {
    slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    fmt.Println("slice = ", slice)
    two := filter(slice, isTwo)
    five := filter(slice, isFive)
    fmt.Println("倍 2 整除的:", two)
    fmt.Println("被 5 整除的:", five)
}
```