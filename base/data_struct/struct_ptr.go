package main 
import "fmt"

type Book struct{
	title string
	author string
	bookId int
}

func main(){

	var book Book 
	book.title = "title"
	book.author = "author"
	book.bookId = 1

	var num int = 10
	printBook(&book, &num)
}

func printBook(book *Book, num *int){
	//注意此时使用形参book或者num时，前面可以不需要加*，编译器会自动隐式的转换该指针对应的值。
	fmt.Printf("title: %s\n", (*book).title)
	fmt.Printf("author: %s\n", book.author)
	fmt.Printf("bookId: %d\n", book.bookId)
	fmt.Printf("num: %d\n", *num)
	fmt.Printf("num: %d\n", num)
}