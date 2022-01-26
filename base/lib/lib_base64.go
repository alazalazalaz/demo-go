package main

import (
	"encoding/base64"
	"log"
)

func main() {
	testAllBase64Type()
}

func testAllBase64Type() {
	baseedString := "nxSBcHGu5JYZrQ7a/aztVE8prZ4Fp8HR2y5ruH9R7QKQVvqYsF40OgwjhQE0kWfxYm+gpnIDnllUBr8UcqIkIQc17gW7/nKsn1yKy8ZoKK/LWBVdC6bZacqOGY368QJo9hlIf6KcRcw0RU5qv4+FGfLbyyWf2QlJevEciXcmPBRxoaY8gFbSasFUEFtcq8ZArU/iMQQbTN9Wd45lXZT5B/ZfMffh2Dbisxxzkfxpe1u9Z+g5Q7Gqq1nyAVa+jEcRejBcHguVJx9nQQ/X8BjKM5Qh0sVoqAPa4l+HM3wVXc8="
	data, err := base64.StdEncoding.DecodeString(baseedString)
	if err != nil {
		log.Printf("error:%v\r\n", err)
	}
	log.Println(data)
}
