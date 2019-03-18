# Go interface


* 通过关键字type和interface，声明出接口类 `type TestInterface interface {}`
* 因为接口类型与其他数据类型不同，它是没法被实例化的。既不能通过调用new函数或make函数创建出一个接口类型的值
* 我们通过 interface 来定义对象的一组行为方法，如果某个对象实现了某个接口类型的所有方法，则此对象就是这个接口的实现类型
 	 实现约束:
		 1. 两个方法的签名需要完全一致
		 2. 两个方法的名称要一模一样


代码实现:

```
type Pet interface {
    SetName(name string)
    Name() string
    Category() string
}

type Dog struct {
    name string
}

func (dog *Dog) SetName(name string) {
    dog.name = name
}

func (dog Dog) Name() string {
    return dog.name
}

func (dog Dog) Category() string {
    return "dog"
}

func TestDog(t *testing.T) {
    dog := Dog{"little pig"}
    _, ok := interface{}(dog).(Pet)
    fmt.Printf("Dog implements interface Pet: %v\n", ok) // Dog implements interface Pet: false
    _, ok = interface{}(&dog).(Pet)
    fmt.Printf("*Dog implements interface Pet: %v\n", ok) // *Dog implements interface Pet: true
    var pet Pet = &dog
    fmt.Printf("This pet is a %s, the name is %q.\n",
        pet.Category(), pet.Name()) //  This pet is a dog, the name is "little pig".
	
    dog.SetName("monster")
    fmt.Printf("This pet is a %s, the name is %q.\n",
        pet.Category(), pet.Name()) // This pet is a dog, the name is "monster".
}
```

###  空接口
1. 空接口可以表示任何类型
2.  2.通过断言来将空接口转换为制定   //  v,ok := p.(int) // ok = true 时转换成功

###  go 接口最佳实践
####  倾向于使用小的接口定义,很多接口只包含一个方法
```
type Reader interface {
	Read(p []byte) (n int, err error)
}
```
#### 较大的接口定义,可以由多个小接口定义组合而成
```
type ReadWrite interface {
	Reader
	Write
}
```

####  只依赖与必要的功能最小接口**
```
func storeData(reader Reader) error {
  …
}
```




