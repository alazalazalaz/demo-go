package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"runtime"
	"time"
)

func main() {
	printGroutine()
	ctx, _ := context.WithTimeout(context.Background(), 2 * time.Second)
	done := make(chan error, 1)//必须设置为1，如果不设置会导致goroutine泄露，详见：https://segmentfault.com/a/1190000039731121
	go func() {
		defer func() {
			if err := recover(); err != nil{
				var buf [4096]byte
				n := runtime.Stack(buf[:], false)
				log.Printf("捕获到panic：%v\n", err)
				log.Printf("panic stack:\n %s", string(buf[:n]))
				done <- errors.New(fmt.Sprintf("%s", err))
			}
		}()
		requestErr := HandelRequest(ctx)
		done <- requestErr
	}()
	printGroutine()

	select {
	case err := <-done:
		log.Println("chan over")
		log.Println(err)
	case <-ctx.Done():
		log.Println("ctx over")
		log.Println(ctx.Err())
	}

	log.Println("over")
	time.Sleep(10 * time.Second)
	log.Println("over")
	printGroutine()
}

func HandelRequest(ctx context.Context) error{
	log.Println("【go】doing request...")
	log.Println("【go】ready to panic...")
	panic("手动panic")
	time.Sleep(time.Second * 5)
	log.Println("【go】request done ...")
	return nil
}

func printGroutine(){
	log.Printf("【runtime】current goroutine num is %d\n", runtime.NumGoroutine())
}


