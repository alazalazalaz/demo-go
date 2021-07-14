package main 

import (
	"errors"
	"fmt"
)


func main(){
	_, err := Sqrt(-1)
	if err != nil {
		fmt.Println(err)
	}
}

func Sqrt(num float64) (float64, error){
	if num < 0 {
		return 0, errors.New("math: square root of negative number")
	}

	return 0, nil
}