package main

import (
	"fmt"
	"strconv"
)

func main(){
	fmt.Println(strconv.ParseBool("a"))
	fmt.Println(strconv.ParseBool("1"))
	fmt.Println(strconv.ParseBool("false"))
}

// bool 转整形
func btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}

// 整形转bool
func itob(i int) bool {
	return i != 0
}
