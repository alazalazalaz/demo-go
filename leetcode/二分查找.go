package main

import "fmt"

func main(){
	nums := []int{-99999,-99998,-9999,-999,-99,-9,-1}
	re := searchRange(nums,  0)
	re2 := searchHalf(nums, -99998)
	fmt.Println("result:", re)
	fmt.Println("result2:", re2)
}

//递归方法
func searchRange(nums []int, target int) []int {
	n := len(nums)
	result := []int{-1, -1}
	if n <=0 {
		return result
	}
	//递归二分查找
	resultIndex, ok := _find(nums, 0, n-1, target)
	if ok {
		minIndex, maxIndex := resultIndex, resultIndex
		for i:= n-1; i>=0; i--{
			if nums[resultIndex] == nums[i]{
				if i>maxIndex{
					maxIndex = i
				}
				if i<minIndex{
					minIndex = i
				}
			}
		}
		return []int{minIndex, maxIndex}
	}
	return result
}

func _find(nums []int, begin, end, target int) (int, bool){
	if nums[begin] == target{
		return begin, true
	}
	if nums[end] == target{
		return end, true
	}
	if end - begin <= 1{
		return 0, false
	}

	mid := (end - begin)/2
	if target > nums[mid]{
		if index, ok := _find(nums, mid+1, end-1, target); ok == true{
			return index, ok
		}
	}else if target == nums[mid]{
		return mid, true
	}else{
		if index, ok:= _find(nums, begin+1, mid-1, target); ok == true{
			return index, ok
		}
	}
	return 0, false
}

//非递归方法
func searchHalf(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left < right{
		mid := ((right - left)/2) + left
		if target <= nums[mid]{
			right = mid
		}else{
			left = mid + 1
		}
	}
	return left
}