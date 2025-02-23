package main

import "fmt"

func vals() (int, int) {
	return 3, 7
}

func main() {

	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	// 忽略第一个值
	_, c := vals()
	fmt.Println(c)
}
