package main

import (
	"fmt"
	"log"
	"reflect"
	"time"
	"unsafe"
)

/**
什么是通道？
通道是go里面goroutine之间相互通信的一个管道。类似pipe。
 */
func main(){
	//what is channel
	//whatIsChannel()
	//创建
	//createNoBufferChan()
	//createBufferChan()

	//关闭通道
	//closeChan()
	//
	////读取通道
	//readChan()
	//
	////单向通道
	//singleWayChan()

	//chan做定时器和sleep的区别
	sleepOrTimerTicker()

}

/**
使用make(chan type)创建的通道是无缓冲通道，也叫同步通道
无缓冲通道特点：
1、因为没有buffer所以，发送后，对方goroutine必须立即接受，如果没有goroutine存在，则发送方会panic of deadlock
比如
make(chan int) 创建一个接受Int的通道
make(chan []int)创建一个接受[]int的通道
如果使用var c1 chan int创建，创建后由于该通道没有初始化，直接使用会报错，建议使用make创建
 */
func createNoBufferChan(){
	c1 := make(chan int)

	go func(c1 chan int) {//如果没有子协程去接受c1(也就是说如果没有这个goroutine)，那么主协程发送后会panic of deadlock
		time.Sleep(time.Second * 2)
		c1Result := <-c1
		fmt.Println("createNoBufferChan 接受c1=", c1Result)
	}(c1)

	c1<-1	//因为是no buffer，所以发送的时候会阻塞，直到接收方准备好接收。
	fmt.Println("createNoBufferChan 写入c1完成")
	time.Sleep(time.Second * 2)
	fmt.Println("createNoBufferChan end")
}

/**
创建有缓冲的通道
make(chan type, num) type表示类型，num表示缓冲大小，只要num>0就表示是有缓冲的通道
 */
func createBufferChan(){
	ch := make(chan int, 1)
	ch<-1//这里主协程直接发送，就不会报错，因为会发送到缓冲里面去
	//ch<-2//但是如果发送第二次，则会报错，因为缓冲大小只有一个
	//go func(ch chan int) {
	//	time.Sleep(time.Second*5)
	//	fmt.Println(<-ch)
	//
	//}(ch)

	//fmt.Println("createBufferChan", <-ch)
	log.Println("end")
}

func closeChan(){
	ch := make(chan int, 1)
	ch<-1
	fmt.Println("closeChan", <-ch)
	close(ch)
	fmt.Println("read a closed chan", <-ch)//继续读取一个已关闭的chan，如果里面有值会返回值，返回完后继续读取会返回该类型的空值，此处返回0
	//ch<-2 // 继续往一个已关闭的chan写入，会导致panic
	//close(ch) // 再次关闭一个已关闭的chan会panic
}

/**
读取通道中的内容，两个方法
1、i, ok := <-ch   类似map
2、for i := range ch  类似slice
 */
func readChan(){
	ch1, ch2 := make(chan int), make(chan int)
	//写入
	go func() {
		for i:=0; i<5; i++{
			ch1<-i
		}
		close(ch1)
	}()

	//第一种读取方式，类似map
	go func() {
		for {
			i, ok := <-ch1
			if ok == false {
				break
			}
			fmt.Println("readChan ch1:", i)
			ch2<- i * i
		}
		close(ch2)
	}()

	//第二种读取方式，range
	for i := range ch2 {//如果ch2没有被关闭，range会一直读，同时发现没有goroutine了(也就不会有人往ch2写了)就会panic of deadlock
		fmt.Println("readChan ch2:", i)
	}

	fmt.Println("readChan over")
}

/**
单向通道，顾名思义，让通道只能读或者只能写
 */
func singleWayChan(){
	ch1, ch2 := make(chan int), make(chan int)
	go _singleWayWrite(ch1)
	go _singleWayRead(ch2, ch1)
	for i:= range ch2{
		fmt.Println("singleWayChan ch2:", i)
	}
	fmt.Println("singleWayChan over")
}

/**
只写，注意形参的写法 chan<-
writeCh <- 写入值
chan <- type
 */
func _singleWayWrite(writeCh chan<- int){
	for i:=0; i<5; i++{
		writeCh<-i
	}
	close(writeCh)
}

/**
读&写
读出值 <- chan	//读取chan写法
type <- chan	//形参写法
 */
func _singleWayRead(writeCh chan<- int, readCh <-chan int){
	for i := range readCh{
		writeCh<-i * i
	}
	close(writeCh)
}

func whatIsChannel(){
	ch := make(chan int)
	fmt.Printf("ch size is :%d\r\n", unsafe.Sizeof(ch))
	fmt.Printf("ch type is :%v\r\n", reflect.TypeOf(ch))
}


func sleepOrTimerTicker(){
	//testSleep()
	testTicker()
	log.Println("sleepOrTimerTicker over")
}

func testSleep(){
	log.Println("testSleep begin")
	for i := 0; i < 3; i++ {
		log.Printf("testSleep for i=%d \r\n", i)
		time.Sleep(time.Second * 5)
		log.Printf("testSleep for i=%d sleep over\r\n", i)
	}
}

func testTicker(){
	t := time.NewTicker(time.Second * 1)
	i := 0
	for range t.C{
		i++
		log.Printf("testTicker for i=%d \r\n", i)
		time.Sleep(time.Second * 5)
		log.Printf("testTicker for i=%d sleep over\r\n", i)
	}
}