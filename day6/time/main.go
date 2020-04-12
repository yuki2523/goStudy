package main

import (
	"fmt"
	"time"
)

// days 获取当前时间,time.Time的方法:Year,Month,Day,Weekday,Hour,Minute,Second
func days() {
	// 一般时间,time.Weekday其实就是int的别名
	now := time.Now()
	fmt.Printf("%v,%T\n", now, now)                      // 2020-04-12 09:16:14.8571809 +0800 CST m=+0.002991601,time.Time
	fmt.Printf("%#v,%T\n", now.Year(), now.Year())       // 2020,int
	fmt.Printf("%#v,%T\n", now.Month(), now.Month())     // 4,time.Month
	fmt.Printf("%#v,%T\n", now.Day(), now.Day())         // 12,int
	fmt.Printf("%#v,%T\n", now.Weekday(), now.Weekday()) // 0,time.Weekday
	fmt.Printf("%#v,%T\n", now.Hour(), now.Hour())       // 9,int
	fmt.Printf("%#v,%T\n", now.Minute(), now.Minute())   // 19,int
	fmt.Printf("%#v,%T\n", now.Second(), now.Second())   // 29,int
}

// timestamp 根据当前时间取时间戳
func timestamp() {
	now := time.Now()
	fmt.Printf("%#v,%T\n", now.Unix(), now.Unix())         // 1586654796,int64 算到秒
	fmt.Printf("%#v,%T\n", now.UnixNano(), now.UnixNano()) // 1586654796699553900,int64 算到纳秒
}

// timestampToTime 时间戳int64类型的一串数转 time.Time
func timestampToTime(timestamp int64) {
	// time.Unix 要两个参数,第一个是秒,第二个是纳秒,int64类型,第二个一般可以给个0
	t := time.Unix(timestamp, 0)
	fmt.Printf("%v,%T\n", t, t) // 2020-04-12 09:26:36 +0800 CST,time.Time
}

// timeDuration 时间间隔,增、减、判断相等、前后
func timeDuration() {
	// time.Duration 类型
	timeObj := time.Unix(int64(1586654796), 0)
	now := time.Now()
	fmt.Printf("%v,%T\n", time.Minute, time.Minute) // 1m0s,time.Duration

	// 增Add,参数 time.Duration 类型,返回 time.Time 类型数据
	t1 := now.Add(time.Second * 10)
	// 可以看到结果有 +... ,但是由于程序运行需要时间,因此实际上是和输入的时间间隔存在误差
	fmt.Printf("%v,%T\n", t1, t1) // 2020-04-12 10:19:42.4102634 +0800 CST m=+10.003988501,time.Time

	// 减Sub,参数 time.Time 类型,返回值 time.Duration 类型
	t2 := now.Sub(timeObj)
	fmt.Printf("%v,%T\n", t2, t2) // 57m37.9455986s,time.Duration

	// 判断相等Equal,参数也是 time.Time 返回 bool 值
	fmt.Printf("%v\n", now.Equal(timeObj)) // false

	// 判断先后,Before 和 After ,返回 bool 值
	fmt.Println(now.Before(timeObj)) // false
	fmt.Println(now.After(timeObj))  // true
}

// setInterval 定时器
func setInterval() {
	timer := time.Tick(time.Second * 2)
	fmt.Printf("%v,%T\n", timer, timer) // 0xc0000180c0,<-chan time.Time
	for i := range timer {
		fmt.Println("已经经过了", i, "秒")
	}
	// 已经经过了 2020-04-12 10:30:53.0110678 +0800 CST m=+2.004394501 秒
	// 已经经过了 2020-04-12 10:30:55.0107491 +0800 CST m=+4.004075801 秒
	// 已经经过了 2020-04-12 10:30:57.0123484 +0800 CST m=+6.005675101 秒
	// 已经经过了 2020-04-12 10:30:59.0110164 +0800 CST m=+8.004343101 秒
	// 已经经过了 2020-04-12 10:31:01.0110985 +0800 CST m=+10.004425201 秒
}

// timeFormat 时间格式化
func timeFormat() {
	now := time.Now()

	// 2006/01/02 03:04:05.000 PM Mon Jan 记住这个就行了,Go的生日
	fmt.Println(now.Format("2006/01/02 03:04:05.000 PM Mon Jan")) // 2020/04/12 10:35:44.714 AM Sun Apr
	fmt.Println(now.Format("2006/1/2 15:4:5.000"))                // 2020/4/12 10:35:44.714
}

// timeStringParse 解析字符串格式的时间
func timeStringParse() {
	// 加载时区
	loc, err1 := time.LoadLocation("Asia/Shanghai")
	if err1 != nil {
		fmt.Println("时区加载失败", err1)
		return
	}
	fmt.Printf("%v,%T\n", loc, loc) // Asia/Shanghai,*time.Location

	// 根据上面加载的时区解析时间
	// ParseInLocation 参数: 布局, 想要解析的时间, 时区, 返回 time.Time 类型数据
	// 想要解析的时间也不能乱写,可以看看我之前测试报错信息
	// 布局里有PM Mon Jan,在时间字符串里这三个元素也不能缺,而且时间上那些数字也要和布局里的样式匹配
	// 时间字符串解析失败 parsing time "2020/4/12 22:22:22.222" as "2006/01/02 03:04:05.000 PM Mon Jan": cannot parse "4/12 22:22:22.222" as "01"
	// 时间字符串解析失败 parsing time "2020/04/12 22:22:22.222": hour out of range
	// 时间字符串解析失败 parsing time "2020/04/12 10:22:22.222 PM" as "2006/01/02 03:04:05.000 PM Mon Jan": cannot parse "" as "Mon"

	timeObj, err2 := time.ParseInLocation("2006/01/02 03:04:05.000 PM Mon Jan", "2020/04/12 10:22:22.222 PM Sun Apr", loc)
	if err2 != nil {
		fmt.Println("时间字符串解析失败", err2)
		return
	}

	fmt.Printf("%v,%T\n", timeObj, timeObj) // 2020-04-12 22:22:22.222 +0800 CST,time.Time
}

func main() {
	// days()
	// timestamp()
	// timestampToTime(1586654796)
	// timeDuration()
	// setInterval()
	// timeFormat()
	timeStringParse()
}
