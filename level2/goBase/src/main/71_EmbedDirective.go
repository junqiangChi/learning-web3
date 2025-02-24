package main

import (
	"embed"
	"fmt"
)

//嵌入指令 相当于java 注解

// 这个目录文件不能有../ 必须是同级目录的文件
//
//go:embed folder/single_file.txt
var fileString string

//go:embed folder/single_file.txt
var fileByte []byte

// 嵌入多个文件
// embed.FS 类型，它实现了一个简单的虚拟文件系统。
//
//go:embed folder/single_file.txt
//go:embed folder/*.hash
var folder embed.FS

func main() {

	println(fileString)
	fmt.Println(fileString)
	print(string(fileByte))

	content1, _ := folder.ReadFile("folder/file1.hash")
	print(string(content1))

	content2, _ := folder.ReadFile("folder/file2.hash")
	print(string(content2))
}
