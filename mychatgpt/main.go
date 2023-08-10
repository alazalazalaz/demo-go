package main

import (
	"demo/mychatgpt/consts"
	"demo/mychatgpt/models"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	//listModels()

	//models.CreateCompletion()

	//models.CreateChatCompletion("chinese", []string{"english"})
	for i := 0; i < 1; i++ {
		models.CreateChatCompletion("", []string{"ja"}, "你好")
	}
}

func listModels() error {
	req, err := http.NewRequest("GET", consts.OPENAI_API_V1+consts.OPENAI_MODEL_MODELS, nil)
	if err != nil {
		return err
	}

	req.Header.Add("Authorization", "Bearer "+consts.OPENAI_API_KEY)
	req.Header.Set("Content-Type", "application/json")
	rep, err := http.DefaultClient.Do(req) //这个http.DefaultClient可以自定义一个，避免每次都去初始化一个client
	if err != nil {
		log.Printf("http do failed err=%s", err.Error())
		return err
	}
	defer rep.Body.Close()
	resultBuf, err := ioutil.ReadAll(rep.Body)
	if err != nil {
		log.Printf("read failed:%v", err)
		return err
	}

	fmt.Println(string(resultBuf))
	return nil
}
