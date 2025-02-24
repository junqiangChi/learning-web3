package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func scanResponse(r *http.Response) {
	scanner := bufio.NewScanner(r.Body)
	for i := 0; scanner.Scan() && i < 5; i++ {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

func main() {

	resp, err := http.Get("https://gobyexample.com")
	if err != nil {
		panic(err)
	}

	scanResponse(resp)
	resp1, _ := http.Get("https://yuanbao.tencent.com/")
	scanResponse(resp1)
	defer resp.Body.Close()

	fmt.Println("Response status:", resp.Status)

}
