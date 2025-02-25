package main

import "fmt"

func main() {

	// 2个缓冲
	messages := make(chan string, 2)

	// 可以放两次
	messages <- "buffered"
	messages <- "channel"

	// 取两次
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
