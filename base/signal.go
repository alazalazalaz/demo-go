package main 

import(
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main(){
	

	fmt.Println("启动")

	

	go func(){
		c := make(chan os.Signal)
		
		//监听
		signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

		fmt.Println("子进程阻塞等待")
		s := <-c
		fmt.Println("子进程监听到信号", s)
		exitFunc()
	}()

	go func(){
		time.Sleep(time.Duration(5)*time.Second)
		fmt.Println("哈哈哈哈哈哈")
	}()
	
	sum := 0
	for {
		sum++
		fmt.Println(sum)
		time.Sleep(time.Duration(1)*time.Second)
	}
	
	fmt.Println("父进程这里执行不到了")
}


func exitFunc(){
	fmt.Println("开始退出...")
	fmt.Println("执行清理...")
	fmt.Println("结束退出...")
	os.Exit(0)
}
