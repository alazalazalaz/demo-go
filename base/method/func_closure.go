package main 
import "fmt"

func main(){
	nextNum := getNum()//闭包哦
	fmt.Println(nextNum())//返回1
	fmt.Println(nextNum())//2
	fmt.Println(nextNum())//3

	nextNum2 := getNum()
	fmt.Println(nextNum2())
	fmt.Println(nextNum2())
}

func getNum() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}