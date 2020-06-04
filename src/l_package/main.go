package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	_ "l_package/utils"
	"l_package/utils/timeutil"
)

//import "l_package/utils" // 绝对路径
/**
init 函数和main 函数
1.这两个函数都是go语言中的保留函数，init 用于初始化信息，main用于作为程序的入口
2。 这个两个 函数定义的时候不能有参数和返回值，只能由go程序自动调用个，不能被饮用
3。init 函数可以定义在任意包中，可以又多个，main 函数只能在main包下，并且只能有一个
4 执行顺序
	先i执行 init 函数，然后执行 main 函数
	对于同一个 go 文件中，调用顺序是从上到下的，之后顺序调用各个文件中的 init 函数
	对于同一个包下，讲文件名安装字符串进行排序，之后顺序调用各个文件中的 init 的函数
 	不同包情况下，
		存在依赖 最后被依赖的最先被初始化，
		不存在依赖，按照 main 包中 的 引入顺序执行 init
存在依赖的包之间，不能循环倒入
一个多可以被其他多个包 import 但是只能被初始化一次
如果只是为了 执行包中的 init 函数，不执行其他方法，可以使用
import 下划线"包路径和包名" import _"package"
*/

func main() {
	timeutil.PrintTime()
	//utils.Count()
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/laravel6?charset=utf8")
	if err != nil {
		fmt.Println("错误信息", err)
	}
	fmt.Println("链接成功", db)
}
