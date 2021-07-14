package main

import (
	"fmt"
	"time"
)

func main(){
	selectTest()
}

func selectTest() {
	// 2个管道
	output1 := make(chan string)
	output2 := make(chan string)
	// 跑2个子协程，写数据
	go _test1(output1)
	go _test2(output2)
	// 用select监控
	select {//如果注释掉_test1()和_test2()两个方法，也就是没有goroutine往outpu1和outpu2写入数据，则这个select会报panic of deadlock
	case s1 := <-output1:
		fmt.Println("s1=", s1)
	case s2 := <-output2:
		fmt.Println("s2=", s2)//输出_test2.1，然后退出主协程，所以看不到_test1的打印了
	//default:
	//	fmt.Println("default")
	//	time.Sleep(time.Second* 5)
	//	fmt.Println("default end")
	}
	fmt.Println("selectTest over")
	time.Sleep(time.Second* 3)
}
func _test1(ch chan string) {
	time.Sleep(time.Second * 2)
	ch <- "_test1.1"
	ch <- "_test1.2"
	ch <- "_test1.3"
}

func _test2(ch chan string) {
	time.Sleep(time.Second * 2)
	ch <- "_test2.1"
	ch <- "_test2.2"
	ch <- "_test2.3"
}
