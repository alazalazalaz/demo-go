package main

import (
	"encoding/base64"
	"fmt"
)

func main(){
	str := "test string"
	baseStr := base64.StdEncoding.EncodeToString([]byte(str))

	fmt.Println(baseStr)

	decodeByte, err := base64.StdEncoding.DecodeString(baseStr)
	fmt.Println(string(decodeByte))
	fmt.Println(err)
}