//结构体是表示相同或不同数据类型的数据集合

package main 
import "fmt"

/*
1.定义
 */
type Books struct{
	title string 
	author string 
	bookId int
}

func main(){
	var book1 Books

	book1.title = "我是book 1"
	book1.author = "www"
	book1.bookId = 1

	fmt.Printf("book1 title:%s\n", book1.title)
	fmt.Printf("book1 author:%s\n", book1.author)
	fmt.Printf("book1 bookId:%d\n", book1.bookId)

	/*
	2. 匿名定义
	 */
	person := struct {
		name string
		age int
	}{}
	person.name = "小明"
	person.age = 10
	fmt.Println(person)

	/*
	3.匿名定义+初始化
	 */
	person2 := struct{
		name string
	}{
		name : "小明",
	}
	fmt.Println(person2)

	fmt.Println(NewCatByName("judy"))
	fmt.Println(NewCatByColor("red"))
}

type Cat struct {
	Color string
	Name  string
}

func NewCatByName(name string) *Cat {
	return &Cat{
		Name: name,
	}
}

func NewCatByColor(color string) *Cat {
	return &Cat{
		Color: color,
	}
}