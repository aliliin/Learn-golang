<h1>Go 切片slice</h1>

#### slice 切片类型基本概念
- slice 总是指向一个底层的 array，slice 的声明也可以像 array 一样，只是不需要定义长度
- slice 并不是真正意义上的动态数组，而是一个引用类型。
- slice 是引用类型，所以当改变其中元素的值时，其他的所有引用都会改变该值。
- slice 可以从一个数组或一个已经存在的slice中再次声明。slice 通过 array[i:j] 来获取，其中 i 是数组的开始位置，j 是结束位置，但不包含 array[j]，它的长度是 j - i。
- 声明 slice 时，方括号内没有任何字符，


#### slice 切片类型基本使用
- 基本声明语法
```
	var sliceType []type
	// 声明一个含有 int 类型元素的 slice
	var aSlice =  []int{1,2,3,4,5}
	fmt.Print(aSlice)     // 全部输出 [1,2,3,4,5]
	fmt.Print(aSlice[0])  // 输出 1
```
- make 指定 slice 完整的属性

```
// 第一个参数是类型，
// 第二个参数包含多少个元素，
// 第三个参数是 slice 的初始容量(如果不设置初始化容量，默认的内存长度为你的值的长度)
	s1 := make([]int,3)
	fmt.Print(len(s1),cap(s1)) // 3,3
	s2 := make([]int,3,10)
	fmt.Print(len(s2),cap(s2)) // 3,10
```
- 练习
```
	// 声明一个含有 8 int 类型元素的 数组
	var array = [8]int{1,3,4,6,7,8,9,10}
    // 声明三个含有 int 类型的 slice
	var bSlice,cSlice,dSlice []int
	dSlice = array[:2]  // 等价于 aSlice = array[0:3] aSlice 包含元素0和1不包含2 结果为:1,3
	bSlice = array[3:]  // 等价于 aSlice = array[3:5] aSlice 结果为：6，7,8,9,10
	cSlice = array[:]   // 等价于 aSlice = array[0:8] aSlice 包含元素所有的array 的元素
	// slice 上面的一些简便的操作解释
	/**
	 * slice 默认是从0开始的。arr[:n] 和 arr[0:n] 是一个意思
	 * slice 的第二个序列默认是数组的长度， arr[n:] 等价于 arr[n:len(arr)]
	 * 如果从一个数组里面直接获取 slice 可以这样子 arr[:] ,因为默认第一个序列是 0 第二个是数组的长度，即等价于 arr[0：len(arr)]
	 */

   // 从 slice 中获取 slice
	var eSlice []int
	eSlice = cSlice[1:3] // 包含了 cSlice[1],cSlice[2] 也就是 3，4
	eSlice = cSlice[0:3] // 包含了 cSlice[0],cSlice[1],cSlice[2] 也就是1,3,4
	// 注：读取 slice 时索引以被 slice 的切片为准，索引不可以超过被 slice 切片的容量cap() 的值

	// 声明一个含有10个元素元素类型为byte的数组
	var ar = [10]byte {'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}

	// 声明两个含有byte的slice
	var aa, bb []byte

	// a指向数组的第3个元素开始，并到第五个元素结束，
	aa = ar[2:5]
	// 现在 aa 含有的元素: ar[2]、ar[3]和ar[4]
	fmt.Print(string(aa)) // c d e
	// bb 是数组 ar 的另一个 slice
	fmt.Print(len(aa),cap(aa)) // 3  8 长度为3，容量为 8
	bb = aa[0:5] // 对 aa 的 bb 可以在 cap 范围内扩展，此时 bb 包含：c,d,e,f,g
	fmt.Print(string(bb)) // c,d,e,f,g
```


#### slice 中可以使用的内置的函数
- len 获取 slice的最大长度
- cap 获取 slice 的最大容量
- copy 函数 copy 从源 slice 的 src 中复制元素到目标 dst ，并且返回复制的元素的个数
- append 像 slice 里面追加一个或者多个元素，然后返回一个和 slice 一样类型的 slice
` 注意： append 函数会改变slice所引用的数组的内容，从而影响到引用同一数组的其他slice ，但当slice中没有剩余空间
	（即（ cap - len) = 0 ）时，此时将动态分配新的数组空间，返回的 slice 数组指针将指向这个空间。而原数组的内容将保持不变；
	其他引用此数组的 slice 则不受影响。`
