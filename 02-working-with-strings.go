/**
 * working-with-strings.go
 *
 * 创建,操作字符串string类型。
 * 更多字符串操作都放在标准库"string"里面
 * See: http://golang.org/pkg/strings/
 */
package main

// 包含多个包，可以使用()这样的格式。
// 如果包含没有用到的包，编译器将报错
import (
	"fmt"     // 用于标准输出
	"strings" // 用于操作字符串
)

func main() {

	// 创建一个字符串变量
	str := "HI, I'M UPPER CASE"

	// 转换小写，Go中所有导出的函数、变量都要大写开头,采用驼峰命名，如 ‘.ToLower’
	lower := strings.ToLower(str)

	// 打印字符串
	fmt.Println(lower)

	// 检查字符串是否含有另一个字符串
	if strings.Contains(lower, "case") {
		fmt.Println("Yes, exists!")
	}

	// 打印前5个字符
	fmt.Println("First Five: " + str[:5])

	// 以指定字符或者单词，切割字符串
	sentence := "I'm a sentence made up of words"
	words := strings.Split(sentence, " ") // 返回 []string类型
	fmt.Printf("%v \n", words)

	// 如果是以空白字符(whitespace)切割，使用Fields方法比Split更好
	// Split(str, " "),以一个空格切割，而Fields以所有空白符切割
	fileds := strings.Fields(sentence)
	fmt.Printf("%v \n", fileds)

}

// 在终端运行本程序
// $go run working-with-strings.go
