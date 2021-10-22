package main

import (
	"fmt"
	"log"
	url2 "net/url"
)

func main(){
	uri := "?name=xiaozhang&pwd=123"
	url := "http://www.baidu.com/a/b" + uri
	re, err := url2.ParseQuery(uri)
	fmt.Println(re)
	fmt.Println(err)

	fmt.Println("1、url encode")
	urlEncode := url2.QueryEscape(url)
	fmt.Println(urlEncode)

	fmt.Println("2、url decode")
	urlDecode ,err := url2.QueryUnescape(urlEncode)
	fmt.Println(urlDecode)
	fmt.Println(err)

	fmt.Println("3、urlStruct")
	urlStruct, err := url2.ParseRequestURI(url)
	fmt.Println(urlStruct.Scheme)// http
	fmt.Println(urlStruct.Hostname())// www.baidu.com
	fmt.Println(urlStruct.Host)// www.baidu.com
	fmt.Println(urlStruct.Path)// /a/b
	fmt.Println(urlStruct.RawQuery)// name=xiaozhang&pwd=123
	fmt.Println(urlStruct)
	fmt.Println(err)

	fmt.Println("4、autoImplodeQuery")
	fmt.Println(autoImplodeQuery(url, "sex", "male"))

	urlParse()
}

func urlParse(){
	url := "https://sqs.us-west-2.amazonaws.com/172169962929/pf-message-queue-gold"
	a, b, c := parseAddressTemplate(url)
	log.Println(a, b, c)
}

func parseAddressTemplate(addressTemplate string) (string, string, error) {
	url_, err := url2.Parse(addressTemplate)
	if err != nil {
		return "", "", err
	}

	return url_.Scheme + "://" + url_.Host, url_.Path[1:], nil
}

// 自动在url后面拼接参数，无需判断?或者&符号
func autoImplodeQuery(url string, key string, value string) string {
	urlStruct, err := url2.ParseRequestURI(url)
	if err != nil {
		return url
	}

	urlValue, _ := url2.ParseQuery(urlStruct.RawQuery)

	urlValue.Add(key, value)

	return fmt.Sprintf("%s://%s%s?%s", urlStruct.Scheme, urlStruct.Host, urlStruct.Path, urlValue.Encode())
}