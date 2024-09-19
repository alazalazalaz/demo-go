package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, err := http.Get("http://www.baidu.com")
	defer resp.Body.Close()
	if err != nil {
		fmt.Println("请求失败:", err)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取响应体失败:", err)
		return
	}

	fmt.Println(string(body))
}
