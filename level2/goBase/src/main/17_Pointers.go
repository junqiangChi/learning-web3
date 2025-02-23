package main

import "fmt"

func zeroval(ival int) {
	ival = 0
}

// 修改指针对应内容值
func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	// 使用指针可以改变原值
	// 变量i iptr = &i表示i的指针，*iptr表示指针对应的变量
	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	fmt.Println("pointer:", &i)
}
