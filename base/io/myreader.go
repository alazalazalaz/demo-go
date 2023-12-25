package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"strings"
)

func main() {
	stringReader()
	bufferWrite()
}

// io.Reader接口定义了Read(p []byte) (n int, err error)方法，我们可以使用它从Reader中读取一批数据。在Reader中：
// 一次最多读取len(p)长度的数据
// 读取遭遇到error(io.EOF或者其它错误), 会返回已读取的数据的字节数和error
// 即使读取字节数< len(p),也不会更改p的大小
// 当输入流结束时，调用它可能返回 err == EOF 或者 err == nil，并且n >=0, 但是下一次调用肯定返回 n=0, err=io.EOF
//
// 打印为：
// 2023/12/22 14:46:04 reading:4, hell
// 2023/12/22 14:46:04 reading:4, o wo
// 2023/12/22 14:46:04 reading:3, rld
// finish read
func stringReader() {
	reader := strings.NewReader("hello world")
	b := make([]byte, 4)
	for {
		n, err := reader.Read(b)
		if err != nil {
			//出错
			if err == io.EOF {
				//读完
				fmt.Println("finish read")
				break
			}

			log.Printf("err:%v", err)
			return
		}

		log.Printf("reading:%d, %s", n, string(b[:n]))
	}
}

func bufferWrite() {
	providers := []string{
		"hello",
		"world",
		"golang",
		"is great",
	}
	var write bytes.Buffer
	for _, s := range providers {
		n, err := write.Write([]byte(s))
		if err != nil {
			fmt.Printf("write err:%v", err)
			return
		}

		if n != len(s) {
			fmt.Printf("failed to write data, write len:%d, target len:%d", n, len(s))
			return
		}
	}

	fmt.Printf("finish:%v\n", write.String())
	fmt.Printf("finish:%v\n", write.String())
}
