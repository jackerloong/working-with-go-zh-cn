/**
 * pointer.go
 *
 * 使用对象的指针去修改对象
 * See: http://tour.golang.org/#28
 */

package main

import "fmt"

// 定义Dog对象类型
type Dog struct {
	Name  string
	Color string
}

func main() {

	// 创建对象实例并且初始化
	Spot := Dog{Name: "Spot", Color: "brown"}

	// 获取对象的指针
	SpotPointer := &Spot

	// 通过指针修改字段
	SpotPointer.Color = "black"

	fmt.Println(Spot.Color)
}
