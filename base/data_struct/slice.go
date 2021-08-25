package main

import (
	"fmt"
	"unsafe"
)

/*
slice切片是对数组的抽象
Go 数组的长度不可改变，所以有局限性，就出现了一个长度可以分配的"动态数组"，也就是切片
*/
func main(){
	/*
	1.定义
	*/
	//方法1，make函数(类型，长度，容量)
	// make([]type, length, capacity)
	var numbers = make([]int, 3, 5)

	//方法2，通过数组赋值
	var array = [4]int {1, 2, 3, 4}
	var emptyArray []int //已声明，但未初始化定义

	/*
	2.切片截取
	 */
	//从startIndex到endIndex-1取值(左闭右开)
	// s := arr[startIndex:endIndex]
	s1 := array[:] //所有
	s2 := array[:2]//下标0开始到1，不含下标2
	s3 := array[2:]//下标2到最后

	fmt.Printf("emptyArray len=%d, cap=%d\n", len(emptyArray), cap(emptyArray))
	fmt.Printf("number len=%d, cap=%d\n", len(numbers), cap(numbers))
	fmt.Printf("array len=%d, cap=%d\n", len(array), cap(array))
	fmt.Printf("s[:] = %d\n", s1)
	fmt.Printf("s[:2] = %v\n", s2)
	fmt.Printf("s[2:] = %v\n", s3)

	/*
	3.切片判空
	 */
	//numbers == nil

	/*
	4.切片拷贝&追加
	 */
	//append()和copy()函数
	var slice [] int
	logg(slice)

	//追加切片
	slice = append(slice, 0)
	logg(slice)
	slice = append(slice, 5, 100, 7)
	logg(slice)

	//创建slice1为slice的两倍容量
	slice1 := make([]int, len(slice), cap(slice) * 2)
	logg(slice1)
	//拷贝slice的内容到slice1，和linux的cp命令相反
	copy(slice1, slice)
	logg(slice)
	logg(slice1)

	//类似array_merge方法
	var array1 = []int {1, 2}
	var array2 = []int {3, 4}

	arrayMerge := append(array1, array2...)
	logg(arrayMerge)

	/*
	5.切片删除
	 */
	a := []int{1, 2, 3}
	//删除第一个元素
	a = a[1:]
	//删除最后一个元素
	a = a[:len(a) - 1]

	//切片是引用类型
	var abc = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	ss1 := abc[:4]
	ss2 := abc[3:7]
	fmt.Println(ss1)//[1 2 3 4]
	fmt.Println(ss2)//[4 5 6 7]
	ss1[3] = 100
	fmt.Println(ss1)//[1 2 3 100]
	fmt.Println(ss2)//[100 5 6 7]
	fmt.Println(abc)//[1 2 3 100 5 6 7 8 9]

	/*
	6. 切片大小
	*/
	var arr = [5]int{1, 2}
	fmt.Printf("intArray size:%d\r\n", unsafe.Sizeof(arr))//5*8=40字节
	sliX := []int{1, 2, 3, 4}
	fmt.Printf("sliX size:%d\r\n", unsafe.Sizeof(sliX))//始终是24

	/*
	7、数组和切片在函数内外修改的影响
	*/
	var arrX = [3]int{5, 6}
	fmt.Printf("before arrX=%v, &arrX=%p \r\n", arrX, &arrX)
	changeArrX(arrX)
	fmt.Printf("after arrX=%v, &arrX=%p \r\n\r\n", arrX, &arrX)

	var sliceX = []int{5, 6}
	fmt.Printf("before sliceX=%v, &sliceX=%p \r\n", sliceX, &sliceX)
	changeSliceX(sliceX)
	fmt.Printf("after sliceX=%v, &sliceX=%p \r\n\r\n", sliceX, &sliceX)

	/**
	8、append是否会影响slice，结论是不会哦
	 */
	var sliceY = []int{100,200}
	fmt.Printf("before changeSliceByAppend() siliceY=%v\n", sliceY)
	sliceYY := changeSliceByAppend(sliceY)
	fmt.Printf("after changeSliceByAppend() siliceY=%v, sliceYY=%v\n", sliceY, sliceYY)

}

func changeArrX(arrX [3]int){
	arrX[0] = 100
	arrX[1] = 200
	fmt.Printf("inner arrX=%v, &arrX=%p \r\n", arrX, &arrX)
}

func changeSliceX(sliceX []int){
	sliceX[0] = 100
	sliceX[1] = 200
	fmt.Printf("inner sliceX=%v, &sliceX=%p \r\n", sliceX, &sliceX)
}


func logg(slice []int){
	fmt.Printf("logg: %d \n", slice)
}

func changeSliceByAppend(sliceY []int) []int{
	sliceY = append(sliceY, 300)
	sliceY = append(sliceY, 400)
	return sliceY
}
