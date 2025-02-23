package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

// 状态协程
type readOp struct {
	key  int
	resp chan int
}
type writeOp struct {
	key  int
	val  int
	resp chan bool
}

func main() {

	var readOps uint64
	var writeOps uint64

	reads := make(chan readOp)
	writes := make(chan writeOp)

	go func() {
		var state = make(map[int]int)
		for {
			select {
			case read := <-reads:
				// 3 read通道收到数据， 从state读取对应key的值，放到对应 read的 resp 通道中
				read.resp <- state[read.key]
			case write := <-writes:
				// 3). write通道收到数据后，从state 中写入key对应的数据，
				state[write.key] = write.val
				// 4). 以write的 resp通道中放入 true 表示写入完成
				write.resp <- true
			}
		}
	}()

	for r := 0; r < 100; r++ {
		go func() {
			// 这个死循环会因为主协程的停止而中断，
			for {
				// 1. 构建 readOp
				read := readOp{
					key:  rand.Intn(5),
					resp: make(chan int)}
				// 2. 将readOp 发送到 read 通道
				reads <- read
				// 4. readOp的 resp 通道接收数据
				<-read.resp
				// 5. 给 readOps 计数变量 使用原始累加 1
				atomic.AddUint64(&readOps, 1)
				// 6. 睡眠1s 防止请求过于频繁
				time.Sleep(time.Millisecond)
			}
		}()
	}

	for w := 0; w < 10; w++ {
		go func() {
			for {
				// 1). 构建 writeOp
				write := writeOp{
					key:  rand.Intn(5),
					val:  rand.Intn(100),
					resp: make(chan bool)}
				// 2). 将 writeOp 发送到 writes 通道
				writes <- write
				// 5). writeOp的resp通道接收数据
				<-write.resp
				// 6). 对writeOps 计数变量 使用原始累加 1
				atomic.AddUint64(&writeOps, 1)
				// 7). 睡眠1s 防止请求过于频繁
				time.Sleep(time.Millisecond)
			}
		}()
	}

	time.Sleep(time.Second)

	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOps:", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOps:", writeOpsFinal)
}
