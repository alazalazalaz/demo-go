package main

import "fmt"

func main(){
	digits := []int{1, 2, 3}
	result := plusOne(digits)
	fmt.Println(result)
}

func plusOne(digits []int) []int {
	len1 := len(digits)
	var newArray []int
	j , addOne := 0, 0
	for i := len1 - 1; i >= 0; i--{
		temp := digits[i] + addOne
		if i == len1 - 1{
			temp = digits[i] + 1
		}

		if temp == 10{
			newArray = append(newArray, 0)
			addOne = 1
		}else{
			newArray = append(newArray, temp)
			addOne = 0
		}
		j++
	}
	if addOne == 1{
		newArray = append(newArray, 1)
	}
	var resultArray []int
	len2 := len(newArray)
	k := 0
	for j := len2 - 1; j >= 0; j--{
		resultArray = append(resultArray, newArray[j])
		k++
	}
	return resultArray
}