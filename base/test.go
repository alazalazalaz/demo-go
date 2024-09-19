package main

import (
	"encoding/hex"
	"fmt"
)

func main() {
	s, err := hex.DecodeString("xx")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(s)
}
