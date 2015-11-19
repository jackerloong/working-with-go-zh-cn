/**
 * memcache-objects.go
 *
 * 如何用hoisie/redis包 缓存对象
 * 内存缓存将对象存储为[]byte格式，存储前需要编码和解码
 *
 * 本例使用JSON数据格式，比较适合人类阅读和调试
 * 如果有性能和空间方面的需求，可以查看"encoding/gob"包用其他格式
 */

package main

import (
	"encoding/json"
	"fmt"

	"github.com/hoisie/redis"
)

// 定义 Dog对象类型
type Dog struct {
	Name  string
	Color string
}

func main() {

	// 创建一个redis客户端
	var client redis.Client
	// 尝试从缓存中拿数据
	val, err := client.Get("Dog")

	// 检查缓存命中
	if err != nil {
		fmt.Println("Error fetching from", err)
	} else {
		fmt.Println("Cache hit")

		dog, err := DecodeData(val)
		if err != nil {
			fmt.Println("Error decoding data from memcache", err)
		} else {
			fmt.Println("Dog name is: ", dog.Name)
		}
	}

	// 创建对象实例
	spot := Dog{Name: "Spot", Color: "brown"}

	err = client.Set("Dog", EncodeData(spot))
	if err != nil {
		fmt.Println("Error setting memcache item", err)
	}

}

func DecodeData(raw []byte) (dog Dog, err error) {
	err = json.Unmarshal(raw, &dog)
	return dog, err
}

func EncodeData(dog Dog) (raw []byte) {
	enc, err := json.Marshal(dog)
	if err != nil {
		fmt.Println("Error encoding Action to JSON", err)
	}
	return enc
}

// 运行本程序两次，第一次显示未命中，第二次命中

// 运行本程序需要redis支持，确保你的主机安装了redis，然后启动redis服务
// $ redis-server
