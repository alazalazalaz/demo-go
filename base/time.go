package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("当前时间精确格式 ", time.Now())
	fmt.Println("当前时间戳 ", time.Now().Unix())
	fmt.Println("当前时间格式ymdhis ", time.Now().Format("2006-01-02 15:04:05")) //快速记忆,ymdHis对应 612345
	fmt.Println("RFC3339:", time.Now().Format(time.RFC3339))
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
}
