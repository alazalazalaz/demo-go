package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"log"
)

type body2 struct{
	OrderId string `json:"order_id"`
}

func main(){
	bodyBytes := []byte("abc")
	bodyNumBytes := []byte("123")

	m := md5.New()
	m.Write(bodyBytes)
	configMd5Str := hex.EncodeToString(m.Sum(nil))
	m.Write(bodyNumBytes)//再次调用m.Write()相当于追加字符串，也就是 cacheMd5Str 会等于md5("abc123")
	cacheMd5Str := hex.EncodeToString(m.Sum(nil))
	log.Printf("Md5:%v\n %v\n", configMd5Str, cacheMd5Str)

	bodyNum2Bytes := []byte("123")
	m2 := md5.New()
	m2.Write(bodyNum2Bytes)
	bodyNum2Md5 := hex.EncodeToString(m2.Sum(nil))
	log.Println(bodyNum2Md5)

	order := &body2{
		OrderId: "1",
	}

	jsonBytes, err := json.Marshal(order)
	log.Println(order, jsonBytes, err) //&{1} [123 34 79 114 100 101 114 73 100 34 58 34 49 34 125] <nil>
	log.Println(string(jsonBytes)) //{"OrderId":"1"}
}