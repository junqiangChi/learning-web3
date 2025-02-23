package main

import (
	"fmt"
	"time"
)

func main() {

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	// select 语句是一种用于处理多个通道操作的强大工具，它主要用于在多个通道操作上进行非阻塞的选择。
	// select 语句的行为类似于 switch 语句，但它关注的是通道的通信操作（发送和接收），而不是值的比较
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
}
