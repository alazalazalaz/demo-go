package main

import "fmt"

func main(){
	mylist := []int{2, 3, 1}
	nextPermutation(mylist)
	fmt.Println("result:", mylist)
}

func nextPermutation(nums []int)  {
	n := len(nums)
	if n <= 1{
		return
	}
	begin := 0
	rightMin := nums[n-1]
	for i := n-1; i>0; i--{
		if nums[i-1] < nums[i]{
			if i == n -1{
				nums[i-1], nums[i] = nums[i], nums[i-1]
			}else{
				nums[i-1], rightMin = rightMin, nums[i-1]
			}
			begin = i
			_update(nums, begin)
			return
		}
	}
	_update(nums, begin)
	return
}

func _update(nums []int, begin int){
	for i:= begin; i< len(nums); i++{
		for j:= i+1; j< len(nums); j++{
			if nums[j]<nums[i]{
				nums[j], nums[i] = nums[i], nums[j]
			}
		}
	}
	return
}
