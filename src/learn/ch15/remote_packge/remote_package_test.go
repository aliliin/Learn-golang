package remote

import (
	"fmt"
	"testing"
)

func TestConcurrentMap(t *testing.T) {
	// m := cm.CreateConcurrentMap(99)
	// m.Set(cm.StrKey("key"), 10)
	// t.Log(m.Get(cm.StrKey("key")))
	// defer_call()
	pase_student()
}

func defer_call() {

	defer func() { fmt.Println("打印前") }()
	defer func() { fmt.Println("打印中") }()
	defer func() { fmt.Println("打印后") }()
	panic("触发异常")
}

type student struct {
	Name string
	Age  int
}

func pase_student() {
	m := make(map[string]*student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	for _, stu := range stus {
		m[stu.Name] = &stu
		fmt.Println(stu)
	}
}
