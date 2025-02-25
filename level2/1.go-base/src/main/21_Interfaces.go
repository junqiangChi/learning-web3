package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}

type aaa struct {
	a, b float64
}

type rect1 struct {
	width, height float64
}
type circle struct {
	radius float64
}

// 结构体struct 类似于java中的类class的概念
// 当把接口中所有的函数都实现，且必须是结构体方法，则认为实现了接口
func (r rect1) area() float64 {
	return r.width * r.height
}
func (r rect1) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func detectCircle(g geometry) {
	if c, ok := g.(circle); ok {
		fmt.Println("circle with radius", c.radius)
	}
}

func main() {
	r := rect1{width: 3, height: 4}
	c := circle{radius: 5}

	measure(r)
	measure(c)

	detectCircle(r)
	detectCircle(c)

	// 使用 int 类型调用 Add 函数
	sumInt := Add(10, 20)
	fmt.Println("Sum of ints:", sumInt)

	// 使用 float64 类型调用 Add 函数
	sumFloat := Add(10.5, 20.5)
	fmt.Println("Sum of floats:", sumFloat)
}
