package main

import (
	"fmt"
	"github.com/thinkeridea/go-extend/exunicode/exutf8"
	"log"
	"reflect"
	"unicode/utf8"
	"unsafe"
)

/*GO中的字符串，默认使用的是UTF-8编码*/
func main(){
	//在golang中，字符的本质是一个int8的整数，也就是byte这个类型。这个整数表示的是这个字符对应的ASCII字符。范围为-2^7~2^7-1(-256~255)。
	//我们回顾一下，26个字母表示的ASCII码范围。

	//定义一个字符
	var c byte = 'a'

	//输出一个字符，会直接输出它对应的ASCII码
	log.Println(c)//97

	//格式化输出字符，才会输出它对应的值
	log.Printf("%c", 169)//a

	//一个byte是int8，只有一个字节大小，如果用中文(3/4个字节)会溢出，编译失败
	//var z byte = '中'//constant 20013 overflows byte

	//可以用一个int32或者int64来存放单个中文汉字，然后格式化输出
	var zh int = '中'//注意是单引号哦，表示单个字符
	log.Printf("%c", zh)//中

	//再来说说字符串的定义
	//使用双引号或者反引号
	var s1,s2,s3 = "012", "allen", "严"
	log.Printf("len(s1):%d, len(s2):%d, len(s3):%d", len(s1), len(s2), len(s3))
	log.Printf("[]byte(s1):%v, []byte(s2):%v, []byte(s3):%v", []byte(s1), []byte(s2), []byte(s3))
	log.Printf("[]rune(s1):%v, []rune(s2):%v, []rune(s3):%v", []rune(s1), []rune(s2), []rune(s3))

	fmt.Println('a' - 'b')
	name := "abc一"
	by := []byte(name)
	byRune := []rune(name)
	empty := ""
	emptyBytes := []byte(empty)

	fmt.Printf("size of name is :%d \r\n", unsafe.Sizeof(name))
	fmt.Println(name, by, emptyBytes, byRune)//abc一 [97 98 99 228 184 128] [] [97 98 99 19968]
	fmt.Printf("%T---%T\n", name, by)//string---[]uint8

	/*1.字符串截取和长度*/
	begin := 1
	end := -1
	strlenEx("abcd123490", begin, end)//截取字符串，输出 bcd12349
	strlenEx("一二三四五六七八90", begin, end)//截取含中文的字符串，输出 二三四五六七八9

	fmt.Println(len("abcd123490"), len("一二三四五六七八90"))//字节长度，输出10 26

	fmt.Println(utf8.RuneCountInString("一二三四五六七八90"))//字符长度，输出10

	/*2.字符串ASCII*/
	theme := "我爱CD😄"
	for i:=0; i<len(theme); i++ {
		//下标的方式可以直接获取到字符对应的ASCII
		fmt.Printf("theme[%d]=%c %d\r\n", i, theme[i], theme[i])
	}


	for _, s := range theme{
		//使用unicode方式输出正常中文字符
		fmt.Printf("unicode: %c %d \r\n", s, s)
	}

	/*3.字符串拼接，除了用+号，还可以使用bytes.Buffer */

	/*4.printf的值
	%v	按值的本来值输出
	%+v	在 %v 基础上，对结构体字段名和值进行展开
	%#v	输出 Go 语言语法格式的值
	%T	输出 Go 语言语法格式的类型和值
	%%	输出 % 本体
	%b	整型以二进制方式显示
	%o	整型以八进制方式显示
	%d	整型以十进制方式显示
	%x	整型以十六进制方式显示
	%X	整型以十六进制、字母大写方式显示
	%U	Unicode 字符
	%f	浮点数
	%p	指针，十六进制方式显示*/
}

//截取字符串长度
func strlen(str string, begin int, end int){
	newStr := str[begin : end]
	fmt.Println(str, begin, end, newStr)
}
//此方法可兼容中文
func strlenEx(str string, begin int, end int){
	newStr := exutf8.RuneSubString(str, begin, end)
	fmt.Println(str, newStr)
}

//ASCII码中，数字0-9是48-57表示
//ASCII码中，大写字母A-Z是65-90表示
//ASCII码中，小写字母a-z是97-122表示
func isPalindrome(s string) bool {
	for k, s1 := range s{
		fmt.Println(reflect.TypeOf(s1), reflect.ValueOf(s1), k, s1, s[k])
		//int32 65 0 65 65
		//int32 97 1 97 97
		//int32 48 2 48 48
		//int32也就是rune
	}
	sByte := []byte(s)
	for _, sByte1 := range sByte{
		fmt.Println(reflect.TypeOf(sByte1), reflect.ValueOf(sByte1), sByte1)
		//uint8 65 65
		//uint8 97 97
		//uint8 48 48
		//uint8 也就是byte
	}
	return true
}

