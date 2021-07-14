package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var isQuit bool

/**
context是上下文，为什么要上下文这个东西？
问题：
如何优雅的让子协程退出？
1、通过定义一个全局变量，让主协程修改全局变量，子协程循环判断全局变量判断是否该退出
2、channel方式
3、官方推荐context
 */
func main(){
	//全局变量方法
	_quitWorkerMethod1()
	//channel方法
	_quitWorkerMethodChannel()
	//context方法，如果有多个协程需要取消，则直接传入ctx则可。
	_quitWorkerMethodContext()
}

func _quitWorkerMethodContext(){
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go _quitWorkerMethodContextGo(ctx)
	time.Sleep(time.Second * 3)
	cancel()
	wg.Wait()
	fmt.Println("_quitWorkerMethodContext over")
}

func _quitWorkerMethodContextGo(ctx context.Context){
	_quitWorkerMethodContextGo2(ctx)
	LOOP:
		for {
			fmt.Println("worker1 context")
			time.Sleep(time.Second)
			select {
			case <-ctx.Done():
				break LOOP
			default:

			}
		}
		wg.Done()
}

func _quitWorkerMethodContextGo2(ctx context.Context){
LOOP:
	for {
		fmt.Println("worker2 context")
		time.Sleep(time.Second)
		select {
		case <-ctx.Done():
			break LOOP
		default:

		}
	}
	wg.Done()
}

func _quitWorkerMethodChannel(){
	ch := make(chan string)
	wg.Add(1)
	go _quitWorkerMethodChannelGo(ch)
	time.Sleep(time.Second * 3)
	ch<-"quit"
	wg.Wait()
	fmt.Println("_quitWorkerMethod2 over")
}

func _quitWorkerMethodChannelGo(ch chan string){
	LOOP:
		for {
			fmt.Println("worker channel")
			time.Sleep(time.Second)
			select {
			case <-ch:
				break LOOP
			default:

			}
		}
		wg.Done()
}


func _quitWorkerMethod1(){
	wg.Add(1)
	go _quitWorkerMethod1Go()
	time.Sleep(time.Second * 3)
	isQuit = true
	wg.Wait()
	fmt.Println("_quitWorkerMethod1 over")
}

func _quitWorkerMethod1Go(){
	for {
		fmt.Println("worker")
		time.Sleep(time.Second)
		if isQuit == true{
			break
		}
	}
	wg.Done()
}

