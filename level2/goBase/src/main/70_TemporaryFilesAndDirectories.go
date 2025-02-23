package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func check3(e error) {
	if e != nil {
		panic(e)
	}
}

// 临时文件和目录
func main() {

	f, err := os.CreateTemp("", "sample")
	check3(err)

	fmt.Println("Temp file name:", f.Name())

	// 延迟触发
	defer os.Remove(f.Name())

	_, err = f.Write([]byte{1, 2, 3, 4})
	check3(err)

	dname, err := os.MkdirTemp("", "sampledir")
	check3(err)
	fmt.Println("Temp dir name:", dname)

	// 延迟触发
	defer os.RemoveAll(dname)

	fname := filepath.Join(dname, "file1")
	err = os.WriteFile(fname, []byte{1, 2}, 0666)
	check3(err)
}
