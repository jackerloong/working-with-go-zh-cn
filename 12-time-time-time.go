/**
 * time-time-time.go
 *
 * 如何格式化、匹配和计算日期跟时间
 * See: http://golang/pkg/time
 *
 */

package main

import (
	"fmt"
	"time"
)

func main() {

	// 创建 time 实例
	now := time.Now()
	epoch := now.Unix()

	fmt.Println("Now: ", now)
	fmt.Println("Unix Time: ", epoch)

	// 按指定的格式格式化时间，返回字符串输出
	// 你可以指定一个基本的格式， 如: "Mon, Jan 2, 2006 at 3:04pm"
	// 这样now.Format()将根据now的值，返回一个格式相同的字符串，如:
	// "Tue, Nov 17, 2015 at 10:09pm"
	fmt.Println(now.Format("Mon, Jan 2, 2006 at 3:04pm"))

	// 你可以抓取时间对象里面的任何一个部分
	// 更多参考: http://golang.org/pkg/time/#Time
	fmt.Println("Year: ", now.Year())
	fmt.Println("Month: ", now.Month())

	// 几个用于时间格式化的内建常量,
	// time.RFC850 = "Monday, 02-Jan-06 15:04:05 MST"
	// 查看更多时间常量: http://golang.org//pkg/time/#pkg-constants
	fmt.Println(now.Format(time.RFC850))
	fmt.Println(now.Format(time.RFC1123))

	// 指定一个特定日期，需要连同时区一起指定
	// 你可以通过 LoadLocation建立一个时区
	// 也可以使用 time.UTC 常量
	est, _ := time.LoadLocation("EST")
	july4 := time.Date(1776, 7, 4, 12, 15, 0, 0, est)

	fmt.Println("July 4, 1776 was a ", july4.Format("Monday"))

	// 时间比较函数，Before(), After(), Equal()
	if july4.Before(now) {
		fmt.Println("July 4 is before Now ")
	}

	// 时间加减法
	// 返回一个 Duration对象, 实质就是一个int64,
	diff := now.Sub(july4)

	// 转换为大概的天数
	days := int(diff.Hours() / 24)
	fmt.Printf("July 4 was about %d days ago \n", days)

	// 可以通过Duration增加日期
	twodaysDiff := time.Hour * 24 * 2
	twodays := now.Add(twodaysDiff)
	fmt.Println("Two Days: ", twodays.Format(time.ANSIC))

	// 匹配日期
	// 可以通过预设的格式匹配用户输入的日期
	input_form := "1/2/2006 3:04pm"
	user_str := "4/16/2014 11:38am"
	user_date, err := time.Parse(input_form, user_str)
	if err != nil {
		fmt.Println(">>> Error parsing date string")
	}
	fmt.Println("User Date: ", user_date.Format("Jan 2, 2006 @ 3:04pm"))

}
