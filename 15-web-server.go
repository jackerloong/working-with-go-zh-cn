/**
 * web-server.go
 *
 * 一个web服务器程序，标准库net/http提供了丰富的工具用于创建web服务器
 *
 * 更多http可查看作者 Lanyou 项目: https://github.com/mkaz/lanyon
 * 一个转换markdown到html得服务
 */

package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	// 关联URL与处理该URL的函数
	http.HandleFunc("/hello", helloRequest)
	http.HandleFunc("/", getRequest)

	// 开始web服务
	log.Println("Listening on http://localhost:9999/")
	log.Fatal(http.ListenAndServe(":9999", nil))

}

// 处理 "/hello"请求，基本响应
func helloRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, Welcom to here")
	return
}

// 基本的文件服务响应，根据URL请求的路径
// 显示请求的文件或文件索引
func getRequest(w http.ResponseWriter, r *http.Request) {
	file_requested := "./" + r.URL.Path
	http.ServeFile(w, r, file_requested)
	return
}
