# Go 错误处理 Panic 和 Recover
#Learning
1. Go中没有错误机制
2. error 类型实现了 error 接口
```
type error interface {
	Error() string
}
```
3. 可以通过 errors.New 来实现快速创建错误信息
```
errors.New("n must be in the range [0,10]
```

#### Panic 的基本定义和使用
1. panic  用于不可以恢复的错误
2. Panic 退出前会执行 defer 指定的内容、
```
func TestPanic(t *testing.T) {
	defer func() {
		fmt.Println("最后结果依旧执行!") // 这一部分的代码依旧执行
	}()
	fmt.Println("执行开始")
	panic(errors.New("错误信息!"))
}
// 输出大致结果如下
开始
最后结果依旧执行!
--- FAIL: TestPanic (0.00s)
panic: 错误信息! [recovered]
panic: 错误信息!
```
#### os.Exit 的基本定义和使用
1. os.Exit 退出时不调用defer 指定的函数
2. os.Exit 退出时不输出当前调用栈信息
```
func TestExit(t *testing.T) {
	defer func() {
		fmt.Println("这段代码不会执行!") 
	}()
	fmt.Println("开始执行")
	os.Exit(-1) // 直接退出 exit status 255
}
// 输出结果如下
开始执行
exit status 255

```

#### Recover 的使用
```
func TestRecover(t *testing.T) {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recoverd 开始\n", err)
		}
	}()
	fmt.Println("开始执行")
	panic(errors.New("错误信息!"))
}
// 输出结果如下:
开始执行
recoverd 开始
 错误信息!
```
#### 异常处理总结
	当一个函数在执行过程中出现了异常或遇到 panic()，正常语句就会立即终止，然后执行 defer 语句，再报告异常信息，最后退出。如果在 defer 中使用了 recover() 函数,则会捕获错误信息，使该错误信息终止报告。


