package main

import (
	"fmt"
	"net/url"	
)

func main(){
	str := "https://cong5.net/post/golang?name=张三&age=20&sex=1&check[0]=1&check[1]=2"
	decode, _ := url.QueryUnescape(str)

	data, _ := url.Parse(decode)

	fmt.Println(str,)
	fmt.Println(decode)
	fmt.Println(data)
	fmt.Println(data.Query().Get("age"))
	fmt.Println(data.Query().Get("check[0]"))
	fmt.Printf("%T--%T\n", decode, data)
}
