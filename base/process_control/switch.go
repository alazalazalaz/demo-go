package main 
import "fmt"

func main() {
	a := 1
	switch a{
	case 1:
		fmt.Printf("1\n")
	case 10:
		fmt.Printf("2\n")
	case 100:
		fmt.Printf("3\n")
	default:
		fmt.Printf("none\n")
	}


	marks := "A"
	switch {
	case marks == "A":
		fmt.Printf("A\n")
	case marks == "B":
		fmt.Printf("B\n")
	default:
		fmt.Printf("others\n")
	}
}


