package main

import "demo/myazuregpt/models"

func main() {
	models.Chat("", []string{"ja", "en"}, "你好")
}
