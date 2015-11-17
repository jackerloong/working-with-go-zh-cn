/**
 * sorting.go
 *
 * 示例如何排序数组和map
 * See: http://golang.org/pkg/sort/
 */

package main

import (
	"fmt"
	"sort"
)

func main() {

	// 创建两个slice， 用于字母序排序
	abc := []string{"jkl", "ghi", "abc", "def"}
	nums := []int{4, 2, 12, 5, 1, 3}

	// 显示排序前的slice
	fmt.Println("Raw ABC : ", abc)
	fmt.Println("Raw Nums: ", nums)

	// 用sort.Strings() 方法排序
	sort.Strings(abc)
	fmt.Println("Sorted ABC : ", abc)

	// 用 sort.Ints() 方法排序
	sort.Ints(nums)
	fmt.Println("Sorted Nums: ", nums)

	// 虽然abc是 []string类型，但不是StringSlice类型，尽管 "type StringSlice []string"
	// []string并没有对应的方法，而StringSlice类型实现了接口方法
	// 因此,abc需要转化为StringSlice类型才能被Reverse操作
	// 详情查看: https://golang.org/pkg/sort/#Reverse
	sort.Sort(sort.Reverse(sort.StringSlice(abc)))
	fmt.Println("Reverse ABC:", abc)

	//-------------------------
	// Sort Maps
	// ------------------------
	// sort map by key
	hash := map[string]int{
		"c": 3,
		"a": 1,
		"b": 2,
		"e": 5,
		"d": 4,
	}

	// 查看排序前的map
	for k, v := range hash {
		fmt.Printf("%s => %v \n", k, v)
	}

	// 创建一个数组，存储map的key
	keys := make([]string, 0, len(hash))
	for k, _ := range hash {
		keys = append(keys, k)
	}

	// 排序keys数组
	sort.Strings(keys)

	// 循环打印按key排序的map
	for i := range keys {
		fmt.Printf("%s => %v \n", keys[i], hash[keys[i]])
	}

	// 我们没法给内建的map类型排序，只能通过额外keys数组排序
}
