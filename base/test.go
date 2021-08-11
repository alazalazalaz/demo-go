package main

import (
	"fmt"
	"strconv"
	"time"
)

func main(){
	fmt.Println("aaa")
	s := "3"
	curlIntervalSeconds, _ := strconv.ParseInt(s, 10, 64)
	time.Sleep(time.Duration(curlIntervalSeconds) * time.Second)
	fmt.Println("bbb")
}