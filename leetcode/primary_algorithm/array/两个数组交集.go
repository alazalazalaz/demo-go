package main

import "fmt"

func main(){
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}
	result := intersect(nums1, nums2)
	fmt.Println(result)
}

func intersect(nums1 []int, nums2 []int) []int {
	var result []int
	result = myFind(nums1, nums2, result)
	return result
}

func myFind(nums1 []int, nums2 []int, result []int) []int{
	len1, len2 := len(nums1), len(nums2)
	for i := 0; i < len1; i++{
		for j := 0; j < len2; j++{
			if nums1[i] == nums2[j]{
				result = append(result, nums1[i])
				fmt.Println(result)
				if i == len1 - 1{
					nums1 = nums1[:i]
				}else{
					temp := i+1
					nums1 = append(nums1[:i], nums1[temp:]...)
				}
				if j == len2 - 1{
					nums2 = nums2[:j]
				}else{
					temp2 := j+1
					nums2 = append(nums2[:j], nums2[temp2:]...)
				}
				re := myFind(nums1, nums2, result)
				return re
			}
		}
	}
	return result
}
