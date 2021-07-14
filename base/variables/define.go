package main

//变量的声明，声明的时候内存会自动初始化比如字符串为空，int为0等等，c则不会
func main(){
	//标准格式 var name type
	//声明
	var name,family string
	//定义
	name = "xiong"
	//var name = "xiong" 重复声明会报错
	//name := "xiong" := 属于声明+定义的操作，也相当于重复声明，会报错
	println(name, family)

	//简洁格式 name:=value
	sex:=1
	println(sex)

	height, width, shape := 100, 200, "round"
	println(height, width, shape)
}
