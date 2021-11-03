package main

import (
	"context"
	"log"
	"runtime"
	"time"
)

func main() {
	printGroutine()

	//创建子协程
	totalGoNum := 10
	for i := 0; i < totalGoNum; i++{
		go func() {
			//子协程设置超时2s
			ctx, cancel := context.WithTimeout(context.Background(), time.Second * 5)
			defer cancel()

			done := make(chan error)
			//必须设置为1，如果不设置会导致goroutine泄露，详见：https://segmentfault.com/a/1190000039731121
			//如果不设置chan的buffer，孙协程往里面写的时候，发现没有其他协程(此处为子协程)在读取，所以会死锁，该goroutine会泄露不会被释放掉。
			//创建孙协程向done写入
			go func() {
				time.Sleep(time.Second * 10)//延迟写入，比如孙协程是个耗时任务
				done<- nil
				log.Println("孙协程写入完毕")
			}()

			select {
			case <-done:
				log.Println("收到孙协程的写入")
			case <-ctx.Done():
				log.Println("子协程超时")
			}

		}()
	}

	time.Sleep(time.Second * 20)
	printGroutine()
}

func printGroutine(){
	log.Printf("【runtime】current goroutine num is %d\n", runtime.NumGoroutine())
}


