package learn
// 数据类型
import (
	"fmt"
	"unsafe"
)
// 初始化 数据类型学习方法
func init() {
	// 	无符号整型
	var uint8Data uint8 = 1 // 1
	var int32Data int32 = 1
	var intData   int = 1 // 根据系统的位数改变 64位是8
	var uintData  uint = 1
	var float32Data float32 = 1 // float 类型后面数据不能不添加
	var float64Data float64 = 1 // float 类型后面数据不能不添加
	var boolTrueType bool = true // 不能是其他类型，只能是true和false 不能是1
	var  byteData byte = 1 // 无符号整型 uint8 的大小一样
	var  runeData rune  = 1 // 无符号整型 uint32  的大小一样 4
	// 数据类型的默认值
	var i int32 //  0
	var f float32 // 0
	var b bool // false
	var c complex64 // (0+0i)
	var s string // 空值
	type 整型 int32 // 类型别名 是通过type 定义的。类型别名不能参与计算，相同别名的类型是可以的参与计算的
	var  zhengxing 整型 = 1
	fmt.Print("类型别名之后默认值为")
	fmt.Print(zhengxing) // 数据类型为 0
	fmt.Print("\n")
	fmt.Print("类型别名之后的字节大小为")
	fmt.Print(unsafe.Sizeof(zhengxing)) // 4
    fmt.Print("\n")
	fmt.Print(i) // 默认值 0
	fmt.Print(f) // 默认值 0
	fmt.Print(b) // 默认值 false
	fmt.Print(c) // 默认值 (0+0i)
	fmt.Print(s) // 默认值 空值

	fmt.Print("\n") //换行
	fmt.Print(unsafe.Sizeof(uint8Data)) // 1
	fmt.Print("\n")
	fmt.Print(unsafe.Sizeof(int32Data)) // 4
	fmt.Print("\n")
	fmt.Print(unsafe.Sizeof(intData)) // 8
	fmt.Print("\n")
	fmt.Print(unsafe.Sizeof(uintData)) // 8
	fmt.Print("\n")
	fmt.Print(unsafe.Sizeof(float32Data)) // 4
	fmt.Print("\n")
	fmt.Print(unsafe.Sizeof(float64Data)) // 8
	fmt.Print("\n")
	fmt.Print(boolTrueType) // true
	fmt.Print("\n")
	fmt.Print(unsafe.Sizeof(byteData)) // 1
	fmt.Print("\n")
	fmt.Print(unsafe.Sizeof(runeData)) // 4
	fmt.Print("\n")

}

// 调用本包的方法 实现初始化init 方法
func LearnTypes(){

}