package main

import (
	"fmt"
	"sort"
)

func main(){
	nums := []int{1, 3, 7, 3 ,5}
	sort.Sort(sort.IntSlice(nums))
	fmt.Println(nums)
	testResult()
}

//使用sort.Slice可以排序任意结构体的数据
//比如排序[]result
type result struct{
	key string
	value int
}
func testResult(){
	resultA := result{
		key: "test1",
		value: 10,
	}
	resultB := result{
		key: "test2",
		value: 20,
	}
	resultSlice := []result{resultA, resultB}
	fmt.Printf("排序前:%v\r\n", resultSlice)
	sort.Slice(resultSlice, func(i, j int) bool {
		return resultSlice[i].value>resultSlice[j].value //按照倒序排列
	})
	fmt.Printf("排序后:%v\r\n", resultSlice)
}

func isAnagram(s string, t string) bool {
	sliceS := []byte(s)
	sliceT := []byte(t)
	sort.Slice(sliceS, func(i, j int) bool {return sliceS[i] > sliceS[j]})
	sort.Slice(sliceT, func(i, j int) bool {return sliceT[i] > sliceT[j]})
	return string(sliceS) == string(sliceT)
}