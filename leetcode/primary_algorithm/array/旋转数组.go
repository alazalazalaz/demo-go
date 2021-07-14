package main

import "fmt"

func main(){
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(nums, 3)
	fmt.Println(nums)
}

func rotate(nums []int, k int)  {
	if len(nums) <= 1{
		return
	}
	numsTemp := make([]int, len(nums), len(nums))
	copy(numsTemp, nums)
	nextIndex := 0
	for i:= 0; i < len(nums); i++{
		temp := numsTemp[i]
		if i + k > len(nums){
			nextIndex = (i + k) % len(nums)
		} else if i + k == len(nums){
			nextIndex = 0
		} else {
			nextIndex = i + k
		}
		nums[nextIndex] = temp
	}
}