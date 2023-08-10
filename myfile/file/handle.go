package file

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func ReadFile(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Printf("open file error:%v", err)
		return nil, err
	}
	defer file.Close()
	content, err := ioutil.ReadAll(file)
	contentArray := strings.Split(string(content), "\n")
	return contentArray, nil
}
