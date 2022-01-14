package main

import (
	"fmt"
	"log"
	"strings"
)

func main() {
	sp("") //结果长度为1的空数组 result[0]等于空
	sp("1")
	sp("1,2")
	sp("1,2,3")
	sp("1,,2")

	trim()

}

func sp(str string) {
	_str := strings.Split(str, ",")
	fmt.Println(_str[0], str, _str, len(_str))
}

func trim() {
	s := "  abc  d "
	log.Println(fmt.Sprintf("-%s-", strings.Trim(s, " ")))
}
