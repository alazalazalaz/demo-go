package main

import (
	"fmt"
	"github.com/shopspring/decimal"
	"log"
	"math"
)

func main() {

	int64ToFloat64()
	float64ToString()

	return
	var num = 1.1
	var total = float64(0)
	for i := 0; i < 10; i++ {
		total += num
	}
	//正确结果应该是11
	println(int(total))            //输出10
	fmt.Printf("%5.5f\r\n", total) //输出11.00000

	//关于浮点数的除法，如果左侧(除数)是明确的int或者float，那右侧的数据类型则会自动变为和左侧变量类型相同。比如：
	var bigNum int = 9999
	var bigNumf float64 = 9999
	fmt.Printf("%v\r\n", bigNum/100)  //输出int 99
	fmt.Printf("%v\r\n", bigNumf/100) //输出float 99.99
	//如果左侧是具体的数值常量而不是变量，则左侧类型会自动等于右侧的类型。比如：
	fmt.Printf("%v\r\n", 110/100)   //输出int 1
	fmt.Printf("%v\r\n", 110/100.0) //输出float 1.1
	//所以浮点数的除法得的结果类型是以左侧为准，左侧是常量则已右侧为准

	println("------")
	var num1 = 19.9
	log.Println(num1 * 100) //等于1989.9999999999998
	//采用decimal包解决
	re, _ := decimal.NewFromFloat(num1).Mul(decimal.NewFromInt(100)).Float64()
	log.Println(re)
}

func int64ToFloat64() {
	var a int64 = 1
	fmt.Println(float64(a))
	fmt.Println(float64(a / 100)) //精度丢失，等于0
	fmt.Println(float64(a) / 100) //等于0.01
}

func float64ToString() {
	fmt.Println("float64ToString")
	var a float64 = 0.1234
	fmt.Printf("%0.0f\n", a) // 0, 不含小数点和小数
	fmt.Printf("%0.2f\n", a) // 0.12, 含小数点和小数

	var newPrice float64 = 0.1234
	newPrice = math.Ceil(newPrice)
	fmt.Printf("%0.0f\n", newPrice) // 0, 不含小数点和小数
	fmt.Printf("%0.2f\n", newPrice) // 0.12, 含小数点和小数

}
