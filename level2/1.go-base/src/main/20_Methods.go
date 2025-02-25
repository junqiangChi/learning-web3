package main

import "fmt"

// 结构体的方法
type rect struct {
	width, height int
}

// go 语言为了方便开发者编写代码，提供了语法糖。
// 当你使用指针接收者去访问结构体的属性时，Go 会自动解引用该指针。也就是说，r.width 实际上等价于 (*r).width
func (r *rect) area() int {
	return r.width * r.height
}

func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height: 5}

	fmt.Println("area: ", r.area())
	fmt.Println("perim:", r.perim())

	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim:", rp.perim())
}
