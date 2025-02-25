package main

import (
	"fmt"
	"slices"
)

func main() {

	var ss [3]string //数组
	var s []string   //切片
	fmt.Println(ss)
	fmt.Println("uninit:", s, s == nil, len(s) == 0)

	s = make([]string, 3) // make 设置容量
	// cap(s) 查看容量
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	// 创建空切片
	c := make([]string, len(s))
	// 从s复制到c
	copy(c, s)
	fmt.Println("cpy:", c)

	l := s[2:5] // 取第3位到第5位的元素 s[2] s[3] s[4]
	fmt.Println("sl1:", l)

	l = s[:5] // 取开始到第5位元素 s[0] - s[4]
	fmt.Println("sl2:", l)

	l = s[2:] // 取第3位到结束的元素 s[3] - s[len(s) -1]
	fmt.Println("sl3:", l)

	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	t2 := []string{"g", "h", "i"}
	// 使用Equal对于元素内容
	if slices.Equal(t, t2) {
		fmt.Println("t == t2")
	}

	// 二维切片
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
	slice := make([]int, 3, 5) //长度为3，容量为5
	fmt.Println("len:", len(slice), "cap:", cap(slice))
	fmt.Println("slice:", slice)

	ints := make([]int, 2)
	ints[0] = 1
	fmt.Println("len:", len(ints), "cap:", cap(ints))

}
