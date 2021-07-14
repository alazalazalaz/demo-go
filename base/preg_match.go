package main

import (
	"fmt"
	"regexp"
)

func main(){
	phone := "12211112222"
	preg := `^1[3-9][0-9]{9}$`
	reg := regexp.MustCompile(preg)
	if reg == nil {
		fmt.Println("mustcompile error")
		return
	}

	if reg.MatchString(phone){
		fmt.Println("ok")
	}else{
		fmt.Println("error format")
	}

	mail := "22@qq.jp"
	pregMail := `^[\s\S]+@[\s\S]{2,}\.[\s\S]{2,}$`
	regErr := regexp.MustCompile(pregMail)
	if regErr == nil {
		fmt.Println("mustcompile error")
		return
	}

	if regErr.MatchString(mail){
		fmt.Println("ok")
	}else{
		fmt.Println("error format")
	}


}
