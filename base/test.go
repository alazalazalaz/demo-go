package main

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

func main() {
	data := "Language: \nEnglish\n\nTranslated text:```\n Convide Mo Lead Kabayan Natin ```"
	language, text := "", ""
	result := regexp.MustCompile("Language:([\\s\\S]*)Translated[ text]*:([\\s\\S]*)").FindStringSubmatch(data)
	if len(result) < 3 {
		result = regexp.MustCompile("Translated[ text]*:([\\s\\S]*)").FindStringSubmatch(data)
		if len(result) < 2 {
			fmt.Println("err")
			return
		}
		language, text = "", result[1]
	} else {
		language, text = result[1], result[2]
	}

	language = strings.ReplaceAll(language, "\n", "")
	text = strings.ReplaceAll(text, "\n", "")

	trimLetters := []string{" ", "```", "\""}
	for _, v := range trimLetters {
		language = strings.TrimSuffix(strings.TrimPrefix(language, v), v)
		text = strings.TrimSuffix(strings.TrimPrefix(text, v), v)
	}

	fmt.Println(language)
	fmt.Println(text)
	fmt.Println(time.Now().Format(time.StampMilli))
	time.Sleep(time.Second * 1)
	fmt.Println(time.Now().Format(time.StampMilli))
}
