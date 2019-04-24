package main

import (
	"fmt"
	"net/http"
	"time"
)

// http 服务
func main() {
	// 定义不同的url下的不同的处理逻辑
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World!")
	})

	// 定义不同的url下的不同的处理逻辑
	http.HandleFunc("/time/", func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()
		timeStr := fmt.Sprintf("{\"time\":\"%s\"}", t)
		w.Write([]byte(timeStr))
	})

	// 启动一个server 设置端口
	http.ListenAndServe(":8080", nil)
}
