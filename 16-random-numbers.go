/**
 * random-numbers.go
 *
 * 示例用GO生成随机数
 */

package main

import (
	"fmt"
	"math/rand"
	"time"
)

type MathFunc func(int, int) int

func main() {

	// rand包函数 Intn(n), 返回一个[0, n)之间的随机整数
	// 然而如果没有指定生成器的种子，每次都会生成 81, 89, 47
	fmt.Println(rand.Intn(100))
	fmt.Println(rand.Intn(100))
	fmt.Println(rand.Intn(100))
	fmt.Println("----")

	// 使用当前时间的毫秒作为种子，用于随机数生成
	rand.Seed(time.Now().UnixNano())
	fmt.Println(rand.Intn(100))
	fmt.Println(rand.Intn(100))
	fmt.Println(rand.Intn(100))
	fmt.Println("---")

	// 下面显示一些其他的随机数函数
	// 更多请查阅官方文档: http://golang.org/pkg/math/rand/
	fmt.Println("Random Int:", rand.Int())
	fmt.Println("Random Float:", rand.Float32())
	fmt.Println("Random Permutation of N ints:", rand.Perm(8))

	// NormFloat64 函数返回一个正太分布的随机数
	// 它的平均值为0，标准差为1.
	// 如果调用次数够多，plot出来就会生成 贝尔曲线
	fmt.Println("Normalized:", rand.NormFloat64())

}
