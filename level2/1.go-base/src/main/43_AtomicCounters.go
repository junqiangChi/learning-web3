package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {

	// 原子计数 用于在
	var ops atomic.Uint64

	// 等待所有协程处理完成
	var wg sync.WaitGroup

	// 使用普通变量，计算时多个协程会覆盖计算，所以值会比50000小
	a := 10
	for i := 0; i < 50; i++ {
		wg.Add(1)

		go func() {
			for c := 0; c < 1000; c++ {
				a += 1
				ops.Add(1)
			}

			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("a: ", a)
	fmt.Println("ops:", ops.Load())
}
