package timeutil

import (
	"fmt"
	"time"
)

func PrintTime() {
	fmt.Println(time.Now())
}

func init() {
	fmt.Println("timeutils 包下的 init() 函数被执行")
}
