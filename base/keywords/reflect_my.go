package main

import (
	"fmt"
	"reflect"
)

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(reflect.TypeOf(nums).Name())        //空
	fmt.Println(reflect.TypeOf(nums).Kind())        //slice
	fmt.Println(reflect.TypeOf(nums).String())      //[]int
	fmt.Println(reflect.TypeOf(nums).Elem().Kind()) //int
	fmt.Println(reflect.TypeOf(nums))               //[]int
	fmt.Println(reflect.ValueOf(nums))              //[1 2 3]

	//使用字符串调用函数
	stringCall("getMoney")
	stringCall("getJob")
}

func stringCall(name string) {
	funcMap := map[string]interface{}{
		"getMoney": getMoney,
		"getJob":   getJob,
	}
	nameV := reflect.ValueOf(funcMap[name]) //funcMap[name]这玩意儿的值是个func

	//params := make([]reflect.Value, 0)
	nameV.Call(nil)
}

func getMoney() {
	fmt.Println("im getMoney function")
}
func getJob() {
	fmt.Println("im getJob function")
}
