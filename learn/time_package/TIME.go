package main

import (
	"fmt"
	"time"
)

//时间类型
func timeType() {
	//获取当前时间
	now := time.Now()
	year := now.Year()
	month := now.Month()
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()
	fmt.Printf("%d-%d-%d %d:%d:%d", year, month, day, hour, minute, second)
}

//时间戳
func timeStamp() {
	now := time.Now()
	//时间戳
	unix := now.Unix()
	//纳秒时间戳
	nano := now.UnixNano()
	fmt.Printf("当前时间戳为:%v\n", unix)
	fmt.Printf("纳秒时间戳为:%v\n", nano)
	//将时间戳转换为时间格式(默认不是纳秒级的，若要将纳秒级的时间戳转为正常时间则需要</1e9>)
	timer := time.Unix(nano/1e9, 0)
	fmt.Printf("对应时间为：", timer)
}

//时间操作
func timeOption() {
	addTime()
	subTime()
	equalTime()
	beforeTime()
	afterTime()
}

//时间相加(时间+时间间隔)
func addTime() {
	now := time.Now()
	later := now.Add(time.Hour)
	fmt.Println(later)
}

//求两个时间之间的差值
func subTime() {
	now := time.Now()
	before := now.Add(-time.Hour)
	fmt.Println(before)
	fmt.Println(now.Sub(before))
}

//时间是否相同
//判断两个时间是否相同，会考虑时区的影响，因此不同时区标准的时间也可以正确比较。
func equalTime() {
	now := time.Now()
	later := now.Add(-time.Hour)
	fmt.Println(now.Equal(later))
}

//是否在这个时间之前
func beforeTime() {
	now := time.Now()
	after := now.Add(time.Second)
	isBefore := now.Before(after)
	fmt.Println(isBefore)
}

//是否在这个时间之后
func afterTime() {
	now := time.Now()
	before := now.Add(-time.Second)
	isAfter := now.After(before)
	fmt.Println(isAfter)
}

//定时器
func timer() {
	//	定义一个时间间隔为1s的定时器
	tick := time.Tick(time.Second)
	for i := range tick {
		fmt.Println(i)
	}
	//	超时处理
	//after := time.After(time.Second * 3)
}

//时间格式化
func formatterTime() {
	now := time.Now()
	// 24小时制
	//时间类型有一个自带的方法Format进行格式化
	//Go语言中格式化时间模板不是常见的Y-m-d H:M:S而是使用Go的诞生时间2006年1月2号15点04分
	fmt.Println(now.Format("2006-01-02 15:04 Mon Jan"))
	// 12小时制
	fmt.Println(now.Format("2006-01-02 03:04:05.000 PM Mon Jan"))
	fmt.Println(now.Format("2006/01/02 15:04"))
	fmt.Println(now.Format("15:04 2006/01/02"))
	fmt.Println(now.Format("2006/01/02"))
}

//解析字符串格式的时间
func parseTime() {
	//	加载时区
	location, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println(err)
		return
	}
	parseTime, err := time.ParseInLocation("2006/01/02 15:04:05", "2020/02/01 22:02:13", location)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(parseTime)
}
func main() {
	//timeType()
	//timeStamp()
	//timeOption()
	//timer()
	//formatterTime()
	parseTime()
}
