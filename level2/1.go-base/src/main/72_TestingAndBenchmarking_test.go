package main

import (
	"fmt"
	"testing"
)

// go test xxx_test.go 执行指定的test文件
// go test 执行当前目录下对应包的所有test文件中的所有单元测试及基准测试
// go test -v xxx_test.go 指定执行测试文件中的所有单元测试及基准测试
// go test -run=functionName 指定单元测试函数名
// go test -bench=functionName 指定基准测试函数名
// go test -run=functionName  xxx_test.go 指定测试文件及测单元试函数
// -v 是用来输出详细的测试结果
// 文件名必须以 _test.go 结尾
// 函数名以Test开始表示单元测试 以Benchmark开头表示基准测试

/*
*
单元测试：主要目的是验证代码的功能正确性
基准测试：主要目的是评估代码的性能
*/
func IntMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func TestIntMinBasic(t *testing.T) {
	ans := IntMin(2, -2)
	fmt.Println("sdfadfadsfasfdas")
	if ans != -2 {

		t.Errorf("IntMin(2, -2) = %d; want -2", ans)
	}
}

func TestIntMinTableDriven(t *testing.T) {
	var tests = []struct {
		a, b int
		want int
	}{
		{0, 1, 0},
		{1, 0, 0},
		{2, -2, -2},
		{0, -1, -1},
		{-1, 0, -1},
	}

	for _, tt := range tests {

		testname := fmt.Sprintf("%d,%d", tt.a, tt.b)
		t.Run(testname, func(t *testing.T) {
			ans := IntMin(tt.a, tt.b)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})

	}
}

func BenchmarkIntMin(b *testing.B) {

	for i := 0; i < b.N; i++ {
		IntMin(1, 2)
	}
}
