package main

import "fmt"

func main() {
	//会不会掉下去的问题
	testS()
}

func testS() {
	//结论：只会打印1和3，不会打印2
	for a := 0; a < 4; a++ {
		switch a {
		case 1:
			fmt.Printf("%d\n", a)
		case 2:
		case 3:
			fmt.Printf("%d\n", a)
		}
	}

}
