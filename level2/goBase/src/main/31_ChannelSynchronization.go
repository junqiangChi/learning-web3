package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}

func main() {

	done := make(chan bool, 1)
	go worker(done)
	// <-done 会阻塞 主协程 直到done收到值
	// 如果删除这个 就不会阻塞主协程，worker函数的协程会随着主协程结束而强制关闭，这样worker函数协程可能未完成内容
	<-done
}
