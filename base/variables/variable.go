package main 
import "fmt"

// var identifier type
//函数体外定义的为全局变量
//函数体外必须用var定义

//第一种定义方式
//指定类型，如果不给初始值，系统会默认
// var v_name v_type = v_value
// var声明的变量系统都会自动初始化
var g0 int //声明，未初始化，int默认为0
var g1, g2 int = 100, 200
var gs1 string = "gs1"
var gb1 bool

// int 默认 0 
// float64 默认 0
// bool 默认 false
// string 默认 ""
// 指针 默认 nil

//第二种定义方式
//不给类型，直接赋值，系统默认给类型
// var v_name = v_value
var g10 = 10
var g20 = "xxx"


//第三种定义方式
//省略var用:=，直接给值，
//特点：
//1、只能用于函数体内
//2、必须手动初始化
//3、系统根据初始化自动判断类型
// v_name := v_value
func main(){

	i := 1 
	j, k := 10, 20
	i, i2 := 10, 20
	// i, j := 100, 100//左侧没有新的变量会报错
	fmt.Println(i, j, k, i2)//输出10 10 20 20

	var x int//只声明，不赋值，会默认0， c语言中会默认一个随机数，所以c中必须给一个初始值
	var y = 10
	fmt.Printf("x=%d, y=%d, x+y=%d\n", x, y, x+y)

}