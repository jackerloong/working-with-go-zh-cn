/**
 * guess-game.go
 *
 * 猜一个1-100的数字
 */

package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 当程序启动时，init()函数会被自动调用
func init() {
	// 新的随机数种子
	rand.Seed(time.Now().UnixNano())
}

func main() {

	// 用于存放user所猜测的值, 来自标准输入
	var guess int

	// 用于统计所猜次数的变量
	var count int

	// 随机生成一个要猜得数, 加1是因为区间[0, 100)
	num := rand.Intn(100) + 1

	for {
		// 提示用户输入
		fmt.Print("Guess: ")
		_, err := fmt.Scanf("%d", &guess)
		if err == nil {
			count += 1 //递增猜测次数
			if guess > num {
				fmt.Println(" >> Too high")
			} else if guess < num {
				fmt.Println(" >> Too low")
			} else {
				fmt.Printf("Correct! It took you %d guesses!\n", count)
				break
			}
		} else { // 如果输入的数据出错，比如输入的不是一个数字
			fmt.Println(">> Please input a number")
		}
	}
}
