package main

import (
	"fmt"
	"sync"
	"time"
)

var map1 map[string]string
var wg1 sync.WaitGroup

var map2 map[string]string
var wg2 sync.WaitGroup
var mutex2 sync.Mutex
var rmutex2 sync.RWMutex

var map3 sync.Map
var map3Lock sync.Mutex

//sync.map
func main() {
	//如何保证一个全局变量map的key，只会被写入一次？
	//方法1：常规的单例，会有并发问题,concurrent map writes
	//storeMapFunc1()
	//
	////方法2：使用map+mutex
	////结论：不会有并发问题，
	//storeMapFunc2()

	//方法2：使用sync.map
	storeMapFunc3()
}

func storeMapFunc1() {
	map1 = make(map[string]string)
	key := "A"
	num := 10

	for i := 0; i < num; i++ {
		go func() {
			v := getV1(key, fmt.Sprintf("v%d", i))
			fmt.Println(v)
		}()
	}
	time.Sleep(time.Second * 2)
	fmt.Println(getV1(key, "over"))
}

func getV1(key string, value string) string {
	if v, isExist := map1[key]; isExist {
		return v
	}

	map1[key] = value //这里会报panic =》 concurrent map writes
	return map1[key]
}

func storeMapFunc2() {
	map2 = make(map[string]string)
	key := "B"
	num := 10

	for i := 0; i < num; i++ {
		go func() {
			v := getV2(key, fmt.Sprintf("v%d", i))
			fmt.Println(v)
		}()
	}
	time.Sleep(time.Second * 2)
	fmt.Println(getV2(key, "over"))
}

func getV2(key string, value string) string {
	if v, isExist := map2[key]; isExist {
		return v
	}

	//加锁，不会出现并发问题，但是依旧会被多次修改
	mutex2.Lock()
	defer mutex2.Unlock()

	//double check可以防止并发修改，也可以把锁放在函数开始就不需要double check了，也可以修改为读锁。
	if v, isExist := map2[key]; isExist {
		return v
	}

	map2[key] = value
	return map2[key]
}

func getV2ByReadMutex(key string, value string) string {
	rmutex2.RLock()
	defer rmutex2.RUnlock()

	if v, isExist := map2[key]; isExist {
		return v
	}

	rmutex2.Lock()
	defer rmutex2.Unlock()

	map2[key] = value
	return map2[key]
}

//!!!惊天大问题1，此方法中i会被协程读取到10！！！理论上只期望协程读到0-9.
//问题2，如果不加lock的话，这10次打印中，每次可能不一样。
func storeMapFunc3() {
	key := "C"
	num := 10

	for i := 0; i < num; i++ {
		//j := i
		go func() {
			v := getV3(key, fmt.Sprintf("i=%d", i))
			fmt.Printf("%s-%s\r\n", key, v)
		}()
	}
	time.Sleep(time.Second * 1)
}

func getV3(key string, value string) string {
	map3Lock.Lock()
	defer map3Lock.Unlock()
	if v, isExist := map3.Load(key); isExist {
		return v.(string)
	}

	map3.Store(key, value)
	return value
}
