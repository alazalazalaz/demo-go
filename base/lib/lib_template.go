package main

import (
	"bytes"
	"html/template"
	"log"
)

func main(){
	templateName := "www.baidu.com/{{.GameId}}/{{.ServerId}}/abc"
	temp, err := template.New("queueName").Parse(templateName)
	if err != nil {
		log.Fatalf("template.New error, err:%v", err)
	}

	templateStruct := struct {
		GameId string
		ServerId string
	}{
		GameId: "1",
		ServerId: "2",
	}
	buf := bytes.NewBufferString("")
	if err := temp.Execute(buf, templateStruct); err != nil {
		log.Fatalf("execute err:%v", err)
	}

	log.Printf("success")
	log.Println(buf.String())
}