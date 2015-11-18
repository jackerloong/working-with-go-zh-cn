/**
 * regular-expressions.go
 *
 * 使用Go语言的正则表达式库
 * 查看包文档: http://golang.org/pkg/regexp
 */

package main

import (
	"fmt"
	"regexp"
)

func basic_regexes() {
	// 创建一个正则表达式模式
	// pattern， 匹配1个或者多个数字
	pattern := "[0-9]+"

	// test 字符串
	str := "The 12 monkeys ate 48 bananas"

	// 编译模式
	re, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Println("Error compiling regex", err)
	}

	// 1. 测试编译过后的模式，匹配一个字符串
	if re.MatchString(str) {
		fmt.Println("Yes, matched a number")
	} else {
		fmt.Println("No, no match")
	}

	// 2. 返回第一个匹配
	result_two := re.FindString(str)
	fmt.Println("Number matched:", result_two)

	// 3. 返回n个匹配到得值， -1表示返回所有结果
	// 本例返回2个
	result_three := re.FindAllString(str, 2)
	for i, v := range result_three {
		fmt.Printf("Match %d: %s\n", i, v)
	}

	// 4. 替换匹配的字符串
	result_four := re.ReplaceAllString(str, "xx")
	fmt.Println("Result:", result_four)
}

func case_insensitive() {
	// GO语言元字符串，里面的内容不会被转义
	// 如"\n", 是会被转义的，而`\n`不会，跟python的r'\n'一个意思
	ptn := `(?i)^t`
	fmt.Println(ptn)

	str := "To be or not to be"

	re, err := regexp.Compile(ptn)
	if err != nil {
		fmt.Println("Error compiling regex", err)
	}

	// 匹配字符串
	result := re.FindString(str)
	fmt.Println("Result:", result)
}

func main() {
	basic_regexes()
	case_insensitive()
}
