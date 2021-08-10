package main

import (
	"bytes"
	"fmt"
	"github.com/yuin/goldmark"
	_ "github.com/yuin/goldmark"
	"github.com/yuin/goldmark/renderer/html"
	"io/ioutil"
	"os"
)

//what is goldmark
//goldmark是一个markdown的解析器，把markdown解析为html。
func main(){
	currentPath, _ := os.Getwd()
	fmt.Printf("current path : %s\r\n", currentPath)
	var buf bytes.Buffer
	f, err := os.Open(currentPath + "/data.md")
	if err != nil {
		fmt.Printf("err: %s", err.Error())
		os.Exit(0)
	}
	defer f.Close()
	source, err2 := ioutil.ReadAll(f)
	if err2 != nil{
		fmt.Printf("err2: %s", err2.Error())
		os.Exit(0)
	}

	md := goldmark.New(goldmark.WithRendererOptions(
		html.WithXHTML(),
		html.WithUnsafe(),
		))
	fmt.Println("raw content: ", string(source))
	if err := md.Convert(source, &buf); err != nil {
		fmt.Printf("err: %s", err.Error())
		os.Exit(0)
	}
	fmt.Printf("=============================================================================\r\nconvert content: %s", buf.String())
}