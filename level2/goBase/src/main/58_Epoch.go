package main

import (
	"fmt"
	"time"
)

func main() {

	now := time.Now()
	fmt.Println(now)

	// 秒级时间戳
	fmt.Println(now.Unix())
	// 毫秒级时间戳
	fmt.Println(now.UnixMilli())
	// 纳秒级时间戳
	fmt.Println(now.UnixNano())

	// 时间戳格式化
	fmt.Println(time.Unix(now.Unix(), 0))
	fmt.Println(time.Unix(0, now.UnixNano()))
}
