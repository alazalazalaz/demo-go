package main

import (
	"fmt"
	"sync"
)

var x, x2 int64
var wg sync.WaitGroup//实现多个goroutine同步。
var lock sync.Mutex//写锁，当一个进程获取写锁后，其他进程获取读锁或者写锁都失败。
//var rlock sync.RWMutex//读锁，当一个进程获取读锁后，其他进程可正常获取读锁，获取写锁失败。

func main(){
	//有问题的并发代码
	concurrentQuestion()
	//数据竞争问题，修复方式1
	concurrentFix1()
}

/**
两个goroutine同时去修改x，会出现数据竞争问题。
比如goroutine1把x修改到一半，for还没循环完，x就被goroutine2获取了，此时groutine2再修改，肯定是个错误的结果
 */
func concurrentQuestion(){
	wg.Add(2)
	go _add()
	go _add()
	wg.Wait()
	fmt.Println("concurrentQuestion测试结束")
}

func _add(){
	for i:=0; i<5000; i++{
		x = x + 1
	}
	fmt.Println("concurrentQuestion x=", x)
	wg.Done()
}

/**
加锁
这样能保证最终结果正确，但是加锁会影响效率
 */
func concurrentFix1(){
	wg.Add(2)
	go _add2()
	go _add2()
	wg.Wait()
	fmt.Println("concurrentQuestion测试结束")
}

func _add2(){

	for i:=0; i<5000; i++{
		lock.Lock()
		x2 = x2 + 1
		lock.Unlock()
	}
	fmt.Println("concurrentFix1 x2=", x2)
	wg.Done()
}
