package main

import (
	"fmt"
	"time"
)

func main() {
	var numbers = []int{1, 2, 3, 4}
	for i, num := range numbers{
	//或者如下 _ 表示忽略掉返回的索引
	//for _, num := range numbers{ 
		fmt.Printf("index: %d, value: %d\n", i, num)
	}

	for i := range numbers{
		fmt.Printf("index: %d\n", i)//输出key
	}

	names := make(map[string]string)
	names["a"] = "abc"
	names["b"] = "def"
	for i := range names{
		fmt.Printf("index: %s\n", i)//输出key
	}

	//是否构成永动机
	_isForeverMachine()

}

/**
不会构成永动机，虽然Nums的值会被改变为如下，但是只会循环len(nums)次
[1 2 3 1]
[1 2 3 1 2]
[1 2 3 1 2 3]
 */
func _isForeverMachine(){
	nums := []int{1, 2, 3}
	for _, v := range nums{
		nums = append(nums, v)
		fmt.Println(nums)
		time.Sleep(time.Second)
	}
}

