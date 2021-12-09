package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Test struct{

}

type Animal struct{
	Id 	int
	Name string
}

func main(){
	base := &Animal{Id:0, Name:""}
	baseString, _ := json.Marshal(base)

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


func testJsonUnmarshal(){
	s := ``// s为空，还真会报错
	animal := &Animal{}
	err := json.Unmarshal([]byte(s), animal)
	fmt.Println(err) //unexpected end of JSON input
}

func jsonDecode(baseString []byte, i int){
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

func testInterface(){
	ss := `{}`
	data := &message{}
	if err := json.Unmarshal([]byte(ss), data); err != nil {
		log.Fatalln(err)
	}

	log.Fatalln(data)
}