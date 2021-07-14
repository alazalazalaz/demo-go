package main 
import "fmt"

var g, g2 int = 100, 100

func main(){
	//局部变量和全局变量可以重名，但是会优先取局部变量
	var a, g int

	a = 20
	g = 20

	fmt.Println(a, g, g2)
}