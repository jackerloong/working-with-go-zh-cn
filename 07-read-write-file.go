/**
 * read-write-files.go
 *
 * 使用 io/ioutils 包读写文件
 * See: http://golang.org/pkg/io/ioutil/
 */

package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {

	filename := "./extras/rabbits.txt"

	// 读文件，将文件所有内容放入content变量中
	// 如果出错，将错误放在第二个返回值中
	// 如果没有出错，err将会是nil
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		// log 是一个很小巧实用的工具，它可以输出信息
		// 在本例中，一个致命的log还将终止本程序
		log.Fatalln("Error reading file", filename)
	}

	// 返回的content 是 []byte类型，而不是字符串string类型
	// 因此，必须转换为string再打印
	fmt.Println(string(content))

	// 将内容写入一个新文件
	outfile := "outfile.txt"
	err = ioutil.WriteFile(outfile, content, 0644) //0644为Unix文件权限格式
	if err != nil {
		log.Fatalln("Error writing file: ", err)
	}
}
