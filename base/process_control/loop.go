package main

import (
	"log"
)

func main() {
	testLoop()
}

// LOOP1标签后面得跟上for循环，
func testLoop() {
	log.Println("testLoop begin")
LOOP1:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			log.Printf("i=%d, j=%d", i, j)
			if j > 3 {
				break LOOP1 //此处break会跳出整个LOOP1大循环，而不只是for j的循环。
			}
		}
	}

	log.Println("testLoop end")

}
