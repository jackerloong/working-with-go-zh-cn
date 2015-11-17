/**
 * json.go
 *
 * go语言操作json数据格式，Marshal()将结构体封装成JSON，Unmarshal则相反
 * 是否注意到一个趋势，这些都是标准库的内容
 * See: http://golang.org/pkg/encoding/json/
 */

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// 创建一个结构体，用于匹配json格式
type Person struct {
	Name, City string
}

// 创建一个Person实例
var person Person

// 创建一个Person类型数组
var people []Person

func main() {

	// 用于匹配的JSON 字符串
	json_str := "{ \"name\": \"Jacker\", \"City\":\"GuangZhou\"}"

	// JSON Unmarshal方法接受 []byte类型，因此需要类型转换
	// 第二个参数为一个指针，指向存放的结构体对象
	if err := json.Unmarshal([]byte(json_str), &person); err != nil {
		fmt.Println("Error parsing JSON: ", err)
	}

	// 输出结果
	fmt.Printf("Name: %v, City: %v\n", person.Name, person.City)

	// 从文件读json数据
	file, err := ioutil.ReadFile("./extras/names.json")
	if err != nil {
		fmt.Println("Error reading file")
	}

	// 文件names.json里面含有一列person对象，因此使用 数组people保存结果
	if err := json.Unmarshal(file, &people); err != nil {
		fmt.Println("Error parsing JSON", err)
	}

	// 输出结果
	fmt.Println(people)

	// 编码 Go对象为JSON数据
	json, err := json.Marshal(people)
	if err != nil {
		fmt.Println("JSON Encoding Error", err)
	}
	fmt.Println(string(json))
}
