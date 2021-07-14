package main 

import(
	"fmt"
	"time"
	ra "math/rand"
)


func main(){

	// fmt.Println(Random(0, 0))//报错
	fmt.Println(Random(0, 1))//结果只有0
	fmt.Println(Random(0, 2))//结果为0或者1

}

//@return 左闭右开，结果会包含min，不会有max
func Random(min, max int) int {
	ra.Seed(time.Now().Unix())
	return ra.Intn(max-min) + min
}