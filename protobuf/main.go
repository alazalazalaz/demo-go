package main

import (
	"demo/protobuf/pb/test"
	"fmt"
	"google.golang.org/protobuf/proto"
	"io"
	"log"
	"net/http"
)

/**
测试 protobuf 的使用
*/
func main(){

	http.HandleFunc("/pro", proHandler)
	err := http.ListenAndServe("127.0.0.1:8001", nil)
	if err != nil {
		fmt.Println("监听端口失败")
	}
	fmt.Println("over")
}


func proHandler(w http.ResponseWriter, r *http.Request){
	//生成student的序列化数据
	stu := &test.Student{
		Name: "xiaozhang",
		Male: true,
		Scores: nil,
	}
	s := stu.String()
	io.WriteString(w, s)

	_, err := proto.Marshal(stu)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}

	io.WriteString(w, "over")

}


