package main

import "fmt"

func mayPanic() {
	panic("a problem")
}

// recover 可以阻止 panic 中止程序，并让它继续执行

func main() {

	// recover 必须在 Deferred 函数中调用，也就是使用 defer 关键字。当封闭函数 panic 时，延迟激活，其中的 recover 调用将捕获 panic
	defer func() {
		if r := recover(); r != nil {

			fmt.Println("Recovered. Error:\n", r)
		}
	}()
	// 会先执行这个
	mayPanic()

	fmt.Println("After mayPanic()")
}
