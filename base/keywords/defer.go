package main

import (
	"fmt"
	"time"
)

/**
defer会在该函数return之前执行。
 */

func main(){
	//testParam()

	fmt.Println(testReturn())
}

/**
go的参数是使用值传递，defer后面的函数参数也一样。
比如下面这个例子，
第一个会失败，失败原因是defer后面紧接的是一个值拷贝，也就是说time.sine在调用时就计算了
第二个会成功，因为传递的是匿名函数，匿名函数会拷贝指针传递给defer，所以defer执行时是执行的最新的。
 */
func testParam(){
	beforeTime := time.Now()
	defer fmt.Println("第一个defer:", time.Since(beforeTime)) //输出0s或者xxxns约等于0s

	defer func() {
		fmt.Println("第二个defer:", time.Since(beforeTime)) //输出1s
	}()

	time.Sleep(time.Second)
}

func testReturn() (result int){
	result = 1
	defer func() {
		result++
	}()

	return result
}