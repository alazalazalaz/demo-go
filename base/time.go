package main

import (
	"fmt"
	"time"
)

func main(){

	fmt.Println("当前时间精确格式 ", time.Now())
	fmt.Println("当前时间戳 ", time.Now().Unix())
	fmt.Println("当前时间格式ymdhis ", time.Now().Format("2006-01-02 15:04:05"))//快速记忆,ymdHis对应 612345
	begin := time.Now().Unix()
	time.Sleep(time.Duration(5)*time.Second)
	end := time.Now().Unix()

	fmt.Println(begin, end)
}
