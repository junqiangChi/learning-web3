package main

import (
	"fmt"
	"math"
)

const s string = "constant"

func main() {
	fmt.Println(s)

	const n = 500000000

	const d = 3e20 / n
	fmt.Println(d)

	// 数值常量在给定 type 之前没有类型，例如通过显式转换
	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))

	const (
		A = 0
		B = 1
		C
	)
	fmt.Println(A, B, C)
}
