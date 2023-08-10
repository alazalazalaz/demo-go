package main

import (
	"demo/myfile/file"
	"fmt"
	"os"
)

func main() {
	filePath := "myfile/data/translatemapping.txt"
	lines, err := file.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	writePath := "myfile/data/translatemapping-result.txt"
	resultFile, err := os.OpenFile(writePath, os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		fmt.Println("open file failed,err:", err)
		return
	}
	defer resultFile.Close()

	for _, v := range lines {
		n, err := resultFile.WriteString(v + `"},` + "\r\n")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(n)
	}
}
