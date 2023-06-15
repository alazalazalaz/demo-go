// 指针是一个概念，指通过内存地址来访问变量
// 声明指针(变量)，和声明普通变量一致，只不过多一个*号
// var var_name *var_type
// * 的作用是指定该变量作为一个指针，
// var_name 存放的内容是一个内存地址
// 获取该内存地址所表示的值通过 *var_name来获取

package main

import "fmt"

type user struct {
	Name string
	Age  int
}

func main() {

	var x int = 10 //声明实际变量
	var ip *int    //声明指针变量
	ip = &x        //把实际变量的内存地址赋给指针变量

	//a变量内存地址是(&x)
	fmt.Printf("before changeIp() func, x=%d, &x=%x \r\n", x, &x)
	//变量的值是(值是a的内存地址)
	fmt.Printf("ip=%#x, *ip=%d\n", ip, *ip)
	fmt.Printf("*&x=%d\n", *&x)

	//changeIp(ip)
	//fmt.Printf("after changeIp() func, x=%d, &x=%#x \r\n", x, &x)

	fmt.Printf("before copyIp() func, ip=%#x, &ip=%#x \r\n", ip, &ip)
	copyIp(ip)
	fmt.Printf("after copyIp() func, ip=%#x, &ip=%#x \r\n", ip, &ip)

	emptyPtr()

	//浅拷贝
	u1 := &user{
		Name: "浅拷贝",
		Age:  1,
	}
	u1Result := sallowCopyPtr(u1)
	fmt.Println(u1)       // &{sallowCopyPtr 1}
	fmt.Println(u1Result) // &{sallowCopyPtr 1}

	//深拷贝
	u2 := &user{
		Name: "深拷贝",
		Age:  1,
	}
	u2Result := deepCopyPtr(u2)
	fmt.Println(u2)       //&{深拷贝 1}
	fmt.Println(u2Result) //&{deepCopyPtr 1}
}

func sallowCopyPtr(u1 *user) *user {
	result := &user{}
	result = u1
	result.Name = "sallowCopyPtr"
	return result
}

func deepCopyPtr(u2 *user) *user {
	result := &user{}
	*result = *u2
	result.Name = "deepCopyPtr"
	return result
}

func changeIp(ip *int) {
	*ip = 20
	fmt.Printf("inner changeIp() func, ip=%#x, *ip=%d\r\n", ip, *ip)
}

func copyIp(ip *int) {
	fmt.Printf("inner copyIp() func, ip=%#x, &ip=%#x \r\n", ip, &ip)
	ip = nil
	fmt.Printf("inner copyIp() func ip=>nil, ip=%#x, &ip=%#x \r\n", ip, &ip)
}

func emptyPtr() {
	var ptr *int

	fmt.Printf("空指针ptr=%x\n", ptr) //输出：0
	// fmt.Printf("空指针*ptr=%x\n", *ptr)//会报错，因为地址是0，找不到值

	if ptr == nil {
		fmt.Printf("ptr == nil为true\n")
	} else {
		fmt.Printf("ptr == nil为false\n")
	}
}
