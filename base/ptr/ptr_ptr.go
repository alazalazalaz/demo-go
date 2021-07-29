//指向指针变量的指针


package main

import (
	"fmt"
	"reflect"
)

func main(){
	var a int 
	var ptr *int
	var pptr **int

	a = 3000

	ptr = &a

	pptr = &ptr

	fmt.Printf("a=%d \n", a)
	fmt.Printf("*ptr=%d \n", *ptr)
	fmt.Printf("**pptr=%d \n", **pptr)

	var dataX = 10
	dataP := &dataX
	fmt.Printf("&dataX=%p", &dataX)
	fmt.Println(reflect.TypeOf(dataP))
}


