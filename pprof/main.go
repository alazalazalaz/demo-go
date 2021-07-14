package main

import (
	"fmt"
	"io"
	"net/http"
	_ "net/http/pprof"
	"time"
)

/**
测试pprof的使用
 */
func main(){
	go func() {
		testPprof()
	}()

	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/test", testHandler)
	http.ListenAndServe("127.0.0.1:1111", nil)
	fmt.Println("over")
}

func testPprof(){
	tick := time.Tick(time.Second / 100)
	var buf []byte
	for range tick {
		buf = append(buf, make([]byte, 1024*1024)...)
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request){
	for {
		io.WriteString(w, "hello word")
		time.Sleep(time.Second)
	}
}

func testHandler(w http.ResponseWriter, r *http.Request){
	io.WriteString(w, "test")
}
