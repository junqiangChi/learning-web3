package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	dat, err := os.ReadFile("E:\\develop\\goprogram\\demo\\test.txt")
	check(err)
	fmt.Print(string(dat))

	f, err := os.Open("E:\\develop\\goprogram\\demo\\test.txt")
	check(err)
	// 从文件的开头读取一些字节。最多允许读取 5 个，但也要注意实际读取了多少。
	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	check(err)
	fmt.Println()
	fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))

	o2, err := f.Seek(6, io.SeekStart)
	check(err)
	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes @ %d: ", n2, o2)
	fmt.Println()
	fmt.Printf("%v\n", string(b2[:n2]))

	// 从当前位置开始查找
	_, err = f.Seek(2, io.SeekCurrent)
	check(err)

	// 从结尾开始查找
	_, err = f.Seek(-4, io.SeekEnd)
	check(err)

	o3, err := f.Seek(6, io.SeekStart)
	check(err)
	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(f, b3, 2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

	_, err = f.Seek(0, io.SeekStart)
	check(err)

	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(5)
	check(err)
	fmt.Printf("5 bytes: %s\n", string(b4))

	err1 := f.Close()
	if err1 != nil {
		return
	}
}
