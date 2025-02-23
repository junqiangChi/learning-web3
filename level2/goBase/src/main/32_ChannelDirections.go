package main

import "fmt"

// 发送方向（Send-Only） chan<-
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// 接收方向（Receive-Only） <-chan
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
