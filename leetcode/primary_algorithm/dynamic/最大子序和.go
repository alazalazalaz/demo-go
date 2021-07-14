package main

import (
	"fmt"
	"math"
)

func main(){
	list := []int{-2,1,-3,4,-1,2,1,-5,4}
	fmt.Println("result:", maxSubArray(list))
}

func maxSubArray(nums []int) int {
	sizeN := len(nums)
	if sizeN <= 0{
		return 0
	}
	if sizeN == 1{
		return nums[0]
	}
	result := math.MinInt64
	for i :=0 ; i<sizeN ; i++{
		n1 := 0
		n := 0
		for j:=i; j<sizeN; j++{
			n = n1 + nums[j]
			n1 = n
			if n > result{
				result = n
			}
		}
	}
	return result
}
