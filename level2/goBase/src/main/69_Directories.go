package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func check2(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	// 0755 为目录权限
	err := os.Mkdir("subdir", 0755)
	check2(err)

	defer os.RemoveAll("subdir")

	// 写空文件
	createEmptyFile := func(name string) {
		d := []byte("")
		check2(os.WriteFile(name, d, 0644))
	}

	createEmptyFile("subdir/file1")

	// 类似于mkdir -p
	err = os.MkdirAll("subdir/parent/child", 0755)
	check2(err)

	createEmptyFile("subdir/parent/file2")
	createEmptyFile("subdir/parent/file3")
	createEmptyFile("subdir/parent/child/file4")

	// 列出目录下所有内容
	c, err := os.ReadDir("subdir/parent")
	check2(err)

	fmt.Println("Listing subdir/parent")
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	// 切换目录 相当于cd
	err = os.Chdir("subdir/parent/child")
	check2(err)

	c, err = os.ReadDir(".")
	check2(err)
	fmt.Println(c)

	fmt.Println("Listing subdir/parent/child")
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	err = os.Chdir("../../..")
	check2(err)

	fmt.Println("Visiting subdir")
	err = filepath.WalkDir("subdir", visit)
}

func visit(path string, d fs.DirEntry, err error) error {
	if err != nil {
		return err
	}
	fmt.Println(" ", path, d.IsDir())
	return nil
}
