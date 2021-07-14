package main 

import(
	"fmt"
)

func main(){

	type human struct{
		name string 
		age int
	}
	
	allen := human{}

	james := map[string]string{}

	type data int 
	var kebo interface{}
	fmt.Println(allen, kebo, james, data)
}