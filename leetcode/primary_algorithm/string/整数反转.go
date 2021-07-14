package main

import "fmt"

func main(){
	fmt.Println(reverse(-2147483412))
}

func reverse(x int) int {
	newNum := 0
	for ; x!= 0 ; {
		lastNum := x%10
		if 10*newNum > 2147483640 || 10*newNum == 214748364 && lastNum>7{
			return 0
		}
		if 10*newNum < -2147483640 || 10*newNum == -214748364 && (lastNum==-8 || lastNum==-9){
			return 0
		}

		newNum = 10*newNum + lastNum
		x = x/10
	}
	return newNum
}
