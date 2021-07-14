package main

import "fmt"

func main(){
	list:= []int{7, 1, 5, 3, 6, 4}
	fmt.Println("result:", maxProfit(list))
}

func maxProfit(prices []int) int {
	n := len(prices)
	if n <= 1{
		return 0
	}
	pastMaxProfit, todayMaxProfit, pastLowestPrice,result := 0, 0, prices[0], 0
	for i:=0; i<n; i++{
		pastLowestPrice = _min(prices[i], pastLowestPrice)
		todayMaxProfit = prices[i] - pastLowestPrice
		result = _max(pastMaxProfit, todayMaxProfit)
		pastMaxProfit = result
	}
	return result
}

func _min(a int, b int)int{
	if a>b{
		return b
	}
	return a
}
func _max(a int, b int)int{
	if a>b{
		return a
	}
	return b
}
