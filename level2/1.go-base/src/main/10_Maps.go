package main

import (
	"fmt"
	"maps"
)

func main() {

	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1:", v1)

	v3 := m["k3"]
	fmt.Println("v3:", v3)

	fmt.Println("len:", len(m))

	delete(m, "k2")
	fmt.Println("map:", m)

	clear(m)
	fmt.Println("map:", m)

	// 从 map 获取值时，可选的第二个返回值指示 map 中是否存在键。这可用于消除缺少的键和具有零值（如 0 或 “”）的键之间的歧义。这里我们不需要值本身，所以我们用空的 identifier_ 忽略了它。
	// 第一个参数表示返回的值，第二个参数表示这个键是否存在于map中，存在为true，反之为false
	// m["k2"] = 0 // 放开注释为 v11: 0 prs: true ， 否则是 v11: 0 prs: false
	v11, prs := m["k2"]
	fmt.Println("v11:", v11, "prs:", prs)
	fmt.Println("prs:", prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

	// 遍历判断
	n2 := map[string]int{"foo": 1, "bar": 2}
	if maps.Equal(n, n2) {
		fmt.Println("n == n2")
	}
}
