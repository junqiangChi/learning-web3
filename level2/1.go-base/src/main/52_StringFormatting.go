package main

import (
	"fmt"
	"os"
)

type point struct {
	x, y int
}

func main() {

	p := point{1, 2}
	// %v 格式结构体的所有字段的值
	fmt.Printf("struct1: %v\n", p)
	// %+v 格式结构体的所有属性的值，并带有字段名称
	fmt.Printf("struct2: %+v\n", p)
	// %+v 格式结构体的所有属性的值，并带有字段名称和 即将产生该值的源代码片段
	fmt.Printf("struct3: %#v\n", p)
	// %T 格式该值的类型
	fmt.Printf("type: %T\n", p)
	// %t 格式bool类型数据
	fmt.Printf("bool: %t\n", true)
	// %d 格式整数类型数据
	fmt.Printf("int: %d\n", 123)
	// %b 将数据格式成二进制形式
	fmt.Printf("bin: %b\n", 14)
	// %c 格式成字符型数据
	fmt.Printf("char: %c\n", 33)
	// %x 格式成十六进制形式
	fmt.Printf("hex: %x\n", 456)
	// %f 格式浮点类型数据
	fmt.Printf("float1: %f\n", 78.9)
	// %e 格式成科学计数法形式
	fmt.Printf("float2: %e\n", 123400000.0)
	fmt.Printf("float3: %E\n", 123400000.0)
	// %s 格式字符串数据  "string"
	fmt.Printf("str1: %s\n", "\"string\"")
	// %q 格式成无转义的字符串类型 "\"string\""
	fmt.Printf("str2: %q\n", "\"string\"")

	fmt.Printf("str3: %x\n", "hex this")
	// %p 格式成指针形式
	fmt.Printf("pointer: %p\n", &p)
	// 格式化整数及小数部分的位数
	fmt.Printf("width1: |%6d|%6d|\n", 12, 345)
	fmt.Printf("width2: |%6.2f|%6.2f|\n", 1.2, 3.45)
	fmt.Printf("width3: |%-6.2f|%-6.2f|\n", 1.2, 3.45)
	fmt.Printf("width4: |%6s|%6s|\n", "foo", "b")
	fmt.Printf("width5: |%-6s|%-6s|\n", "foo", "b")

	// Sprintf 格式化后返回字符串
	s := fmt.Sprintf("sprintf: a %s", "string")
	fmt.Println(s)
	// Fprintf 格式化 + 打印到 io。Writer 以外的 操作系统
	fmt.Fprintf(os.Stderr, "io: an %s\n", "error")
}
