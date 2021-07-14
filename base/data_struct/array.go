
// var variable_name [SIZE] variable_type

//定义一个长度10，类型为float32的数组
// var balance [10] float32
// 取值时，用下标，xxx[num]，此num是0开始哦，上面的SIZE是长度
package main 
import "fmt"

/*数组*/
func main(){
	/*
	1. 定义
	*/

	//var variable_name[SIZE] variable_type
	var n [10] int
	for i := 0; i < 10; i++{
		n[i] = i + 100
	}

	//声明
	var array [5] int
	//定义
	array = [5]int{1, 2, 3, 4, 5}
	//声明+定义
	var array2 = [...]int{1, 2, 3, 4, 5} //...自动计算长度
	//其中数组必须指定长度或者使用...自动计算后面{}中的长度
	//切片则不必，比如var s []int
	//所以区分数组和切片的一个点就是，看定义的时候，是否指定长度。
	//切片定义
	//s1 := array[:2]
	//s2 := []int{1,2}
	//s3 := make([]int, 2, 3)


	fmt.Println(array)
	fmt.Println(array2)

	/*
	2.多维数组
	*/
	var multArray[2][3] int
	multArray = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(multArray)

	/**
	数组是否是引用类型呢？
	 */
	//数组不是哦
	var abc = [3]int{1,2,3}
	_checkType(abc)
	fmt.Println("数组-是否是引用类型呢变换后函数外abc:", abc)
	//切片是哦
	abcS := []int{1, 2, 3}
	_checkTypeS(abcS)
	fmt.Println("切片-是否是引用类型呢变换后函数外abc:", abcS)
}

func _checkType(abc [3]int){
	fmt.Println("数组-是否是引用类型呢abc:", abc)
	abc[1] = 1
	abc[2] = 1
	fmt.Println("数组-是否是引用类型呢变换后abc:", abc)
}

func _checkTypeS(abc []int){
	fmt.Println("切片-是否是引用类型呢abc:", abc)
	abc[1] = 1
	abc[2] = 1
	fmt.Println("切片-是否是引用类型呢变换后abc:", abc)
}

