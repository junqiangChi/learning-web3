package main

import (
	"fmt"
	"sync"
	"time"
)

func worker2(id int) {
	fmt.Printf("Worker %d starting\n", id)

	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {

	var wg sync.WaitGroup

	// WaitGroup 用于等待所有在这里启动的 协程 完成
	for i := 1; i <= 5; i++ {
		// wg 变量，用于跟踪子协程的执行状态
		// wg.Add(1)：在启动每个子协程之前，调用 wg.Add(1) 方法，将 WaitGroup 的计数器加 1，表示有一个新的协程开始工作。
		wg.Add(1)

		go func() {
			//defer wg.Done()：在匿名协程内部，使用 defer 关键字确保 wg.Done() 方法在协程结束时被调用
			defer wg.Done()
			worker2(i)
		}()
	}

	wg.Wait()
}
