package main

import "fmt"

func main(){
	book1 := NewBook()
	fmt.Printf("book1=%v, &book1=%v\n", &book1, book1)

	book1.Author = "wang"

	book2 := NewBook()
	fmt.Printf("book2=%v, &book2=%v\n", &book2, book2)

}

type books struct {
	Id int
	Author string
}

func NewBook() *books{
	return &books{
		Id:1,
		Author: "xiaozhang",
	}
}