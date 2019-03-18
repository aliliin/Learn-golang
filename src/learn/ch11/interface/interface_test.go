package interface_test

import (
	"fmt"
	"testing"
	"time"
)

// 接口定义
type Programmer interface {
	WriteHelloWorld() string
}

// 类定义 接口实现
type GoProgrammer struct {
}

// 	签名要和 interface 接口完全一致
func (g *GoProgrammer) WriteHelloWorld() string {
	return "fmt.Println(\"Hello World\")"
}

func TestInterface(t *testing.T) {
	// 定义
	var p Programmer

	p = new(GoProgrammer)
	// 调用
	t.Log(p.WriteHelloWorld())
}

type Learn interface {
	GolangLearn() string
}

type LearnGolang struct {
}

func (l *LearnGolang) GolangLearn() string {
	return "Learn Golang"
}

func TestLearn(t *testing.T) {
	var study Learn
	study = new(LearnGolang)
	t.Log(study.GolangLearn())
}

type class func(op int) int

// 计算函数运行时长
func timeSpent(inner class) class {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("time spent:", time.Since(start).Seconds())
		return ret
	}
}

func slowFunc(op int) int {
	time.Sleep(time.Second * 1)
	return op
}
func TestFn(t *testing.T) {
	tsSf := timeSpent(slowFunc)
	t.Log(tsSf(10))

}

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
	fmt.Printf("Dog implements interface Pet: %v\n", ok)

	_, ok = interface{}(&dog).(Pet)
	fmt.Printf("*Dog implements interface Pet: %v\n", ok)

	var pet Pet = &dog
	fmt.Printf("This pet is a %s, the name is %q.\n",
		pet.Category(), pet.Name())

	dog.SetName("monster")

	fmt.Printf("This pet is a %s, the name is %q.\n",
		pet.Category(), pet.Name())

	var dog1 *Dog
	fmt.Println("The first dog is nil. [wrap1]")
	dog2 := dog1
	fmt.Println("The second dog is nil. [wrap1]")
	pet = dog2
	if pet == nil {
		fmt.Println("The pet is nil. [wrap1]")
	} else {
		fmt.Println("The pet is not nil. [wrap1]")
	}

}
