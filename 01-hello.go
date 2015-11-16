/**
 * 欢迎 Go Code!
 *
 * 通过阅读一系列Golang源代码学习go语言。
 *
 * hello.go
 * 第一个小程序
 */

// 所有程序必须是包的一部分，要生成可执行程序必须建立main包
package main

// 导入标准库中的包， 查看所有标准库 -> http://golang.org/pkg/
import "fmt"

// main包要求一个main()函数，作为程序的入口(执行起点)
func main() {

	// 定义一个变量，所有变量都有类型。这里使用 ':=' 让go编译器
	// 自动探测变量类型， 此处为 string 类型。
	phrase := "Hello World"

	// 调用fmt包中的Println方法，将变量输出到标准输出(stdout)并换行。
	fmt.Println(phrase)
}

// 为了在终端运行本程序， 需要你的电脑已经安装go环境。
// 安装过程参考: http://golang.org/doc/install,
// 国内上不去的话，参考: https://go-zh.org/doc/install
// 安装好了后，终端执行下面的命令即可运行。
// $ go run hello.go
