package learn

import "fmt"

// slice 类型学习
// 类似动态数组的类型，在Go中被称为 slice 类型
// slice 并不是真正意义上的动态数组，而是一个引用类型。
// slice 是引用类型，所以当改变其中元素的值时，其他的所有引用都会改变该值。
// slice 总是指向一个底层的 array，slice 的声明也可以像 array 一样，只是不需要定义长度
// slice 可以从一个数组或者一个已经存在的 slice 中再次声明
// slice 通过 array[key] 来取值的。
// 注意 声明 slice 时，方括号内没有任何字符

// var sliceType []int

<<<<<<< HEAD
func LearnSlice() {
	sliceType := []int{1, 2, 4, 5}
	fmt.Print(sliceType)    // 全部输出
	fmt.Print(sliceType[0]) // 输出1
	// 声明一个含有 10 个 byte类型的元素 数组
	var arr = [10]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}
	//  声明两个含有 byte 类型的 slice
	var a, b []byte
=======
func LearnSlice(){
	sliceType := []int{1,2,4,5}
	fmt.Print(sliceType) // 全部输出
	fmt.Print(sliceType[0])  // 输出1
	// 声明一个含有 10 个 byte类型的元素 数组
	var arr = [10]byte {'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}
	//  声明两个含有 byte 类型的 slice
	var  a,b []byte
>>>>>>> 120392ef192ab92d5a6fd6c6e1362322de3e2d7b
	a = arr[2:5]
	b = arr[3:5]
	fmt.Print(a)
	fmt.Print(b)
	/**
<<<<<<< HEAD
	slice 有一些简便的操作
	slice 默认是从0开始的。arr[:n] 和 arr[0:n] 是一个意思
	slice 的第二个序列默认是数组的长度， arr[n:] 等价于 arr[n:len(arr)]
	如果从一个数组里面直接获取slice 可以这样子 arr[:] ,因为默认第一个序列是 0 第二个是数组的长度，即等价于 arr[0：len(arr)]
	*/
	var array = [8]int{1, 3, 4, 6, 7, 8, 9, 10}
	var aSlice, bSlice, cSlice []int
	aSlice = array[:2] // 等价于 aSlice = array[0:3] aSlice 包含元素 1，3
	bSlice = array[3:] // 等价于 aSlice = array[3:5] aSlice 包含元素 6，7,8,9,10
	cSlice = array[:]  // 等价于 aSlice = array[0:8] aSlice 包含元素 所有的array 的元素
	fmt.Print(aSlice)
	fmt.Print(bSlice)
	fmt.Print(cSlice)

	// 从 slice中获取 slice
	var dSlice, eSlice []int
	dSlice = cSlice[1:3] // 3,4 包含了 cSlice[1],cSlice[2] 也就是 3，4
	fmt.Print(dSlice)
	eSlice = cSlice[0:3] //  包含了cSlice[0] cSlice[1],cSlice[2] 也就是1， 3，4
	fmt.Print(eSlice)    // 6,7
=======
		slice 有一些简便的操作
		slice 默认是从0开始的。arr[:n] 和 arr[0:n] 是一个意思
		slice 的第二个序列默认是数组的长度， arr[n:] 等价于 arr[n:len(arr)]
		如果从一个数组里面直接获取slice 可以这样子 arr[:] ,因为默认第一个序列是 0 第二个是数组的长度，即等价于 arr[0：len(arr)]
	 */
	 var array = [8]int{1,3,4,6,7,8,9,10}
	 var aSlice,bSlice,cSlice []int
	 aSlice = array[:2] // 等价于 aSlice = array[0:3] aSlice 包含元素 1，3
	 bSlice = array[3:] // 等价于 aSlice = array[3:5] aSlice 包含元素 6，7,8,9,10
	 cSlice = array[:] // 等价于 aSlice = array[0:8] aSlice 包含元素 所有的array 的元素
	 fmt.Print(aSlice)
	 fmt.Print(bSlice)
	 fmt.Print(cSlice)


	// 从 slice中获取 slice
	var dSlice,eSlice []int
	dSlice = cSlice[1:3] // 3,4 包含了 cSlice[1],cSlice[2] 也就是 3，4
	fmt.Print(dSlice)
	eSlice = cSlice[0:3]  //  包含了cSlice[0] cSlice[1],cSlice[2] 也就是1， 3，4
	fmt.Print(eSlice) // 6,7
>>>>>>> 120392ef192ab92d5a6fd6c6e1362322de3e2d7b

	/**
		slice 中有几个内置的函数
		len 获取slice的最大长度
	    cap 获取slice 的最大容量
		copy 函数 copy 从源slice 的src中复制元素到目标 dst ，并且返回复制的元素的个数
		append 像slice 里面追加一个或者多个元素，然后返回一个和slice 一样类型的slice
		注意： append 函数会改变 slice 所引用的数组的内容，从而影响到引用同一数组的其他slice ，但当slice中没有剩余空间
	（即（ cap - len) = 0 ）时，此时将动态分配新的数组空间，返回的 slice 数组指针将指向这个空间。而原数组的内容将保持不变；
	其他引用此数组的slice 则不受影响。

<<<<<<< HEAD
	*/
	//  能指定 slice 完整的属性
	// 第一个参数是类型，第二个参数包含多少个元素，第三个参数是 slice 的初始容量(如果不设置初始化容量，默认的内存长度为你的值的长度)
	s1 := make([]int, 3)
	fmt.Print(len(s1), cap(s1)) // 3,3
	s2 := make([]int, 3, 10)
	fmt.Print(len(s2), cap(s2)) // 3,10

	// 声明一个含有10个元素元素类型为byte的数组
	var ar = [10]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}
=======
	 */
	 //  能指定 slice 完整的属性
	 // 第一个参数是类型，第二个参数包含多少个元素，第三个参数是 slice 的初始容量(如果不设置初始化容量，默认的内存长度为你的值的长度)
	s1 := make([]int,3)
	fmt.Print(len(s1),cap(s1)) // 3,3
	s2 := make([]int,3,10)
	fmt.Print(len(s2),cap(s2)) // 3,10

	// 声明一个含有10个元素元素类型为byte的数组
	var ar = [10]byte {'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}
>>>>>>> 120392ef192ab92d5a6fd6c6e1362322de3e2d7b

	// 声明两个含有byte的slice
	var aa, bb []byte

	// a指向数组的第3个元素开始，并到第五个元素结束，
	aa = ar[2:5]
	//现在a含有的元素: ar[2]、ar[3]和ar[4]
	fmt.Print(string(aa)) // c d e
	// b是数组ar的另一个slice
<<<<<<< HEAD
	fmt.Print(len(aa), cap(aa)) // 3  8
	bb = aa[0:5]                // 对 bb的 aa 可以在cap范围内扩展，此时 bb 包含：c,d,e,f,g
	fmt.Print(string(bb))       // f g
	// b的元素是：ar[3]和ar[4]
}
=======
	fmt.Print(len(aa),cap(aa)) // 3  8
	bb = aa[0:5] // 对 bb的 aa 可以在cap范围内扩展，此时 bb 包含：c,d,e,f,g
	fmt.Print(string(bb)) // f g
	// b的元素是：ar[3]和ar[4]
}

>>>>>>> 120392ef192ab92d5a6fd6c6e1362322de3e2d7b
