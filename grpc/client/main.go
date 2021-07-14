package main

import (
	"context"
	"demo/grpc/pb/student/pb"
	"google.golang.org/grpc"
	"log"
)

const (
	address = "127.0.0.1:8005"
)

//grpc客户端
func main(){
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil{
		log.Fatalf("客户端连接失败:%v", err)
	}
	defer conn.Close()
	c := pb.NewStudentClient(conn)
	r, err := c.GetInfo(context.Background(), &pb.StudentInfoRequest{Id: 100})
	if err != nil {
		log.Fatalf("接口返回错误:%v", err)
	}
	log.Printf("接口返回：%v", r)
}
