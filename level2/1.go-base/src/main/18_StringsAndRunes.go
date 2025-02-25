package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	const s = "สวัสดี"

	// len函数是返回字节长度而不是字符长度
	fmt.Println("Len:", len(s))

	ss := "1" // 1个字节
	fmt.Printf("字节长度: %d, 字符数量: %d\n", len(ss), utf8.RuneCountInString(ss))
	ss = "池" // 3个字节
	fmt.Printf("字节长度: %d, 字符数量: %d\n", len(ss), utf8.RuneCountInString(ss))
	ss = "😀" //4个字节
	fmt.Printf("字节长度: %d, 字符数量: %d\n", len(ss), utf8.RuneCountInString(ss))

	for i := 0; i < len(ss); i++ {
		// 十六进制打印所有字节的ascii值
		fmt.Printf("%x ", ss[i])
	}
	fmt.Println()

	fmt.Println("Rune count:", utf8.RuneCountInString(s))

	// %#U 格式化输出unicode值
	for idx, runeValue := range s {
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
	}

	fmt.Println("\nUsing DecodeRuneInString")
	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at %d\n", runeValue, i)
		w = width

		examineRune(runeValue)
	}
}

func examineRune(r rune) {

	if r == 't' {
		fmt.Println("found tee")
	} else if r == 'ส' {
		fmt.Println("found so sua")
	}
}
