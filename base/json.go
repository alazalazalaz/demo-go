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
	// 测试json marshal
	//testJson()
	//
	//// 测试 json unmarshal 第一个参数为空，是否会报错
	//testJsonUnmarshal()
	//
	//// 测试interface{}的解析
	//testInterface()

	// 测试解析带斜杠的字符串
	//testSlash()

	ttt()
}

func testJson() {
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

type testSlashStruct struct {
	Action struct {
		Data struct {
			AcqCode string `json:"acqCode"`
		} `json:"data"`
	} `json:"action"`
}

func testSlash() {
	s := `{\"action\":{\"data\":{\"acqCode\":\"28470528\",\"backEndUrl\":\"https://test.adyen.com/hpp/nUP.shtml?pp=IntDgsB%2FjYwpkzhO5bjJGXSvnmpuBgxkV2LDmzOjp%2BQue4cLbdboxHFoC96o6wSsOI5eXi4pMAkdww7%2BZcxHnDJYeGc8TIzbBFlnJKSGAT6tyArbCSrldDEzS%2BmRO4GJ\&orderTime=20220513072513\",\"charset\":\"UTF-8\",\"frontEndUrl\":\"https://checkoutshopper-test.adyen.com/checkoutshopper/checkoutPaymentReturn?gpid=GP031997F771DB59C0\&orderTime=20220513072513\",\"merAbbr\":\"NotWorking\",\"merCode\":\"5718\",\"merId\":\"833000000000001\",\"merReserved\":\"{isPreAuth=false\&frontFailUrl=https://checkoutshopper-test.adyen.com/checkoutshopper/checkoutPaymentReturn?gpid=GP031997F771DB59C0\&ady_status=cancelled}\",\"orderAmount\":\"10\",\"orderCurrency\":\"978\",\"orderNumber\":\"AU4L34EX0UW5\",\"orderTime\":\"20220513072513\",\"signMethod\":\"MD5\",\"signature\":\"3edafae572c66f376d2305ab4209460e\",\"transType\":\"01\",\"version\":\"1.0.0\"},\"method\":\"POST\",\"paymentMethodType\":\"unionpay\",\"type\":\"redirect\",\"url\":\"https://checkoutshopper-test.adyen.com/checkoutshopper/checkoutPaymentRedirect?redirectData=X3XtfGC7%21H4sIAAAAAAAAAJWUbW%2FaMBSF%2F4u%2FrgTbeSuRpqkNKWS8DkGrTUiTm7iQERzXcaAZ6n%2FfNdDRdmPSPoCU6yfn%2BhxfZ4c6ecUlW%2FA20wwFO3SvmEjDIuUoQJXICiFZjS4Qf5KZ4uWVhjLFlDaw2yDOlDgBdQNif8A0wBi4LAWgM8Y2abX8G98n7Wu3FZqVNVfJkgl9lSRFJYzO9Grs3MyG3dGw04NfPx7E06gdhaMB4Iqn0DDRY1avudAv22PJYwV1dVSJoR1xXEIuEFsfZHcoqZTiIqmhRTSbgNaGgUkA8TNgMrvlqgRjKPB88yzzLGEaCrF4KPY90pqLfgZJqNo8C7Y2aezLjUXRAIlGflwG8Rc15FrYougZmpwJ8T8jeMEn%2FIEbQ0aNEGp7hHgudRzHx55NHQ%2FQQmWLTLB8XMrX9Gdv0G4NPO%2Bu2%2B11epcUSMnyCS%2BrXL9Dp9f996jijxUvYbeHyF%2FTZgSwS2zsU%2Fh%2FxaZpZqJk%2Be8De5uvdQxu%2BM9Mz7x0%2By7qU0Sx0Hyh9i9YupZGOuxGYW80m37vRMNoEodn6NPxeT56NkZ0pcRM5VBZai3LYN6cN5MlT1ZFpctlISVXDQ1erf3mraRY%2F7F%2BKhzHd7JX%2FbSQWfrxr7ej5DkMO0%2BvzeS8nZpylYnjLMnq3tpQ6xITzyYexj52CHaptbWXyTfGl7hDNl%2BKyehmtvmKL6d1Fd2OV164crrjPt2QHz%2Fd%2BODxcLdOLvcmt9utBd%2BC2sqfUlEejM1kIe%2F4%2FbwJBzRvghuLJSY3dC4pI6KZdB6qYzRwIX4B5U%2FgDmYEAAA%3DYYFBR7o0Sz7yYPV3HGTGS%2FMUdTj2TGXp4Uli8XDhjys%3D\"},\"resultCode\":\"RedirectShopper\"}`

	//json.RawMessage{}
	data := &testSlashStruct{}
	if err := json.Unmarshal([]byte(s), data); err != nil {
		log.Fatalln(err)
	}

	log.Println(data)
}

func ttt() {
	type student struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
		ID   int    `json:"id"`
		SID  int    `json:"sid"`
	}
	msg := ""
	var someOne student
	if err := json.Unmarshal([]byte(msg), &someOne); err == nil {
		fmt.Println(someOne)
	} else {
		fmt.Println(err)
	}
}
