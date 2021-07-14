package main

import "fmt"

/*
关键字type详解
type有两个作用，一个是类型别名，一个是类型定义，
 */
func main(){
	/*
	1.类型别名
	 */
	//语法：type variable_name variable_type
	//给int取一个别名
	type myInt int
	var age myInt
	age = 10
	fmt.Printf("type:%T, value:%d \r\n", age, age)//输出 type:main.myInt, value:10

	//给匿名结构体
	//struct{
	//	header string
	//	body string
	//}
	//定义一个别名http，这也是常用的定义结构体的方式。
	type http struct{
		header string
		body string
	}

	/*
	2.类型定义
	 */
	//myInt2和int一模一样
	type myInt2 = int
	var num myInt2
	num = 20
	fmt.Printf("type:%T, value:%d \r\n", num, num)//输出 type:int, value:20


}
