/**
 * object-oriented-example.go
 *
 * 一个基本的例子: 创建对象和方法
 */

package main

import "fmt"

// 定义 Dog 对象类型
type Dog struct {
	Name  string
	Color string
}

// 除了指定函数是哪个对象的，其余
// 函数名字，参数,返回类型跟普通函数定义一样
func (d Dog) Call() {
	fmt.Printf("Come here %s dog, %s \n", d.Color, d.Name)
}

// 接收对象也可以是指针格式,这样可以改变对象本身
func (d *Dog) Set(name, color string) {
	// 注意不能省'()'
	(*d).Color = color
	(*d).Name = name
}

func main() {

	// 创建Dog对象类型的实例，并且初始化
	Spot := Dog{Name: "Spot", Color: "brown"}

	// 调用对象方法
	Spot.Call()

	// 修改对象
	Spot.Set("C++", "pink")
	Spot.Call()

}
