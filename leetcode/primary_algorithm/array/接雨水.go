package main

import "fmt"

func main(){
	a := []int{0,1,0,2,1,0,1,3,2,1,2,1}
	fmt.Println("result:", trap(a))
}
func trap(height []int) int {
	leftPtr := 0
	total := 0
	n := len(height)
	if n <= 2{
		return 0
	}
	inner := 0
	for rightPtr := 1; rightPtr<n; rightPtr++{
		//如果右边有柱子,并且柱子比当前的高，则左指针往右边挪
		if leftPtr+1 < n && height[leftPtr+1] >= height[leftPtr]{
			leftPtr = leftPtr+1
		}
		//如果右指针每次都要挪动，就放for里面了
		//但是需要计算是否有inner
		if height[rightPtr] < height[leftPtr]{
			inner += height[rightPtr]
		}
		//判断是否有水可取
		if rightPtr - leftPtr > 1 && height[rightPtr] >= height[leftPtr]{
			tmp := _min(height[leftPtr], height[rightPtr]) * (rightPtr - leftPtr - 1) - inner
			total += tmp
			fmt.Println(tmp, total)
			inner = 0
			leftPtr = rightPtr
		}
	}
	return total
}

func _min(a, b int)int{
	if a < b{
		return a
	}
	return b
}
