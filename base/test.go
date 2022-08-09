package main

import (
	"fmt"
	"log"
	"strings"
)

func main() {
	log.Println(IsVailedIp("101.24.90.164"))
}

func IsVailedIp(ip string) bool {
	ids := ""
	idsArray := strings.Split(ids, ",")
	fmt.Println(idsArray)

	idsArray = append(idsArray, "aaa")

	fmt.Println(idsArray)

	newIds := strings.Join(idsArray, ",")

	fmt.Println(newIds)
	newIds = strings.TrimLeft(newIds, ",")

	fmt.Println(newIds)
	return true
}
