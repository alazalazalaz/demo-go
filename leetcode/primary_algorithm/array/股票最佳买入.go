package main

import "fmt"

func main(){
	p := maxProfit([]int{4, 3, 2, 1})
	fmt.Println(p)
}

func maxProfit(prices []int) int {
	totalProfit := 0
	for i := 0; i < len(prices) - 1; i++{
		if prices[i+1] >= prices[i] {
			totalProfit += prices[i+1] - prices[i]
		}
	}
	return totalProfit
}
