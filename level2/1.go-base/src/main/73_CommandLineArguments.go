package main

import (
	"fmt"
	"os"
)

// .\bin\CommandLineArguments.exe 1 2 3 4 5
func main() {

	argsWithProg := os.Args
	// os.Args[1:] 也就是 os.Args[0] 保存了程序全部的参数。
	argsWithoutProg := os.Args[1:]
	// 索引位从1开始
	arg := os.Args[3] // 3

	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg)
}
