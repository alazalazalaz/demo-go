
package main 
import "fmt"

/*
Printf的格式化输出，占位符如下：
获取变量类型用T
%T 		返回变量类型

布尔值
%t 		true 或false

整形
%b 		二进制表示
%d 		十进制表示
%o 		八进制表示
%x 		十六进制表示(a-f)
%X 		十六进制表示(A-F)

浮点数
%b 		无小数部分
%e 		科学计数法
%f 		有小数而无指数， eg: 123.456
%g 		根据情况选择%e 或者 %f 以阐述更紧凑的输出(无末尾的0)

字符串
%s 		字符串
%q 		双引号围绕的字符串
%x 		十六进制，每字节两个字符(a-f)

指针
%p 		十六进制表示，前缀0x

特殊格式 %v 它会自动判断类型选择默认占位符
boot 	%t
int, int8 etc...	%d
float32 etc... 		%g
string 	%s
chan 	%p
pointer %p
*/

func main(){
	var a int = 10
	var b string = "xx"
	var c float32 = 1.23456
	var p *int = &a

	fmt.Printf("a=%d \n", a)
	fmt.Printf("&a=%x \n", &a)
	fmt.Printf("p=%x \n", p)
	fmt.Printf("*p=%d \n", *p)
	fmt.Printf("*&a=%d \n", *&a)
	fmt.Printf("b=%s \n", b)
	fmt.Printf("c=%f \n", c)

}

