package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {

	rawCurl()
}

func rawCurl() {
	resp, err := http.Get("https://pf-chat-en2en.tap4fun.com/k1d2-beta/3048/permanent/avatar")
	if err != nil {
		fmt.Println("err")
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	fmt.Println(resp.StatusCode)
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("err:%v\n", err)
		return
	}

	fmt.Println(data)
}
