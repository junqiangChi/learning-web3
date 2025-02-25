package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func main() {

	match, err := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)
	if err != nil {
		fmt.Println(err)
	}

	// 定义正则变量
	r, _ := regexp.Compile("p([a-z]+)ch")

	fmt.Println(r.MatchString("peach")) //true

	fmt.Println(r.FindString("peach punch")) // peach

	// 查找第一个匹配字符的开始和结束索引 [0, 5]
	fmt.Println("idx:", r.FindStringIndex("peach punch"))
	// 查找n个匹配的字符的开始和结束索引 [[0 5] [6 11]]
	fmt.Println(r.FindAllStringIndex("peach punch", 2))
	// 查找匹配字符及子字符 [peach ea]
	fmt.Println(r.FindStringSubmatch("peach punch"))
	// 查找第一个匹配字符及子字符起始位置 [0 5 1 3]
	fmt.Println(r.FindStringSubmatchIndex("peach punch"))
	// 查找所有匹配的字符串 [peach punch pinch]
	fmt.Println(r.FindAllString("peach punch pinch", -1))
	// 查找所有匹配及子字符串的开始和结束索引  [[0 5 1 3] [6 11 7 9] [12 17 13 15]]
	fmt.Println("all:", r.FindAllStringSubmatchIndex(
		"peach punch pinch", -1))
	// 查找前n个匹配字符串
	fmt.Println(r.FindAllString("peach punch pinch", 2))
	// 字节匹配
	fmt.Println("byte: ", r.Match([]byte("peach")))

	// MustCompile 不检查error
	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println("regexp:", r)

	// 用于将字符串子集替换为其他值
	fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))

	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))
}
