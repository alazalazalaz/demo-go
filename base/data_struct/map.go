package main

import (
	"fmt"
	"unsafe"
)

type user struct{
	age int
}

var mapGlobal map[int]int

/*
map详解
 */
func main(){
	mapGlobal = make(map[int]int)
	mapGlobal[10] = 10
	/*
	1.声明&定义
	 */
	//定义方式1 使用map关键字，此方式未初始化，为nil map，不能用来存放键值对，会panic(assignment to entry in nil map)
	// var variable_name map[key_type]value_type

	//定义方式2 使用make函数，会被初始化
	// var variable_name := make(map[key_type]value_type)

	countryCapitalMap := make(map[string]string)

	countryCapitalMap["中国"] = "北京"
	countryCapitalMap["葡萄牙"] = "里斯本"
	countryCapitalMap["西班牙"] = "马德里"

	/*
	2.遍历
	 */
	//注意：遍历输出元素的顺序与填充顺序无关！！！
	for key := range countryCapitalMap{
		fmt.Printf("key: %s, value: %s \n", key, countryCapitalMap[key])
	}

	for key, v := range countryCapitalMap{
		fmt.Printf("key: %s, value: %s \n", key, v)
	}

	/*
	3.判断某元素是否存在map中
	 */
	value, isExist := countryCapitalMap["美国"]

	fmt.Printf("value : %s, isExist : %v\n", value, isExist)//value : , isExist : false

	clientId := "test"
	clientId, _ = countryCapitalMap["sadfsdfl"]
	fmt.Printf("是否会覆盖clientId, %s\n", clientId)

	/*
	4.删除元素
	 */
	delete(countryCapitalMap, "西班牙")
	delete(countryCapitalMap, "中国")
	delete(countryCapitalMap, "葡萄牙")
	delete(countryCapitalMap, "葡萄牙3")
	fmt.Println("删除后：", countryCapitalMap, len(countryCapitalMap))
	for country := range countryCapitalMap{
		fmt.Printf("删除后 key: %s, value: %s \n", country, countryCapitalMap[country])
	}

	/*
	5.map的结构是一个指针
	*/
	var mapX = make(map[int]int, 3)
	fmt.Printf("unsafe.sizeof(mapX)=%d\r\n", unsafe.Sizeof(mapX))//输出8
	var mapY = map[int]int{
		1: 1,
	}
	fmt.Printf("unsafe.sizeof(mapX)=%d\r\n", unsafe.Sizeof(mapY))//输出8

	/*
	6. map是引用类型吗
	 */
	//普通赋值
	mapLeft, mapRight := make(map[int]int), make(map[int]int)
	mapLeft[1] = 1
	mapRight[2] = 2
	fmt.Printf("mapLeft=%v, mapLeft's *hmap=%p, mapLeft address=%p\n", mapLeft, mapLeft, &mapLeft)
	fmt.Printf("mapRight=%v, mapRight's *hmap=%p, mapRight address=%p\n", mapRight, mapRight, &mapRight)
	mapLeft = mapRight
	fmt.Printf("mapLeft=%v, mapLeft's *hmap=%p, mapLeft address=%p\n", mapLeft, mapLeft, &mapLeft)

	var mapUninit map[int]int
	mapInited := make(map[int]int)
	mapInited[1] = 1
	fmt.Printf("mapUninit=%v, mapUninit's *hmap=%p, mapUninit address=%p\n", mapUninit, mapUninit, &mapUninit)
	fmt.Printf("mapInited=%v, mapInited's *hmap=%p, mapInited address=%p\n", mapInited, mapInited, &mapInited)
	_referenceMap(mapUninit, mapInited)

	/**
	7. 函数传递和赋值
	 */
	mapA, mapB := make(map[int]int), make(map[int]int)
	mapA[1] = 1

	mapB[1] = 100
	mapB[2] = 200
	mapA = mapB
	mapB[2] = 2000
	fmt.Printf("mapA: %v, mapB: %v\n", mapA, mapB)//mapA: map[1:100 2:2000], mapB: map[1:100 2:2000]
	//说明A和B都是指向同一块数据

	mapC := make(map[int]int)
	mapC[1] = 1
	testPass(mapC)
	fmt.Printf("mapC: %v\n", mapC)//map[1:1 2:2]
	mapGlobal[20] = 20
	fmt.Printf("mapC: %v\n", mapC)//map[1:1 2:2]
}

func testPass(m map[int]int){
	m[2] = 2//这个赋值和下面的赋值是不同的！不知道slice有没有这个问题
	m = mapGlobal
	fmt.Printf("mapC inner: %v\n", m)//map[10:10]
}

func _referenceMap(mapUninit map[int]int, mapInited map[int]int){
	fmt.Printf("in func mapUninit=%v, mapUninit's *hmap=%p, mapUninit address=%p\n", mapUninit, mapUninit, &mapUninit)
	fmt.Printf("in func mapInited=%v, mapInited's *hmap=%p, mapInited address=%p\n", mapInited, mapInited, &mapInited)
	//m2 := make(map[int]int, 3)
	//m = m2
	//fmt.Printf("mapInited=%v, mapInited's *hmap=%p, mapInited address=%p\n", mapInited, mapInited, &mapInited)
}