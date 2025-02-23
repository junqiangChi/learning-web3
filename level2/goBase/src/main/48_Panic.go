package main

import "os"

func main() {

	// panic 是 Go 语言中的一个内置函数，用于引发一个运行时错误，使程序进入恐慌状态。
	// 当调用 panic 时，程序会停止当前函数的正常执行流程，开始回溯调用栈，依次执行每个函数的 defer 语句，然后终止程序。
	panic("a problem")

	_, err := os.Create("/tmp/file")
	if err != nil {
		// 类似于java中的 throw
		panic(err)
	}
}
