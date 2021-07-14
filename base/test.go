package main

import (
	"fmt"
	"strings"
)

func main(){
	decodeByte := "111-222-333"
	secInfoArray := strings.Split(string(decodeByte), "-")
	for k, v := range secInfoArray{
		fmt.Println(k, v)
	}

	fmt.Println(secInfoArray[2])
	fmt.Println(secInfoArray[3])
	fmt.Println(secInfoArray[4])
}