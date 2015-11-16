/**
 * for-loops.go
 * GO只有"for"循环，没有 "while", "do-while" 循环
 * 但是"for"可以当"while"用
 */

package main

import "fmt"

func main() {

	// 创建一个字符串数组
	alphas := []string{"abc", "def", "ghi"}

	// 标准 for 循环, 与C相比，省去了园括号
	for i := 0; i < len(alphas); i += 1 {
		fmt.Printf("%d: %s \n", i, alphas[i])
	}

	// 使用迭代器遍历数组，
	// range 返回两个值， i -> index, val -> value
	for i, val := range alphas {
		fmt.Printf("%d: %s \n", i, val)
	}

	// 如果你只关注索引index，而不需要值value
	for i := range alphas {
		fmt.Println(i)
	}

	// 如果你不需要索引而只要取值，可以使用'_'
	// ’_'变量意味着不保存该值，
	for _, val := range alphas {
		fmt.Println(val)
	}

	// 像 "while"一样使用"for"
	x := 0
	for x < 10 {
		fmt.Println(x)
		x++
	}

}
