package main

import (
	"bufio"
	"fmt"
	"os"
)

func check1(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	d1 := []byte("hello\ngo\n")
	// 0644 表示文件权限
	err := os.WriteFile("E:\\develop\\goprogram\\demo\\test.txt", d1, 0644)
	check1(err)

	f, err := os.Create("E:\\develop\\goprogram\\demo\\test1.txt")
	check1(err)

	defer f.Close()

	d2 := []byte{115, 111, 109, 101, 10} // some
	n2, err := f.Write(d2)
	check1(err)
	fmt.Printf("wrote %d bytes\n", n2)

	n3, err := f.WriteString("writes\n")
	check1(err)
	fmt.Printf("wrote %d bytes\n", n3)

	f.Sync()

	// 缓冲流
	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	check1(err)
	fmt.Printf("wrote %d bytes\n", n4)

	w.Flush()

}
