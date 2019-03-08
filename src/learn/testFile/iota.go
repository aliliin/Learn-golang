package learn

import "fmt"

// iota 只能在常量定义中使用 新增一行iota常量声明 iota 就会自动加一
const iotaa = iota
const iotab = iota

const ( // 每新增一行常量声明 iota 就会自动加一
	iotac = iota
	iotad = iota
)

// 跳值使用法
const (
	iotae = iota // 0
	_            // 通过下划线来达到 跳值使用法  会把1 赋值给下划线1
	_            // 2
	iotaf = iota // 3
)

// 插队使用
const (
	iotag = iota // 0
	iotah = 3.14 // 3.14
	iotaj = iota // 2
)

// 单行使用法
const (
	iotay, iotau = iota, iota + 3 // 0 , 3
	iotao, iotas                  // 1,4 iotao 沿用第一个也就是 iotay的，而 iotas 沿用的是iotau的
	ff           = iota           // 2

)

// 有趣的现象
const (
	iotaq  = iota * 2 // 0
	iotaw             // 如果不赋值默认依次继承 ，就成了  2
	iotar             //  4
	iotaps            // 6
	iotap             // 8
)

func init() {

}

func LearnIota() {
	fmt.Print(iotaa)
	fmt.Print("\n")
	fmt.Print(iotab)
	fmt.Print("\n")
	fmt.Print(iotac)
	fmt.Print("\n")
	fmt.Print(iotad)
	fmt.Print("\n")
	fmt.Print(iotae)
	fmt.Print("\n")
	fmt.Print(iotaf)
	fmt.Print("\n")
	fmt.Print(iotag)
	fmt.Print("\n")
	fmt.Print(iotah)
	fmt.Print("\n")
	fmt.Print(iotaj)
	fmt.Print("\n")
	fmt.Print(iotaq)
	fmt.Print("\n")
	fmt.Print(iotaw)
	fmt.Print("\n")
	fmt.Print(iotar)
	fmt.Print("\n")
	fmt.Print(iotaps)
	fmt.Print("\n")
	fmt.Print(iotap)
	fmt.Print("\n")
	fmt.Print(iotay)
	fmt.Print("\n")
	fmt.Print(iotau)
	fmt.Print("\n")
	fmt.Print(iotao)
	fmt.Print("\n")
	fmt.Print(iotas)
	fmt.Print(ff)
}
