// func func_name(params)[return_types]{}

package main 
import "fmt"

func main() {
	fmt.Printf("%d", max(1, 2))

	a, b := swapXY("baidu", "laji")
	fmt.Println(a, b)

	
}

//普通实例
// func max(num1 int, num2 int) int {}
func max(num1, num2 int) int {
	var result int

	if num1 > num2 {
		result = num1
	}else{
		result = num2
	}

	return result
}


//返回多个值
func swapXY(x, y string)(string, string){
	return y, x
}









