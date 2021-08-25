package main

import (
	"fmt"
	"log"
	"time"
)

/**
defer会在该函数return之前执行。
 */

func main(){
	//测试未注册到的defer会不会执行
	fmt.Println(testUnRegisterDefer(1))

	//从defer中返回
	testDeferReturn()

	//testParam()
	fmt.Println(testReturn())

	//测试多个defer
	//结论，栈的方式执行defer，如果defer里面有recover，当遇到panic时，只有最顶部的defer才能recover到。
	multiDefer()
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

//输入1时，返回:
//第二个defer num=1
//第一个defer num=1
//1

//输入20时，返回：
//第一个defer num=400
//20

//说明未注册到的defer不会被执行
func testUnRegisterDefer(num int)int{
	defer func() {
		num = num * num
		fmt.Printf("第一个defer num=%d\n", num)
	}()
	if num > 10 {
		return num
	}

	defer func() {
		fmt.Printf("第二个defer num=%d\n", num)
	}()

	return num
}

func testDeferReturn(){
	name := getName(false)
	log.Printf("name is :%s", name)

}

func getName(right bool)(name string){
	log.Println("getName() begin")
	defer func() {
		log.Println("im getName() defer")
		if err := recover(); err != nil {
			name = fmt.Sprintf("%s", err)
		}

	}()

	if right{
		return "allen"
	}

	panic("panic no name")
	return "unknown"
}

func multiDefer(){
	defer func() {
		if err := recover(); err != nil{
			log.Println("multiDefer recover()!")
		}
		log.Println("defer注册1")
	}()
	log.Println("func multiDefer()")
	multiDefer2()
}

func multiDefer2(){
	defer func() {
		if err := recover(); err != nil{
			log.Println("multiDefer2 recover()!")
		}
		log.Println("defer注册2")
	}()
	log.Println("func multiDefer2()")

	multiDeferPanic()
}

func multiDeferPanic(){
	panic("test panic")
}