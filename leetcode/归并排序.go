package main

import "fmt"

func main(){

	testNums := []int{6, 1, 2, 7, 9, 3, 4, 5, 10, 8}
	result := _mergeSort(testNums, 0, len(testNums) - 1)
	fmt.Println("result:", result)
}

//归并排序
//步骤：
//1、确定函数，主要是参数
func _mergeSort(nums []int, leftPtr, rightPtr int)[]int{
	//2、确定终点，也就是归并到只有一个元素的时候，返回一个slice
	if leftPtr == rightPtr{
		return []int{nums[leftPtr]}
	}
	//3、确定递归公式。这个是重点。
	mid := (leftPtr + rightPtr) / 2
	leftSlice := _mergeSort(nums, leftPtr, mid)
	rightSlice := _mergeSort(nums, mid+1, rightPtr)
	return _merge(leftSlice, rightSlice)
}

func _merge(leftSlice, rightSlice []int)[]int{
	leftPtr, rightPtr, index := 0, 0, -1
	n1, n2 := len(leftSlice), len(rightSlice)
	result := make([]int, n1+n2, n1+n2)
	for leftPtr < n1 || rightPtr < n2 {
		index++
		if leftPtr == n1 && rightPtr < n2{
			//左数组越界，右数组没有越界
			result[index] = rightSlice[rightPtr]
			rightPtr++
			continue
		}
		if rightPtr == n2 && leftPtr < n1{
			//右数组越界，左数组没有越界
			result[index] = leftSlice[leftPtr]
			leftPtr++
			continue
		}
		//两个都没有越界
		if leftSlice[leftPtr] < rightSlice[rightPtr]{
			result[index] = leftSlice[leftPtr]
			leftPtr++
		}else{
			result[index] = rightSlice[rightPtr]
			rightPtr++
		}
	}
	return result
}


