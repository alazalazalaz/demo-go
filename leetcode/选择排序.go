package main

import "fmt"

func main(){

	testNums := []int{4, 3, 2, 6, 1, 4}
	_selectionSort(testNums)
	fmt.Println("result:", testNums)
}

func _selectionSort(nums []int){
	n := len(nums)
	for i:= 0; i<n-1; i++{
		minValue := nums[i]
		for j:=i+1; j<n; j++{
			if nums[j] < minValue{
				minValue = nums[j]
				nums[j], nums[i] = nums[i], nums[j]
			}
		}
	}
}
