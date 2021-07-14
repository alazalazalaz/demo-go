package main

import "fmt"

func main(){
	//make函数
	_make()

	//new函数
	_new()
}

func _make(){
	s1 := make([]int, 0)
	fmt.Printf("make初始化切片为：%v\r\n", s1)

	m1 := make(map[int]string)
	fmt.Printf("make初始化map为：%v\r\n", m1)

	c1 := make(chan int)
	fmt.Printf("make初始化chan为：%v\r\n", c1)
}

func _new(){
	int1 := new(int)
	string1 := new(string)
	fmt.Printf("new初始化int为：%v，字符串为：%v\r\n", int1, string1)

	s1 := new([]int)
	fmt.Printf("new初始化切片为：%v\r\n", s1)

	m1 := new(map[int]string)
	fmt.Printf("new初始化map为：%v\r\n", m1)

}
