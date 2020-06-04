package main

import (
	"fmt"
	"time"
)

/**
time 包
1年 = 356天day
1天 = 24小时 hour
1小时 = 60分钟 minute
1分钟 = 60 秒 second
1秒 = 1000毫秒 millisecond

*/
func main() {
	// 2020-06-04 14:35:18.141997 +0800 CST m=+0.000145006
	t1 := time.Now()
	timeNow := time.Now()
	fmt.Println(timeNow)
	fmt.Println(t1.Unix())

	fmt.Println(t1)
	// 获取指定时间
	t2 := time.Date(2020, 6, 4, 13, 30, 23, 0, time.Local)
	fmt.Println(t2)

	// 时间 和时间字符串直接的转换 时间格式化的日期必须格式 2006年1月2日 15:04:05
	s1 := t1.Format("2006年1月2日 15:04:05")
	fmt.Println(s1)
	s2 := t1.Format("2006/01/02")
	fmt.Println(s2)

	s3 := "2020年10月10日"
	t3, err := time.Parse("2006年01月02日", s3)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(t3)
	fmt.Printf("%T\n", t3)

	fmt.Println(t1.String())
	fmt.Println(t1.Date())  // 年月日
	fmt.Println(t1.Clock()) // 时分秒

	fmt.Println(t1.Year())    // 年
	fmt.Println(t1.YearDay()) // 今年一共过了多少天
	fmt.Println(t1.Month())   // 月
	fmt.Println(t1.Day())     //日
	fmt.Println(t1.Hour())    // 时
	fmt.Println(t1.Weekday()) // 星期几
	fmt.Println(t1.ISOWeek()) // 今年的第几周
	fmt.Println(t1.Unix())    // 当前时间的时间戳
	fmt.Println("-------")
	// 时间间隔
	t5 := t1.Add(24 * time.Hour)
	fmt.Println(t1)
	fmt.Println(t5)
	t6 := t1.AddDate(0, 0, -1)
	fmt.Println(t6)
	h, _ := time.ParseDuration("2h")
	t7 := t1.Add(h)
	fmt.Println(t1)
	fmt.Println(t7)
	negativeM, _ := time.ParseDuration("-2m")
	t8 := t1.Add(negativeM)
	fmt.Println(t1)
	fmt.Println(t8)
	fmt.Println(h)
	fmt.Println("1年3个月4天之前的时间:", t1.AddDate(-1, -3, -4))
	// 睡眠
	time.Sleep(3 * time.Second)
	fmt.Println("睡眠之后的结果")

}
