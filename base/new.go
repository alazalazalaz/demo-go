package main 

import (
	"fmt"
)

func main(){
	obj := new(Person)
	fmt.Println(obj.Id)
}

type Person struct{
	Id 	int 
	Name string
}

func (this *Person) List(){
	fmt.Println("list")
}


