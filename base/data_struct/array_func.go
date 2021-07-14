package main 
import "fmt"

func main(){
	var array = [5]int{1, 2, 3, 4, 6}
	var avg float32
	avg = getAverage(array, 5)

	fmt.Printf("average is : %v \n", avg)
}

func getAverage(arr [5]int, size int) float32 {
	var sum int
	var avg float32

	for i := 0; i < size; i++ {
		sum += arr[i]
	}

	fmt.Println(sum, size)
	avg = float32(sum) /float32(size)
	// avg = sum / size//这样会报错，只允许同类型的变量进行运算
	return avg
}

