package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Test struct {
}

type Animal struct {
	Id   int
	Name string
	Json string
}

type Animal2 struct {
	Id   int
	Name string
	Body MyBody `json:"json"`
}

type MyBody struct {
	Aaa string
	Bbb string
}

func main() {
	b2 := MyBody{Aaa: "11", Bbb: "22"}
	a2 := &Animal2{
		Id:   11,
		Name: "xx",
		Body: b2,
	}
	a2string, _ := json.Marshal(a2)
	log.Println(string(a2string))

	base := &Animal{Id: 333, Name: "xixixixi", Json: `{"a":"b"}`}
	baseString, _ := json.Marshal(base)
	log.Println(string(baseString)) //{"Id":0,"Name":""}
	baseString2, _ := json.Marshal(string(baseString))
	log.Println(string(baseString2)) //"{\"Id\":0,\"Name\":\"\"}"
	baseString3, _ := json.Marshal(string(baseString2))
	log.Println(string(baseString3)) //"\"{\\\"Id\\\":333,\\\"Name\\\":\\\"xixixixi\\\"}\""

	base1 := &Animal{}
	if err := json.Unmarshal(baseString, base1); err != nil {
		log.Fatalln(err)
	}
	log.Println(base1)

	base2 := &Animal{}
	if err := json.Unmarshal(baseString2, base2); err != nil {
		log.Fatalln(err)
	}
	log.Println(base2)
	// dog := &Animal{Id:1, Name:""}
	// cat := &Animal{Id:2, Name:""}
	// bird := &Animal{Id:3, Name:""}

	for i := 0; i < 3; i++ {
		jsonDecode(baseString, i)
	}

	// 测试 json unmarshal 第一个参数为空，是否会报错
	testJsonUnmarshal()

	// 测试interface{}的解析
	testInterface()

}

func testJsonUnmarshal() {
	s := `` // s为空，还真会报错
	animal := &Animal{}
	err := json.Unmarshal([]byte(s), animal)
	fmt.Println(err) //unexpected end of JSON input
}

func jsonDecode(baseString []byte, i int) {
	var ani Animal
	json.Unmarshal((baseString), &ani)

	if i == 0 {
		ani.Id = 1
		ani.Name = "dog"
	}

	fmt.Println(ani)
}

type message struct {
	Extension map[string]interface{}
}

func testInterface() {
	ss := `{}`
	data := &message{}
	if err := json.Unmarshal([]byte(ss), data); err != nil {
		log.Fatalln(err)
	}

	log.Fatalln(data)
}
