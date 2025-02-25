package main

import (
	"bytes"
	"fmt"
	"log"
	"os"

	"log/slog"
)

func main() {

	log.Println("standard logger")

	// log.LstdFlags 用于设置日志的输出格式。
	// log.Lmicroseconds 表示在时间中添加微秒
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	log.Println("with micro")
	// log.Lshortfile 表示在日志中添加文件名和行号
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("with file/line")

	// my: 是前缀
	mylog := log.New(os.Stdout, "my:", log.LstdFlags|log.Lshortfile)
	mylog.Println("from mylog")

	mylog.SetPrefix("ohmy:")
	mylog.Println("from mylog")

	// 调用将日志输出写入 buf
	var buf bytes.Buffer
	buflog := log.New(&buf, "buf:", log.LstdFlags)

	buflog.Println("hello")

	// from buflog:buf:2025/02/24 17:30:10 hello
	fmt.Print("from buflog:", buf.String())

	jsonHandler := slog.NewJSONHandler(os.Stderr, nil)
	myslog := slog.New(jsonHandler)
	/**
	{
	  "time": "2025-02-24T17:32:51.9012864+08:00",
	  "level": "INFO",
	  "msg": "hi there"
	}
	*/
	myslog.Info("hi there")
	/**
	{
	  "time": "2025-02-24T17:32:51.9012864+08:00",
	  "level": "INFO",
	  "msg": "hello again",
	  "key": "val",
	  "age": 25
	}
	*/
	myslog.Info("hello again", "key", "val", "age", 25)
}
