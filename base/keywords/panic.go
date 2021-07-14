package main

import (
	"errors"
	"fmt"
)

//panic 能够改变程序的控制流，调用 panic 后会立刻停止执行当前函数的剩余代码，并在当前 Goroutine 中递归执行调用方的 defer；
//recover 可以中止 panic 造成的程序崩溃。它是一个只能在 defer 中发挥作用的函数，在其他作用域中调用不会发挥作用；
func main(){
	testP1()
	fmt.Println("继续...")
}

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

