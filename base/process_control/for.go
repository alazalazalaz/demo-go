//和c语言的for一样
// for init; condition; post{}

//和c语言的while一样
// for condition {}

//和c语言的for(;;)一样(无限循环)
// for {}

package main

import (
	"fmt"
	"log"
	"sync/atomic"
	"time"
)

func main() {

	var sum = 0
	for i := 0; i < 10; i++{
		sum += i
	}
	fmt.Printf("%d\r\n", sum)

	//是否能构成永动机？
	//确实会哦，但是range不会
	//nums: [1 2 3 1]
	//nums: [1 2 3 1 2]
	//nums: [1 2 3 1 2 3]
	//nums: [1 2 3 1 2 3 1]
	//nums: [1 2 3 1 2 3 1 2]
	//...
	//nums := []int{1, 2, 3}
	//for i:= 0; i < len(nums); i++{
	//	nums  = append(nums, nums[i])
	//	fmt.Println("nums:", nums)
	//	time.Sleep(time.Second)
	//}

	workerCount := int32(10)
	QueueServiceWorkerNumber := 10
	for int(workerCount) >= QueueServiceWorkerNumber {
		log.Println(workerCount, QueueServiceWorkerNumber)
		time.Sleep(1 * time.Second)
	}
	atomic.AddInt32(&workerCount, int32(1))
	log.Println(workerCount)
}