package main

import "fmt"

func main() {
	/**
	 创建：make(map[string]int)
	获取：m[key]
	key 不存在时，获得 value 类型的初始值
	用value，ok：=m[key]来判断是否存在key
	用 delete 删除key
	使用 range 遍历key，或者遍历key，value 队
	不保证遍历的顺序，如需要顺序，需要手动对key排序
	map 使用哈希表，必须可以比较相等
	除了 slice，map function 的内建类型都可以作为key
	struct 类型不饱含上述字段，也可以作为key

	*/
	m := map[string]string{
		"name":    "ccmouse",
		"course":  "golang",
		"site":    "imooc",
		"quality": "notbad",
	}
	m2 := make(map[string]int) // m2 === empty map
	var m3 map[string]int      // m3 === nil
	fmt.Println(m, m2, m3)
	for k, v := range m {
		fmt.Println(k, v)
	}
	fmt.Println("Getting values")
	courseName, ok := m["course"]
	fmt.Println(courseName, ok)
	courseNames, ok := m["coure"]
	fmt.Println(courseNames, ok)
	if courseName, ok := m["coures"]; ok {
		fmt.Println(courseName)
	} else {
		fmt.Println("Key does not exist")
	}
	fmt.Println("Deleting values ")
	name, ok := m["name"]
	fmt.Println(name, ok)

	delete(m, "name")
	name, ok = m["name"]
	fmt.Println(name, ok)
}
