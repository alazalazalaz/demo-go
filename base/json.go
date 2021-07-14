package main 

import(
	"fmt"
	"encoding/json"
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

	// fmt.Println(dog, cat, bird)

}


func jsonEncode(){

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