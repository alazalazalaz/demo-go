package main

import "fmt"

/*
golang关于三个省略号的使用
*/
func main() {
	/*
		1.使用在数组中，表示不定长，数组长度由{}内的定义时给出
	*/
	var numbers = [...]int{1, 2, 3}
	fmt.Println(numbers)

	/*
		2. ...放后面 打散切片
	*/
	arr := []int{1, 2, 3}
	var arr2 []int
	arr2 = append(arr2, arr...) //表示传入的是一个[]int类型的切片，append函数接受的时候是接受到打散的int，比如1, 2 ,3
	arr2 = append(arr2, 4, 5, 6)
	fmt.Println(arr2) // [1 2 3 4 5 6]

	/*
		3. ...放前面 变长函数参数
	*/
	getNum("test", 7, 7, 7) //输出 test [7 7 7]
	//也可以这样传入
	arr3 := []int{8, 8, 8}
	("test2", arr3...) //输出 test2 [8 8 8]

}

/*
@param num ...int是一个切片哦等同于num []int
* num ...int和num []int相同点是对于函数内部使用来说都一样，不同点是函数调用者，前者需要传递n个单独元素，后者只能传递一个切片
*/
func getNum(name string, num ...int) {
	fmt.Println(name, num, num[0])
}
