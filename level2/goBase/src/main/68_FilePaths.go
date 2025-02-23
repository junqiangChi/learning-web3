package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main() {

	// 路径拼接
	p := filepath.Join("dir1", "dir2", "filename")
	fmt.Println("p:", p)

	fmt.Println(filepath.Join("dir1//", "filename"))       // dir1\filename
	fmt.Println(filepath.Join("dir1/../dir1", "filename")) // dir1\filename

	// 目录
	fmt.Println("Dir(p):", filepath.Dir(p)) // Dir(p): dir1\dir2
	// 文件名
	fmt.Println("Base(p):", filepath.Base(p)) // Base(p): filenam

	fmt.Println(filepath.IsAbs("dir/file"))
	fmt.Println(filepath.IsAbs("/dir/file"))

	filename := "config.json"

	// 获取文件扩展名
	ext := filepath.Ext(filename)
	fmt.Println(ext)
	// 获取无扩展名的文件名
	fmt.Println(strings.TrimSuffix(filename, ext))

	rel, err := filepath.Rel("a/b", "a/b/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)

	// 以"a/b" 为基准路径，从"a/c/t/file"中获取相对路径为 ..\c\t\file
	rel, err = filepath.Rel("a/b", "a/c/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)
}
