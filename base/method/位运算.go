package main

import "fmt"

func main(){
	x, y, z := 2, 3, 3
	fmt.Println(x^y)
	fmt.Println(y^z)
	fmt.Println(3^5)

	//左移<<运算，二进制左移n位，右边补0，相当于乘以2，但效率比乘法高
	leftCal()
}

func leftCal(){
	fmt.Println("<<<<<<<<<<")
	fmt.Println(1<<1)//二进制1左移一位，变10，转为十进制为2
	fmt.Println(1<<2)//1=>100
	fmt.Println(1<<3)
	fmt.Println(10<<1)
	fmt.Println(10<<2)
	fmt.Println(10<<23)
	fmt.Println(2<<1)
	fmt.Println(2<<2)
	fmt.Println(2<<3)
}
