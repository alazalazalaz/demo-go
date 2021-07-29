package main

import (
	"fmt"
	"math"
	"unsafe"
)

func main(){

	var i8 int8 = 1
	var i16 int16 = 1
	var i32 int32 = 1
	var i64 int64 = 1
	var i int = 1

	fmt.Printf("int8 size:%d\n", unsafe.Sizeof(i8))//1字节，能表示整数从-128~127(含) uint8=byte表示0~255(含)，byte代表了ASCII的一个字符
	fmt.Printf("int16 size:%d\n", unsafe.Sizeof(i16))//2字节，能表示整数从-32768~32767(含) uint8表示0~65535(含)
	fmt.Printf("int32 size:%d\n", unsafe.Sizeof(i32))//4字节，能表示整数-21亿~21亿多 uint32=rune表示0~42亿多,utf-8就是使用的rune类型
	fmt.Printf("int64 size:%d\n", unsafe.Sizeof(i64))//8字节，贼多
	fmt.Printf("int size:%d\r\n", unsafe.Sizeof(i))//8字节，和int64一样，贼多

	maxNum := int(math.Pow(2, 31) - 1)//或者使用math.MaxInt32
	minNum := int(math.Pow(2, 31) * -1)
	fmt.Println(maxNum, minNum)
	//大范围向小范围转换，会出现精度丢失
	var big int32 = 123456
	var small int8
	small = int8(big)
	fmt.Println(small, big)//64 123456 只会输出int8的最大值64，为啥不是127？？？

	var x int = 10
	fmt.Printf("before changeX() func : x=%d, &x=%#x \r\n", x, &x)
	changeX(x)
	fmt.Printf("after changeX() func : x=%d, &x=%#x \r\n", x, &x)
}

func changeX(x int){
	x = 20
	fmt.Printf("inner changeX() func : x=%d, &x=%#x \r\n", x, &x)
}
