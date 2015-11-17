/**
 * read-write-file-two.go
 *
 * 前一个io/ioutil例子展示了一次性读取整个文件，有时这是不需要的。
 * 本例将使用bufio按行读写文件
 */

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	filepath := "./extras/rabbits.txt"

	// os包提供不依赖操作系统平台的系统调用
	f, _ := os.Open(filepath)

	// defer 指定的语句或函数，不管程序是否发生异常，都将在当前函数结束后执行。
	// 对于打开的文件， defer f.Close关闭是一个好习惯
	defer f.Close()

	// 使用 带缓冲io 读取文件
	// 更多缓冲IO方法参考, http://golang.org/pkg/bufio/
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}

	// 错误检查， if语句先执行初始化，后判断
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// 创建一个新文件
	f, err := os.Create("output.txt")
	if err != nil {
		log.Fatalln("Error creating file: ", err)
	}
	defer f.Close()

	for _, str := range []string{"apple", "banana", "carrot"} {
		bytes, err := f.WriteString(str)
		if err != nil {
			log.Fatalln("Error writing string: ", err)
		}
		fmt.Printf("Wrote %d bytes to file\n", bytes)
	}

}
