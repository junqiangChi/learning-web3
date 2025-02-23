package main

import (
	"fmt"
	"time"
)

func main() {

	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	// 固定速率的定时器
	// 该通道会每隔 200 毫秒接收一个当前时间的 time.Time 类型的值。limiter 就是这个定时器通道。

	limiter := time.Tick(200 * time.Millisecond)

	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	//  突发流量限流
	burstyLimiter := make(chan time.Time, 3)

	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	go func() {
		// 启动协程定时填充通道：
		// 启动一个匿名协程，使用 time.Tick(200 * time.Millisecond) 创建一个定时器，每隔 200 毫秒向
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	for req := range burstyRequests {
		// 每次从 burstyLimiter 通道接收一个值，由于 burstyLimiter 通道在开始时有 3 个值，所以前 3 个请求可以立即处理，之后会按照 200 毫秒的间隔处理，从而实现了突发流量限流。
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}
