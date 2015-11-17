/**
 * read-dir-recursive.go
 *
 * 递归读取目录，打印所有文件夹和文件
 * See: https://golang.org/pkg/path/filepath/
 */

package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	path := "./"

	// Go语言支持将函数作参数传递，也即回调
	// 函数Walk使用(path, function )参数，
	// 对每个path入口(目录/文件) 应用一次function
	filepath.Walk(path, Walker)
}

// Walker 函数被每个文件调用一次，接收filename，
// fileinfo 对象，如果出错还将有个error
func Walker(fn string, fi os.FileInfo, err error) error {
	if err != nil {
		fmt.Println("Walker Error: ", err)
		return nil
	}

	if fi.IsDir() {
		fmt.Println("Directory: ", fn)
	} else {
		fmt.Println("File: ", fn)
	}

	return nil
}
