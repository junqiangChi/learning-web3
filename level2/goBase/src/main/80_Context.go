package main

import (
	"fmt"
	"net/http"
	"time"
)

func hello1(w http.ResponseWriter, req *http.Request) {

	ctx := req.Context()
	fmt.Println("server: hello handler started")
	defer fmt.Println("server: hello handler ended")

	select {
	case <-time.After(11 * time.Second):
		fmt.Fprintf(w, `{"name": "hello"}`)
	case <-ctx.Done(): // 如果请求的上下文被取消（例如客户端关闭了连接），则执行这个分支
		err := ctx.Err()
		fmt.Println("server:", err)
		internalError := http.StatusInternalServerError
		http.Error(w, err.Error(), internalError)
	}
}

func main() {

	http.HandleFunc("/hello", hello1)
	http.ListenAndServe(":8090", nil)
}
