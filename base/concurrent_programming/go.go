package main 

import(
	"fmt"
	"log"
	"runtime"
	"time"
)

var ch chan int = make(chan int)

func main(){
	
	// func1()

	//func2()

	// funcBufferedChan()
	// funcBufferedChan2()

	//子进程修改父进程数据
	//childModifyParentData()

	//打印协程数量
	printGoNum()

}


//细品这个函数
func funcBufferedChan2(){
	chanBuff := make(chan int, 3)

	for i := 0; i < 3; i++ {//开启五个线程，写入chan，但是chan长度为3，依旧不报错
		// go func(){
			// time.Sleep(time.Duration(2)*time.Second)//如果没有sleep主线程没及时消耗就会报错
			chanBuff<-i
		// }()
	}
time.Sleep(time.Duration(2)*time.Second)
	parts := make([]int, 0)
	for {//是因为主线程在不短消耗，并且子线程有sleep，
		part, ok := <-chanBuff
		if !ok {
			break
		}

		parts = append(parts, part)
		if len(parts) == 3 {
			close(chanBuff)
		}

		
		fmt.Println(part, ok)
	}
}

func funcBufferedChan(){
	chanBuff := make(chan int, 3)
	chanBuff<-1
	chanBuff<-2
	chanBuff<-3

	fmt.Println(<-chanBuff)
	fmt.Println(<-chanBuff)
	fmt.Println(<-chanBuff)

}

//这个要报错
func func2(){
	ch <-1//执行完后，主进程阻塞，等待其他进程来消耗这个ch,但是并没有其他进程，所以死锁
	// time.Sleep(time.Duration(2)*time.Second)
	data := <-ch
	fmt.Println(data)
}

//子进程写，主进程读(阻塞读)，
func func1(){
	go loop("goroutine")
	data := <- ch
	fmt.Println(data)

	go loop("goroutine2")
	data = <- ch
	fmt.Println(data)
	// loop("main")
	fmt.Println("over")
}

func loop(str string){
	for i := 100; i < 105; i++ {
		fmt.Println(i, str)
	}

	ch<- 0
	
}

func childModifyParentData(){
	name := "father"
	go func() {
		name = "child"
	}()
	time.Sleep(time.Second * 2)
	log.Println(name)
	log.Println("over")
}

func printGoNum(){
	goroutineNum := runtime.NumGoroutine()
	log.Printf("goroutineNum:%d", goroutineNum)// 1
	go func() {
		log.Println("im child")
		time.Sleep(time.Second * 2)
	}()
	time.Sleep(time.Second * 1)
	goroutineNum2 := runtime.NumGoroutine()
	log.Printf("goroutineNum2:%d", goroutineNum2)// 2

	time.Sleep(time.Second * 3)
	goroutineNum3 := runtime.NumGoroutine()
	log.Printf("goroutineNum3:%d", goroutineNum3)// 1

}