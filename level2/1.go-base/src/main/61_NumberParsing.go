package main

import (
	"fmt"
	"strconv"
)

func main() {

	//  64 表示要整数位精度
	f, _ := strconv.ParseFloat("1.234", 64)
	fmt.Println(f)
	// 第二个参数 表示字符串的进制。当该参数为 0 时，函数会根据字符串的前缀来自动判断进制
	i, _ := strconv.ParseInt("123", 0, 64)
	fmt.Println(i)

	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println(d)

	u, _ := strconv.ParseUint("789", 0, 64)
	fmt.Println(u)
	// Atoi 是基本 base-10 的便捷函数 int 解析
	k, _ := strconv.Atoi("135")
	fmt.Println(k)

	_, e := strconv.Atoi("wat")
	fmt.Println(e)
}
