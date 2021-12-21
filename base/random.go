package main

import (
	"fmt"
	"math/rand"
	"time"
)


func main(){

	// fmt.Println(Random(0, 0))//报错
	fmt.Println(Random(0, 1))//结果只有0
	fmt.Println(Random(0, 2))//结果为0或者1

	//随机字符串
	randStr()
}

//@return 左闭右开，结果会包含min，不会有max
func Random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}

func randStr(){
	n := 10
	keyLetter := []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	keyBytes := make([]byte, n)
	for i := range keyBytes {
		keyBytes[i] = keyLetter[rand.Intn(len(keyLetter))]
	}
	newClientKey := string(keyBytes)

	fmt.Println(newClientKey)
}