/**
 * command-line-args.go
 *
 * 一个展示命令行参数使用的Go程序
 */

package main

import (
	"flag"
	"fmt"
	"os"
)

var str string
var num int
var help bool

func main() {

	// Go语言main函数没有参数， 命令行参数以slice类型存储在os.Args中
	// 第一个参数是 程序自己的名称
	num_args := len(os.Args)

	// 检测是否收到了命令行参数
	if num_args < 2 {
		fmt.Println(">> No args passed in")
	}

	// flag包提供了命令行参数分析匹配功能
	// 这个例子中创建了三个全局变量，然后将指针传给对应的类型函数
	flag.StringVar(&str, "str", "defaut value", "text description")
	flag.IntVar(&num, "num", 5, "text description")
	flag.BoolVar(&help, "help", false, "Display help")
	flag.Parse()

	// 检测 'help' 是否被显示指定
	if help {
		fmt.Println(">> Display help screen")
		os.Exit(1)
	}

	fmt.Println(">> String:", str)
	fmt.Println(">> Number:", num)

	// 最后一个命令行参数
	fmt.Println(">> Last Item:", os.Args[num_args-1])

	// os.Args将包含所有参数包括flags，例如
	// 执行 $ go run command-line-args.go --str=Foo filename
	// os.Args[1] = "--str=Foo"

	// 没有匹配到得参数，都将保存在 flag.Args()中
	// 例如: $ go run command-line-args.go --str=Foo filename a=b
	// 则: args = []string{"filename", "a=b"}
	args := flag.Args()
	if len(args) > 0 {
		fmt.Printf("Flag Arg: %v \n", args)
	}

}

// 测试输出
// $ go run command-line-args.go
// >> No args passed in
// >> String: default value
// Number: 5
// Last Item: /var/.../11-command-line-args.go

// $ go run command-line-args.go --str=Foo --num=8 filename
// >> String: Foo
// >> Number: 8
// >> Last Item: filename
// >> Flag Arg: [filename]
