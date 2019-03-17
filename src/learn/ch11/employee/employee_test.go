package employee_test

import (
	"fmt"
	"testing"
	"unsafe"
)

type Employee struct {
	Id   int
	Name string
	Age  int
}

func TestInitEmployee(t *testing.T) {
	e := Employee{1, "Bob", 20}
	e1 := Employee{Name: "Ali", Age: 30}
	// 注意这里返回的引用/指针, 相当于 e:= &Employee{}
	e2 := new(Employee)
	// 与其他的编程语言不同, 可以不用 -> 的方式来访问指定的属性
	e2.Id = 2
	e2.Age = 23
	e2.Name = "Ali"
	t.Log(e)               // {1 Bob 20}
	t.Log(e1)              //  {0 Ali 30}
	t.Log(e2)              // &{2 Ali 23}
	t.Log(e.Name)          // Bob
	t.Log(e1.Id)           // 0
	t.Logf("e is %T", e)   // e is employee_test.Employee 前面如果加上 &  那么返回的也是 指针类型
	t.Logf("e2 is %T", e2) // e2 is *employee_test.Employee 这个是指针类型
}

// 这一种,定义方式在实力对应方法变调用的时候 ,实例的成员会进行值复制
// func (e Employee) String() string {
/**
输出传进来的e的地址 c00000a0e8
输出传进来的e的地址 c00000a108
*/
// 	fmt.Printf("输出传进来的e的地址 %x\n", unsafe.Pointer(&e.Name))
// 	return fmt.Sprintf("ID :%d-Name:%s-Age:%d", e.Id, e.Name, e.Age)
// }

// 在通常情况下避免内存拷贝我们使用 这种方式
func (e *Employee) String() string {
	/**
	输出传进来的e的地址 c00000a0e8
	输出传进来的e的地址 c00000a0e8
	*/
	fmt.Printf("输出自己的e的地址 %x\n", unsafe.Pointer(&e.Name))
	return fmt.Sprintf("ID :%d/Name:%s/Age:%d", e.Id, e.Name, e.Age)
}

func TestInit(t *testing.T) {
	// e := Employee("10", "dd", 29)
	e := Employee{1, "Bob", 20}
	// 第一种
	// t.Log(e.String()) // ID :1-Name:Bob-Age:20
	fmt.Printf("输出传进来的e的地址 %x\n", unsafe.Pointer(&e.Name))
	// 第二种
	t.Log(e.String()) // ID :1/Name:Bob/Age:20

}

func TestStaff(t *testing.T) {
	type Human struct {
		name string
		age  int
	}
	type Staff struct {
		Human // 隐式的引入 Human 的字段
		wage  int
		age   float32 // 覆盖 Human 的 age
	}
	tom := Staff{Human{"Tom", 18}, 1000, 18.5}
	jerr := Staff{Human: Human{age: 32, name: "Jerr"}, wage: 2000, age: 33.5}
	fmt.Printf("%s is %d old, %.2f years, wage is %d \n", jerr.name, jerr.Human.age, jerr.age, jerr.wage)
	fmt.Printf("%s is %d old, %.1f years, wage is %d \n", tom.name, tom.Human.age, tom.age, tom.wage)
}
