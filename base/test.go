package main

import "fmt"

func main() {
	var needDeleteTokens, needDeleteTokensRetry []string
	needDeleteTokens = getNeed()
	needDeleteTokensRetry = []string{"ff"}
	needDeleteTokens = append(needDeleteTokens, needDeleteTokensRetry...)
	fmt.Println(needDeleteTokens)
}

func getNeed() []string {
	return nil
}
