package main

import (
	"errors"
	"fmt"
	"log"
	"time"
)

//panic 能够改变程序的控制流，调用 panic 后会立刻停止执行当前函数的剩余代码，并在当前 Goroutine 中递归执行调用方的 defer；
//recover 可以中止 panic 造成的程序崩溃。它是一个只能在 defer 中发挥作用的函数，在其他作用域中调用不会发挥作用；
func main(){
	//testP1()
	fmt.Println("继续...")

	// 测试子协程panic后，是否会影响主进程
	//testGo()

	// 子协程是一个for任务，panic后让它还能继续for continue执行
	testGoForPanic()
}

var testGoForPanicNum int = 0

func testP1(){
	fmt.Println("im testP1")
	defer func() {
		if err := recover(); err != nil{
			fmt.Println(err)
		}
		if err := recover(); err != nil{//这里是抓不到的
			fmt.Println("行行行", err)
		}
	}()

	err := errors.New("testP1挂了哦")
	panic(err)
	fmt.Println("这个不会执行")
}


func testGo(){
	// 主进程抓不到子进程的panic
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("main get panic")
		}
	}()

	go func() {
		// 当前协程能抓到当前的panic,如果注释掉recover，panic会导致主进程也挂掉
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("sub get panic")
			}
		}()
		time.Sleep(time.Second * 4)
		panic("sub goroutine panic")
	}()

	for {
		log.Println("main goroutine")
		time.Sleep(time.Second * 1)
	}

	log.Println("main")
}

func testGoForPanic(){
	go func() {
		for  {
			log.Println("sub goroutine print A")
			funcGoA()
			log.Println("sub goroutine print B")
			time.Sleep(time.Second * 1)
		}
	}()

	for {
		log.Println("main goroutine")
		time.Sleep(time.Second * 1)
	}

	log.Println("main")
}

func funcGoA(){
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("funcGoA get panic")
		}
	}()

	//do something
	time.Sleep(time.Second * 4)
	panic("funcGoA goroutine panic")
}
