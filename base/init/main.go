package main

import "fmt"

func init() {
	fmt.Println("init main") //执行顺序1
}

func main() {
	fmt.Println("main") //执行顺序3
}
