package main

import (
	"demo/grpc/pb/student/pb"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct{}

func (s *server) GetInfo(ctx context.Context, req *pb.StudentInfoRequest)(*pb.StudentInfoResponse, error){
	log.Printf("客户端来访GetInfo()...")
	data := &pb.StudentInfoResponse{
		Id: req.Id,
		Name: "xiaozhang",
		Male: true,
		Scores: []int32{1, 2, 3, 4},
	}
	return data, nil
}

//测试grpc
func main(){
	//开启监听端口
	lis, err := net.Listen("tcp", "127.0.0.1:8005")
	if err != nil{
		fmt.Printf("监听端口失败：%v\n", err)
		return
	}

	s := grpc.NewServer()
	pb.RegisterStudentServer(s, &server{})
	if err :=s.Serve(lis); err != nil{
		fmt.Printf("grpc start error, %v\n", err)
		return
	}
	fmt.Println("OVER")
}
