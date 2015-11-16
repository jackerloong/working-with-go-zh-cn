/**
 * functions.go
 * 函数使用示例
 */

package main

// 导入包
import (
	"fmt"  // 格式化标准输出
	"math" // 使用了一些math函数
)

// 接受一个参数的基本函数
// 需要指定参数类型，这里为string
// 函数格式为 function_name(variable type)
func Echo(s string) {
	fmt.Println(s)
}

// 有一个返回值得函数
// 返回值得类型指定，放在函数声明后面
func Say(s string) string {
	phrase := "Hello" + s
	return phrase
}

// 返回值命名变量的函数
// 你可以指定返回值变量的名字，在函数体中可以给返回值赋值
// 这样函数在返回时，可以省去return statement
// 变量的当前值将会随函数返回。
func Say2(s string) (phrase string) {
	phrase = "Hello" + s
	return
}

// 多参数和多返回值函数
// 如果多个变量类型相同，可以在最后一次性指定类型
func Divide(x, y float64) (q, r float64) {
	q = math.Trunc(x / y)
	r = math.Mod(x, y)
	return q, r
}

// main函数用来调用上面定义的函数
func main() {

	Echo("Bonjour tout le monde")

	fmt.Println(Say("Gopher"))

	// 测试多返回值函数
	q, r := Divide(11, 3)

	// 使用Printf格式化输出， %v可以被用于任意数据类型
	fmt.Printf("商: %v, 余数 %v \n", q, r)
}
