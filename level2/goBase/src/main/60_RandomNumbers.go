package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {

	fmt.Print(rand.IntN(100), ",")
	fmt.Print(rand.IntN(100))
	fmt.Println()

	// 默认0-1
	fmt.Println(rand.Float64())

	// 5-10 随机
	fmt.Print((rand.Float64()*5)+5, ",")
	fmt.Print((rand.Float64() * 5) + 5)
	fmt.Println()

	// PCG随机生成算法
	// 42 是随机数生成器的种子 种子是一个初始值
	// 1024 是流（stream）参数 加了随机数的多样性
	// 设置后产生随机数的就固定了

	s2 := rand.NewPCG(42, 1024)
	r2 := rand.New(s2)
	// 94,49
	fmt.Print(r2.IntN(100), ",")
	fmt.Print(r2.IntN(100))
	fmt.Println()

	// 94,49
	s3 := rand.NewPCG(42, 1024)
	r3 := rand.New(s3)
	fmt.Print(r3.IntN(100), ",")
	fmt.Print(r3.IntN(100))
	fmt.Println()
}
