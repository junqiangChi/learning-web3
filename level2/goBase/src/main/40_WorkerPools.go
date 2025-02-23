package main

import (
	"fmt"
	"time"
)

/*
*
id：工作协程的编号，用于区分不同的工作协程。
jobs <-chan int：只读通道，用于接收任务。<-chan 表示这是一个只能接收数据的通道。
results chan<- int：只写通道，用于发送处理结果。chan<- 表示这是一个只能发送数据的通道。
*/
func worker1(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func main() {

	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// 启用3 个协程
	// 向通道发5份数据，3个协程会随机接收通道的5份数据进行处理
	for w := 1; w <= 3; w++ {
		go worker1(w, jobs, results)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= numJobs; a++ {
		<-results
	}
}
