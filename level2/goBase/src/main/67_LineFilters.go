package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	// 控制台输入
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {

		ucl := strings.ToUpper(scanner.Text())
		if ucl == "Q" {
			fmt.Println("已退出!")
			os.Exit(0)
		}
		fmt.Println(ucl)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
