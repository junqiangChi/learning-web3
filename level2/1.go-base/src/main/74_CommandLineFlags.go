package main

import (
	"flag"
	"fmt"
)

// 对命令行main的传参进行命名
// go run CommandLineFlags.go --word=aa --numb=111 --fork=true
// go run CommandLineFlags.go -word=aa -numb=111 -fork=true
// go run CommandLineFlags.go -h 查看帮助
func main() {
	// 仅支持字符串、整数和布尔值选项
	wordPtr := flag.String("word", "foo", "a string")

	numbPtr := flag.Int("numb", 42, "an int")
	forkPtr := flag.Bool("fork", false, "a bool")

	// 用已有的参数来声明一个标志
	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")

	flag.Parse()

	fmt.Println("word:", *wordPtr)
	fmt.Println("numb:", *numbPtr)
	fmt.Println("fork:", *forkPtr)
	fmt.Println("svar:", svar)
	fmt.Println("tail:", flag.Args())
}
