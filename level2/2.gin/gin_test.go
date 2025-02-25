package gin

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"testing"
)

func TestGin(t *testing.T) {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}

func scanResponse(r *http.Response) {
	scanner := bufio.NewScanner(r.Body)
	for i := 0; scanner.Scan() && i < 5; i++ {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

func TestPing(t *testing.T) {
	resp, err := http.Get("http://localhost:8080/ping")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 5; i++ {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

func BenchmarkPing(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resp, err := http.Get("http://localhost:8080/ping")
		if err != nil {
			panic(err)
		}
		scanResponse(resp)
	}
}

func TestUser(t *testing.T) {
	resp, _ := http.Get("http://localhost:8080/user/foo")
	scanResponse(resp)
}

func BenchmarkUser(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resp, _ := http.Get("http://localhost:8080/user/foo")
		scanResponse(resp)
	}
}

func postAdmin() {
	url := "http://localhost:8080/admin"
	// 认证信息
	username := "foo"
	password := "bar"
	// 请求体
	requestBody := []byte(`{"value": "fo1o"}`)

	// 创建一个新的POST请求
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Println("创建请求失败:", err)
		return
	}

	// 设置请求头
	req.Header.Set("Content-Type", "application/json")

	// 设置基本认证
	req.SetBasicAuth(username, password)

	// 创建一个HTTP客户端
	client := &http.Client{}

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("发送请求失败:", err)
		return
	}
	defer resp.Body.Close()

	// 读取响应体
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取响应体失败:", err)
		return
	}
	fmt.Println(string(body))
}
func TestPostAdmin(t *testing.T) {
	postAdmin()
}

func BenchmarkPostAdmin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		postAdmin()
	}
}

func TestTemp(t *testing.T) {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{"title": "Main website"})
	})

	r.Run(":8080")
}
