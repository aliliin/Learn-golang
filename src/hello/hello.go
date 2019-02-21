package hello

import (
    "fmt"
    "mymath"
)
 
func init(){
    fmt.Printf("Hello World!\n")
}
func Hello(){
    mymath.Sqrt("测试导入重复的包")
}