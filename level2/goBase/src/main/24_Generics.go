package main

import "fmt"

// SlicesIndex 泛型
/**
E：这是一个类型参数，它有一个约束 comparable。comparable 是 Go 语言中的一个预定义约束，
它表示该类型的值可以使用 == 和 != 运算符进行比较。常见的可比较类型包括基本数据类型（如 int、float64、string 等）、指针类型、数组类型等。
S：这也是一个类型参数，它的约束是 ~[]E。~ 符号在 Go 泛型中表示底层类型约束，~[]E 意味着 S 可以是任何底层类型为 []E 的切片类型。
也就是说，S 可以是普通的切片类型，也可以是自定义的基于切片类型的别名。例如：
*/
func SlicesIndex[S ~[]E, E comparable](s S, v E) int {
	for i := range s {
		if v == s[i] {
			return i
		}
	}
	return -1
}

type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val  T
}

func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

func (lst *List[T]) AllElements() []T {
	var elems []T
	for e := lst.head; e != nil; e = e.next {
		elems = append(elems, e.val)
	}
	return elems
}

// 泛型类型
// Stack 是一个泛型结构体，用于实现栈数据结构
type Stack[T any] []T

// Push 方法用于将元素压入栈中
func (s *Stack[T]) Push(item T) {
	*s = append(*s, item)
}

// 泛型函数
// Max 是一个泛型函数，用于返回两个同类型元素中的最大值
func Max[T int | float64](a, b T) T {
	if a > b {
		return a
	}
	return b
}

// 泛型约束
type Addable interface {
	int | int64 | float32 | float64
}

func Add[T Addable](a, b T) T {
	return a + b
}

func main() {
	var s = []string{"foo", "bar", "zoo"}

	fmt.Println("index of zoo:", SlicesIndex(s, "zoo"))

	// 显式声明
	_ = SlicesIndex[[]string, string](s, "zoo")

	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)
	fmt.Println("list:", lst.AllElements())
}
