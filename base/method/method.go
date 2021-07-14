package main

import "fmt"

//函数

//定义方式
//1、常规定义
//2、函数作为一个参数来传递，实现回调。
type sayAge func(age int) int

func sayAge1(age int) int{
	fmt.Println(age)
	return age
}
func testSay(x int, sa sayAge){
	sa(x)
}

func main(){
	testSay(120, sayAge1)

	//3、定义成一个变量
	hello := func(name string){
		fmt.Println(name)
	}
	hello("xiaozhang")
}






