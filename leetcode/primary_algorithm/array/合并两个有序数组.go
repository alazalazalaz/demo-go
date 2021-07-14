package main

import "fmt"

func main(){
	nums1 := []int{4, 5, 6, 0, 0, 0}
	nums2 := []int{1, 2, 3}
	merge(nums1, 3, nums2, 3)
	fmt.Println(nums1)
}

func merge(nums1 []int, m int, nums2 []int, n int)  {
	maxNum := m + n
	for i:=maxNum - 1 ; i>=0 ; i--{
		if m == 0{
			nums1[i] = nums2[n-1]
			n--
			continue
		}
		if nums1[m-1] >= nums2[n-1]{
			nums1[i] = nums1[m-1]
			m--
		}else{
			nums1[i] = nums2[n-1]
			n--
		}
		if  n<=0 {
			break
		}
	}
}
