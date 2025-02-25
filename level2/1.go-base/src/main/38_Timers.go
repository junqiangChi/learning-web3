package main

import (
	"fmt"
	"time"
)

func main() {

	// 计时器
	timer1 := time.NewTimer(2 * time.Second)

	// <-timer1.C：这是一个阻塞操作，会一直等待，直到定时器 timer1 触发
	// 即 timer1.C 通道接收到一个值。一旦接收到值，程序会继续执行下一行代码。
	<-timer1.C
	fmt.Println("Timer 1 fired")

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()
	// 在第一个定时器等待期间，timer2被停止，因此 "Timer 2 fired" 不会打印 stop2 is false
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}

	time.Sleep(2 * time.Second)
}
