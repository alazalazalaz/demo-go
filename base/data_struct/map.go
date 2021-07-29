package main

import (
	"fmt"
	"unsafe"
)

/*
map详解
 */
func main(){
	/*
	1.声明&定义
	 */
	//定义方式1 使用map关键字，此方式未初始化，为nil map，不能用来存放键值对
	// var variable_name map[key_type]value_type

	//定义方式2 使用make函数
	// var variable_name := make(map[key_type]value_type)

	var countryCapitalMap = make(map[string]string)

	countryCapitalMap["中国"] = "北京"
	countryCapitalMap["葡萄牙"] = "里斯本"
	countryCapitalMap["西班牙"] = "马德里"

	/*
	2.遍历
	 */
	//注意：遍历输出元素的顺序与填充顺序无关！！！
	for country := range countryCapitalMap{
		fmt.Printf("key: %s, value: %s \n", country, countryCapitalMap[country])
	}

	/*
	3.判断某元素是否存在map中
	 */
	value, isExist := countryCapitalMap["美国"]

	fmt.Printf("value : %s, isExist : %v\n", value, isExist)//value : , isExist : false

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
	5.函数内外的修改
	*/
	var mapX = make(map[int]int, 3)
	fmt.Printf("unsafe.sizeof(mapX)=%d\r\n", unsafe.Sizeof(mapX))
	var mapY = map[int]int{
		1: 1,
	}
	fmt.Printf("unsafe.sizeof(mapX)=%d\r\n", unsafe.Sizeof(mapY))

}
