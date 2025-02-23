package main

import "fmt"

func main() {
	// 通道（Channels）是一种特殊的类型，用于在不同的 goroutine 之间进行安全的数据通信和同步。通道提供了一种类型安全、并发安全的方式来传递数据
	messages := make(chan string)
	// 使用 <- 操作符将数据发送到通道，语法为 通道名 <- 数据。
	go func() { messages <- "ping" }()

	// 从通道中接收数据
	msg := <-messages
	fmt.Println(msg)
}
