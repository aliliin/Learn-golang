

# Go map类型
<hr>
## 基本定义

* 类似其它语言中的哈希表或者字典，以 key => value 形式存储数据  简单的理解为 PHP 中关联数组
* 格式为 map[keyType]valueType
* key 必须是支持 == 或者 != 比较运算的类型，不可以是函数，map 或者 slice 
* map 使用 make() 创建，支持 := 这种简写的方式
* map 是无序的，每次打印出来的 map 都会不一样，它不能通过 index 获取，而必须通过 key 获取
* map 的长度是不固定的，也就是和 slice 一样，也是一种引用类型
* 内置的len函数同样适用于 map，返回 map 拥有的 key 的数量
* map 的值可以很方便的修改，通过 mapType[“one”] = 11 可以很容易的把 key 为 one 的值改为11
* 和 slice 一样,如果两个map 同时指向同一个底层数据,那么任何一个改变.另一个也会改变


## 代码练习

```golang

var numbers map[string]int  // 声明一个 key 是字符串，值为 int 的map 

numbers = make(map[string]int] // 另一种声明方式 map 使用 make() 创建


numbers := make(map[string]int) // 支持 := 这种简写的方式
numbers["one"] = 1          // 赋值 
numbers["two"] = 2          // 赋值
fmt.Println(numbers["one"]) //  而必须通过 key 获取对应的值. 输出 1

rating := map[string]float32{"C": 5, "Go": 4.5, "Python": 4.5, "PHP": 2} // 初始化一个有值的map
csharpRating, ok := rating["C++"] // 新赋值给别的变量
if ok {
	fmt.Println("输出ok的值", csharpRating)
} else {
	fmt.Println("没找到对应的OK的值")
}

delete(rating, "C")      // 删除key为C的元素
fmt.Println(rating["C"]) // 输出 0

// 内置的len函数同样适用于 map，返回 map 拥有的 key 的数量
fmt.Println(len(rating)) // 因为删除了一个key 输出 3
// 和 slice 一样,如果两个map 同时指向同一个底层数据,那么任何一个改变.另一个也会改变
m := make(map[string]string)
m["Hello"] = "你好"
m1 := m
m1["Hello"] = "不好"  // 现在m["Hello"]的值已经是 不好 了


mapa := map[string]int{"one":1,"two":2}
mapb := map[string]int{"yi": 1, "er": 2}
if mapa["one"] == mapb["yi"] { //  key 必须是支持 == 比较运算
	fmt.Println("可以比较")
	}
if mapa["one"] != mapb["er"] { //  key 必须是支持 != 比较运算
	fmt.Println("可以比较")
}

```



