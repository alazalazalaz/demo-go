package main

import "fmt"

func main(){

	testNums := []int{6, 1, 2, 7, 9, 3, 4, 5, 10, 8}
	_quickSort(testNums, 0, len(testNums) - 1)
	fmt.Println("result:", testNums)
}

func _quickSort(nums []int, leftPtr, rightPtr int){
	if rightPtr <= leftPtr{
		return
	}
	baseLeftPtr := leftPtr
	baseRightPtr := rightPtr
	base := nums[leftPtr]
	for leftPtr != rightPtr{
		for rightPtr > leftPtr{
			if nums[rightPtr] < base{
				break
			}
			rightPtr--
		}
		for leftPtr < rightPtr{
			if nums[leftPtr] > base{
				break
			}
			leftPtr++
		}
		if leftPtr != rightPtr{
			//交换
			nums[leftPtr], nums[rightPtr] = nums[rightPtr], nums[leftPtr]
		}
	}
	//两个指针相遇后把基数和相遇位置的值互换
	//相遇后leftPtr和rightPtr相等，都变成了那个k位置的下标
	nums[baseLeftPtr], nums[rightPtr] = nums[rightPtr], nums[baseLeftPtr]
	//左右子序列再递归
	_quickSort(nums, baseLeftPtr, rightPtr-1)
	_quickSort(nums, rightPtr + 1, baseRightPtr)
}

