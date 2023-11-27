package main

import (
	"fmt"
	"io"
	"net/http"
	_ "net/http/pprof"
	"path"
	"time"
)

/**
测试pprof的使用
*/
func main() {
	go func() {
		for i := 0; i < 10; i++ {
			testPprof()
			time.Sleep(time.Second * 10)
		}
	}()

	//http.HandleFunc("/hello", helloHandler)
	//http.HandleFunc("/test", testHandler)
	http.ListenAndServe("127.0.0.1:1111", nil)
	fmt.Println("over")
}

func testPprof() {
	var buf []byte
	buf = append(buf, make([]byte, 1024*1024)...)
	path.Join()
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	for {
		io.WriteString(w, "hello word")
		time.Sleep(time.Second)
	}
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "test")
}
