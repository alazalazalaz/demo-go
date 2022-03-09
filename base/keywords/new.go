package main

import "fmt"

func main() {
	a := []int{0, 0, 0} // 提供初始化表达式。
	a[1] = 10

	b := make([]int, 3) // make slice
	b[1] = 10

	d := make([]int, 1) // make slice
	d[0] = 10
	//d[1] = 100//这样会越界
	d = append(d, 1) //这样就不会了，会自动扩容
	d = append(d, 2)
	fmt.Println(d)

	//c := new([]int) 这样会报错哟
	//c[1] = 10 // ./main.go:11:3: invalid operation: c[1] (type *[]int does not support indexing)

	data := Data{
		Id: 1,
	}

	fmt.Println(data.Info.Age)
	fmt.Println(data.InfoP.Age)

	dataP := &Data{
		Id: 2,
	}
	fmt.Println(dataP.Info.Name)
}

type Data struct {
	Id    int
	Info  Info
	InfoP *Info
}

type Info struct {
	Age  int
	Name string
}
