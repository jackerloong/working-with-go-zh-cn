/**
 * working-with-slices.go
 *
 * Go 语言有数组类型，但是大部分接口都是针对slice（切片).
 * slice构建于数组之上，底层采用数组存储。
 * slice的类型声明为"[]T", T是一个类型
 * 如: "[]string" 是元素类型为string的slice, "[]int" 是元素类型为int的slice
 *
 * 更多slice查看: https://golang.org/ref/spec#Slice_expressions
 */

package main

import "fmt"

func main() {

	// 初始化一个空slice
	var empty []int

	// 初始化一个有数据的slice
	alphas := []string{"abc", "def", "ghi", "jkl"}

	// 向slice中追加元素
	empty = append(empty, 123)
	empty = append(empty, 456)
	fmt.Printf("%v \n", empty)

	// 追加多个值
	alphas = append(alphas, "pqr", "stu")
	fmt.Printf("%v \n", alphas)

	// 获取slice长度
	fmt.Println("Length: ", len(alphas))

	// 获取slice中一个元素
	fmt.Println(alphas[1])

	// 在slice上面创建新的slice, 获取区间为左闭右开"[1, 3)"
	alpha2 := alphas[1:3]
	fmt.Println(alpha2)

	// 检查数组切片中是否含有 该元素
	// 因为没用函数实现这个功能，自能自己通过遍历实现。
	if elemExists("def", alphas) {
		fmt.Println("Exists!")
	}

}

func elemExists(s string, a []string) bool {
	for _, v := range a {
		if v == s {
			return true
		}
	}
	return false
}
