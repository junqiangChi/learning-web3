package main

import (
	"fmt"
	"sync"
)

type Container struct {
	// sync.Mutex 是 Go 语言标准库中用于实现互斥锁的类型
	// 互斥锁用于保护共享资源，确保同一时间只有一个协程（goroutine）可以访问该资源，从而避免数据竞争
	mu       sync.Mutex
	counters map[string]int
}

func (c *Container) inc(name string) {
	// c.mu.Lock()：调用 Lock 方法对互斥锁 mu 进行加锁操作。加锁后，其他协程如果尝试再次对该锁进行加锁操作，会被阻塞，直到当前协程释放该锁
	c.mu.Lock()
	// defer 关键字用于延迟执行 Unlock 方法，即当 inc 方法返回时，无论是否发生异常，都会自动调用 Unlock 方法来释放互斥锁
	defer c.mu.Unlock()
	c.counters[name]++
}

func main() {
	c := Container{
		counters: map[string]int{"a": 0, "b": 0},
	}

	var wg sync.WaitGroup

	doIncrement := func(name string, n int) {
		for i := 0; i < n; i++ {
			c.inc(name)
		}
		wg.Done()
	}

	wg.Add(3)
	go doIncrement("a", 10000)
	go doIncrement("a", 10000)
	go doIncrement("b", 10000)

	wg.Wait()
	fmt.Println(c.counters)
}
