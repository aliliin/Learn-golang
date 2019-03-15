# Go 面向对象


### 实例创建及初始化
* 基本使用

```golang
type Employee struct {
	Id   int
	Name string
	Age  int
}

e := Employee{"0", "Bob", 20}
e1 := Employee{name: "Ali", Age: 30}
// 注意这里返回的引用/指针, 相当于 e:= &Employee{}
e2 := new(Employee)
// 与其他的编程语言不同, 可以不用 -> 的方式来访问指定的属性
e2.Id = "2"
e2.Age = 23
e2.Name = "Ali"
```
* 非指针传递会生成一个新的结构对象，其中每个成员会复制。指针传递，只是传递指针，且该指针指向原有结构。
```
// 这一种,定义方式在实例对应方法变调用的时候 ,实例的成员会进行值复制
func (e Employee) String() string {
  fmt.Printf("输出自己的e的地址 %x\n", unsafe.Pointer(&e.Name)) // c00000a0e8
  return fmt.Sprintf("ID :%d-Name:%s-Age:%d", e.Id, e.Name, e.Age)
}

// 在通常情况下避免内存拷贝我们使用 这种方式
func (e *Employee) String() string {
    fmt.Printf("输出自己的e的地址 %x\n", unsafe.Pointer(&e.Name)) // c00000a0e8
    return fmt.Sprintf("ID :%d/Name:%s/Age:%d", e.Id, e.Name, e.Age)
}

// 测试 String 方法
func TestInit(t *testing.T) {
    e := Employee{1, "Bob", 20}
    // 第一种
    // t.Log(e.String()) // ID :1-Name:Bob-Age:20
    fmt.Printf("输出传进来的e的地址 %x\n", unsafe.Pointer(&e.Name)) // c00000a108

    // 第二种
    t.Log(e.String()) // ID :1/Name:Bob/Age:20
    fmt.Printf("输出传进来的e的地址 %x\n", unsafe.Pointer(&e.Name))  // c00000a0e8

}
```