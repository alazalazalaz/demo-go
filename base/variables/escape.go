package main

import "fmt"

type Data struct{
	Age int
}

func dummy() *Data{
	ss := make([]int, 2)
	fmt.Println(ss)
	var b Data
	d := b
	fmt.Println(d)
	c := new(Data)
	return c
}

func main(){
	fmt.Println(dummy())
}


