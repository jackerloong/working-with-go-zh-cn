/**
 * control-structures.go
 * Go语言中的流程控制结构： if, else, switch
 */

package main

import "fmt"

func main() {

	// 定义一个int变量
	num := 1

	// Go 语言条件判断省去了圆括号"()", 大括号'{'必须而且左大括号不换行
	if num > 3 {
		fmt.Println("Many")
	}

	// Go 是挑剔的， "else"必须与括号"}"在同一行
	if num == 1 {
		fmt.Println("One")
	} else if num == 2 {
		fmt.Println("Two")
	} else {
		fmt.Println("Many")
	}

	// Switch 语句，不同于C，Go语言自动break
	switch {
	case num == 1:
		fmt.Println("One") // 打印之后自动break，不会执行后面的case
	case num == 2:
		fmt.Println("Two")
	case num > 2:
		fmt.Println("Many")
	default:
		fmt.Println("Thrown over boat")
	}

}
