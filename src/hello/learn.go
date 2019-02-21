// 程序所属包 必须所在第一行
package hello
// 导入依赖包 可以导入多个包
// 注意不能导入 未使用的 package 包
import (
	"fmt"
)
// 定义常量 及 赋值 string 是类型
const NAME string = "Gao"
// 全局变量定义 及 赋值 全局变量关键字 var
var  globalVar string = "global variable"
// 一般类型声明
type gaoInt int
// 结构体的声明
type Learn struct {

}
// 声明接口
type ILearnToday interface{

}
// 函数定义
func LearnFunction(){
	fmt.Println("Today Learn Golang! ")
}
// 初始化 函数 相当于这个文件的入口
func init() {
	LearnFunction()
	fmt.Println("Learn Golang!")
}
