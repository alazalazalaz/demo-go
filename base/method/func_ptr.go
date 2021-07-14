//c语言或其他语言中，参数有两种传递
//1、值传递，不会影响实际参数的值
//2、引用传递时(用指针)，如果函数内的形参改变了，则函数外的实参也会变
//go语言中，实际只有一种传递
//那就是指传递，
//1、当传递值时，和其他语言一样
//2、当传递指针时，函数内部会copy一个指针，这个copy的指针指向的当然还是实参的值。所以函数内部改变这个指针的value后，
//实参还是会变，但是如果函数内部删除这个指针，是对外界没有任何影响的。
//举例，可以把函数内外的指针地址(指针变量的地址哈，不是指针内容，指针内容是指向实参的地址，肯定相同的)打印出来，地址是不同的，说明是copy了一个指针，而c语言两个指针的地址都是相同的。
package main 
import "fmt"

func main(){
	var stringX = "xxx"
	changeString(stringX)
	fmt.Println(stringX)//xxx 普通变量传递，不会被修改

	var stringPtr = "ptrxxx"
	fmt.Println("传入之前ptrxxx地址为：", &stringPtr)
	changeStringPtr(&stringPtr)
	fmt.Println(stringPtr)//modify 指针变量传递，会被修改


	var array1 = [3]int{1, 2, 3}
	changeArray(array1)
	fmt.Println(array1)//[1 2 3] 普通数组传递，不会被修改

	var array2 = [3]int{1, 2, 3}
	changeArrayPtr(&array2)
	fmt.Println(array2)//[100 200 3] 指针数组传递，会被修改


	var sss1 = []int{1, 2, 3, 4}
	changeSlice(sss1)
	fmt.Println(sss1)//[100 200 3 4]切片是引用传递，会被修改值哦，但是使用append函数不会，只有修改下标才会被重新赋值

	var sss2 = []int{1, 2, 3, 4}
	changeSlicePtr(&sss2)
	fmt.Println(sss2)//[100 200 3 4]指针切片传递，会被修改


	var mapX = map[string]string{"country":"中国"}
	changeMap(mapX)
	fmt.Println(mapX)//map[country:中国updated] 说明map也是引用传递

	var mapXPtr = map[string]string{"country":"中国"}
	changeMapPtr(&mapXPtr)
	fmt.Println(mapXPtr)//map[country:中国updated]

	//结构体
	cat := animal{}
	cat.age = 2
	cat.name = "cat"
	changeStruct(cat)
	fmt.Println(cat)	//{2 cat} 普通结构体传递，不会被修改

	dog := &animal{
		age: 1,
		name: "dog",
	}
	fmt.Println("传入之前：", &dog)
	changeStructPtr(dog)
	fmt.Println(dog)	//&{100 name updated} 指针结构体传递，会被修改

}


func changeArray(array1 [3]int){
	array1[0] = 100
	array1[1] = 200
}

func changeArrayPtr(array2 *[3]int){
	(*array2)[0] = 100
	(*array2)[1] = 200
}

func changeSlice(s1 []int){
	s1[0] = 100
	s1[1] = 200
}

func changeSlicePtr(s2 *[]int){
	(*s2)[0] = 100
	(*s2)[1] = 200
}

func changeString(s string){
	s = "modify"
}

func changeStringPtr(sPtr *string){
	fmt.Println("传入之后ptrxxx地址为：", &sPtr)
	*sPtr = "modify"
}

func changeMap(m map[string]string){
	m["country"] = "中国updated"
}

func changeMapPtr(m *map[string]string){
	(*m)["country"] = "中国updated"
}

type animal struct{
	age int
	name string
}

func changeStruct(an animal){
	an.name = "name updated"
	an.age = 100
}

func changeStructPtr(an *animal){
	fmt.Println("传入之后：", &an)
	an.name = "name updated"
	an.age = 100
	an = nil //！！！给指针赋值nil是无效的，an依旧存在！！！！
}

