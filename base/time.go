package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("当前时间精确格式 ", time.Now())
	fmt.Println("当前时间戳 ", time.Now().Unix())
	fmt.Println("当前时间格式ymdhis ", time.Now().Format("2006-01-02 15:04:05")) //快速记忆,ymdHis对应 612345
	begin := time.Now().Unix()
	time.Sleep(time.Duration(2) * time.Second)
	end := time.Now().Unix()

	fmt.Println(begin, end)

	//字符串转时间戳
	s1 := "2021-10-26T06:38:13+0800"
	loc, _ := time.LoadLocation("UTC")
	theTime, err := time.ParseInLocation("2006-01-02T15:04:05+0000", s1, loc)
	if err != nil {
		fmt.Printf("err:%v", err)
	}

	fmt.Println(theTime.Unix())
}
