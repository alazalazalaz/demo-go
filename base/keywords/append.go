package main

import "fmt"

/*
append()函数详解。
 */
func main(){

	var num []int
	oldNum := []int{7, 7, 7}
	num = append(num, 1)
	num = append(num, 1, 2, 3)
	num = append(num, []int{4, 5, 6}...)
	num = append(num, oldNum...)
	fmt.Println(num)
}