package main

import (
	"fmt"
	"time"
)

func main() {

	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
		fallthrough //进行穿透
	default:
		fmt.Println("It's after noon")
	}

	// 使用匿名函数 func(i interface{})：这是一个匿名函数的定义，i是函数的参数名，interface{}表示空接口类型。
	// 在 Go 语言中，空接口可以表示任意类型的值，所以这个函数可以接收任意类型的参数。
	// go中没有java中的Object类，而使用 interface{} 空接口
	// switch是默认短路的，但可以使用 fallthrough 关键字进行穿透。但不能使用在`类型`判断的switch中
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}
