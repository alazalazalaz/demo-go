package main

import (
	"fmt"
	"log"
	"time"
)

func main() {

	//取13位时间戳
	//fmt.Println(GetTimestamp())

	fmt.Println("当前时间精确格式 ", time.Now())
	fmt.Println("当前时间戳 ", time.Now().Unix())
	fmt.Println("当前时间格式ymdhis ", time.Now().Format("2006-01-02 15:04:05")) //快速记忆,ymdHis对应 612345
	fmt.Println("RFC3339:", time.Now().Format(time.RFC3339))
	fmt.Println("RFC3339Nano:", time.Now().Format(time.RFC3339Nano))
	fmt.Println("RFC3339Nano:", time.Now().Format(time.RFC3339Nano))
	fmt.Println("RFC3339Nano:", time.Now().Format(time.RFC3339Nano))
	begin := time.Now().Unix()
	time.Sleep(time.Duration(1) * time.Second)
	end := time.Now().Unix()

	fmt.Println(begin, end)

	//字符串转时间戳
	s1 := "2022-06-09 07:49:13"
	loc, err := time.LoadLocation("UTC")
	if err != nil {
		fmt.Println(err)
		return
	}
	theTime, err := time.ParseInLocation("2006-01-02 15:04:05", s1, loc)
	if err != nil {
		fmt.Printf("err:%v", err)
		return
	}

	fmt.Println(theTime.Unix())

	//时间戳转字符串
	fmt.Println(time.Unix(1668737916, 0).Format(time.RFC3339))
	fmt.Println(time.Unix(1668737916, 0).Format("2006-01-02 15:04:05"))
	fmt.Println(time.Unix(1668737916, 0).Local())

	//计算时间差
	calculate()

	//字符串转时间戳
	stringToTimestamp()
}

// get 13 timestamp
func GetTimestamp() int64 {
	return time.Now().UnixMilli()
	//return time.Now().UnixNano() / 1e6
}

func calculate() {
	t := time.Now()
	fmt.Println(t.String()) //2023-07-07 10:29:36.490158 +0800 CST m=+1.005395210
	time.Sleep(time.Millisecond * 10)
	interval := time.Since(t)
	fmt.Println(interval.String())       //10.7666ms
	fmt.Println(interval.Milliseconds()) //10
}

// 字符串转时间戳
func stringToTimestamp() {
	s1 := "Nov 19 17:16:03 2023 +0800"
	//s1 := "Nov 9 17:16:03 2023"
	//s1 := "07:49:13 2022-Jun-9"
	loc, err := time.LoadLocation("UTC")
	if err != nil {
		fmt.Println(err)
		return
	}
	//time.RFC3339
	//http.TimeFormat
	//theTime, err := time.ParseInLocation("15:04:05 2006-Jan-2", s1, loc)
	theTime, err := time.ParseInLocation("Jan _2 15:04:05 2006 -0700", s1, loc)
	if err != nil {
		fmt.Printf("err:%v", err)
		return
	}
	log.Printf("itme:%v", theTime.Unix())
	//itme:1654760953
}

func testTicker() {
	ticker := time.NewTicker(time.Second * 2)
	for {
		select {
		case <-ticker.C:
			fmt.Println("定时器")
		}
	}
}
