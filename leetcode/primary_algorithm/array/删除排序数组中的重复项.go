package main

import "fmt"

func main(){
	rawNum := make([]int, 1, 1)
	rawNum = []int{}
	num := removeDuplicates(rawNum)
	fmt.Println(num)
}

func removeDuplicates(nums []int) int {
	i := 0
	for j := 1; j <= len(nums) - 1; j++{
		if nums[i] == nums[j] {

		}else{
			i++
			nums[i] = nums[j]
		}
	}
	i++
	fmt.Println(i)
	nums = nums[:i]
	fmt.Println(nums)
	return i
}