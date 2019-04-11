package reflect_test

// 反射编程练习
import (
	"fmt"
	"reflect"
	"testing"
)

func TestTypeAndValue(t *testing.T) {
	var f int64 = 10
	t.Log(reflect.TypeOf(f), reflect.ValueOf(f))
	t.Log(reflect.ValueOf(f).Type())
}

func TestType(t *testing.T) {
	var f int64 = 10
	CheckType(f)
}

// 查找参数类型
func CheckType(v interface{}) {
	t := reflect.TypeOf(v)
	switch t.Kind() {
	case reflect.Float32, reflect.Float64:
		fmt.Println("Float")
	case reflect.Int, reflect.Int64, reflect.Int32:
		fmt.Println("Int")
	case reflect.String:
		fmt.Println("String")
	default:
		fmt.Println("Unknown", t)
	}
}

func TestDeepEqual(t *testing.T) {
	a := map[int]string{1: "one", 2: "two", 3: "three"}
	b := map[int]string{1: "one", 2: "two", 3: "three"}
	//	t.Log(a == b)
	t.Log("a==b?", reflect.DeepEqual(a, b))

	s1 := []int{1, 2, 3}
	s2 := []int{1, 2, 3}
	s3 := []int{2, 3, 1}

	t.Log("s1 == s2?", reflect.DeepEqual(s1, s2))
	t.Log("s1 == s3?", reflect.DeepEqual(s1, s3))

	c1 := Customer{"1", "Mike", 40}
	c2 := Customer{"1", "Mike", 40}
	fmt.Println(c1 == c2)
	fmt.Println(reflect.DeepEqual(c1, c2))
}

// 利用反射编写灵活的代码
/**
1.按照名字访问结构的成员
 reflect.ValueOf(*e).FieldByName("Name")

2.按照名字访问结构的方法
 reflect.ValueOf(e).MethodByName("updateAge").Call([]reflect.Value{reflect.ValueOf(1)})

*/

type Employee struct {
	EmployeeID string
	Name       string `format:"normal"`
	Age        int
}

func (e *Employee) UpdateAge(newAge int) {
	e.Age = newAge
}

type Customer struct {
	CookieID string
	Name     string
	Age      int
}

// func TestInvokeByName(t *testing.T) {
// 	e := &Employee{"1", "Mike", 30}
// 	// 按照名字获取成员
// 	// t.Logf("Name: value(%[1]v), Type(%[1]T) ", reflect.ValueOf(*e).FieldByName("Name"))
// 	t.Logf("Name: value(%[1]v), Type(%[1]T) ", reflect.ValueOf(*e).FieldByName("Name"))
// 	if nameField, ok := reflect.TypeOf(*e).FieldByName("Name"); !ok {
// 		t.Error("Failed to get 'Name' field.")
// 	} else {
// 		t.Log("Tag:format", nameField.Tag.Get("format"))
// 	}
// 	reflect.ValueOf(e).MethodByName("updateAge").Call([]reflect.Value{reflect.ValueOf(1)})
// 	t.Log("Updated Age :", e)
// }

func TestInvokeByName(t *testing.T) {
	e := &Employee{"1", "Mike", 30}
	//按名字获取成员
	t.Logf("Name: value(%[1]v), Type(%[1]T) ", reflect.ValueOf(*e).FieldByName("Name"))
	if nameField, ok := reflect.TypeOf(*e).FieldByName("Name"); !ok {
		t.Error("Failed to get 'Name' field.")
	} else {
		t.Log("Tag:format", nameField.Tag.Get("format"))
	}
	reflect.ValueOf(e).MethodByName("UpdateAge").
		Call([]reflect.Value{reflect.ValueOf(1)})
	t.Log("Updated Age:", e)
}
