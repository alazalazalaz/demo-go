
// 声明多维数组
// var variable_name [SIZE1][SIZE2]...[SIZEN] variable_type

package main 
import "fmt"




func main(){
	//定义一个二维数组
	var a [3][4]int

	//初始化
	a = [3][4]int{
		{0, 1, 2, 3},
		{4, 5, 6, 7},
		{8, 9, 10, 11},//此处必须有逗号
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 4; j++ {
			fmt.Printf("element a[%d][%d] = %d \n", i, j, a[i][j])
		}
	}
}

