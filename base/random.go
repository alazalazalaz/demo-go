package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// 通过new source防止种子交叉覆盖
	testSeedWithNew()
	// 如果两个副本的种子相同，则生成出来的随机数也是相同的
	//testSeedWithSameResult()

	// fmt.Println(Random(0, 0))//报错
	//fmt.Println(Random(0, 1)) //结果只有0
	//fmt.Println(Random(0, 2)) //结果为0或者1
	//
	////随机字符串
	//randStr()
}

// 所以最终的生成随机数的方案为：
//1、单独建立一个source，
//2、副本启动时设置一个独立的种子(比如用启动时间纳秒级别)，这样该副本后续使用这个source生成的随机数都不会重复
//3、如果多个副本启动时间不同，那多个副本之间也不会随机数重复
func testSeedWithNew() {
	//使用新的source
	r := rand.New(rand.NewSource(123))
	a1 := r.Int()
	a2 := r.Int()
	fmt.Println(a1, a2)
	// rand库自带的source
	rand.Seed(123)
	a3 := rand.Int() // 5361704182646325489
	a4 := rand.Int() // 241876450138978746
	fmt.Println(a3, a4)

	//a1_1 := r.Int() // 2305561650894865143
	//a3_1 := rand.Int() // 2305561650894865143
	//fmt.Println(a1_1, a3_1)

	rand.Seed(1234)
	a1_1 := r.Int()    // 2305561650894865143
	a3_1 := rand.Int() // 2041104533947223744 这里变了！因为用rand库重新设置了种子
	fmt.Println(a1_1, a3_1)

}

// 下面这个代码放在任何副本执行，结果都一样！！！
// 也就是说只要种子相同，不同副本按照顺序生成的随机数肯定是相同的，哪怕source不同，生成的也是相同的。
// 但是如果是单副本，一直调用rand.Int()生成的随机数是永远也不会相同的。
func testSeedWithSameResult() {
	rand.Seed(123)
	a1 := rand.Int() // 5361704182646325489
	a2 := rand.Int() // 241876450138978746
	fmt.Println(a1, a2)

	rand.Seed(123)
	a3 := rand.Int() // 5361704182646325489
	a4 := rand.Int() // 241876450138978746
	fmt.Println(a3, a4)
}

//@return 左闭右开，结果会包含min，不会有max
func Random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}

func randStr() {
	n := 10
	keyLetter := []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	keyBytes := make([]byte, n)
	for i := range keyBytes {
		keyBytes[i] = keyLetter[rand.Intn(len(keyLetter))]
	}
	newClientKey := string(keyBytes)

	fmt.Println(newClientKey)
}
