package models

import (
	"bytes"
	"demo/mychatgpt/consts"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
)

type completionRequest struct {
	//required
	Models string `json:"model"`
	Prompt string `json:"prompt"`
	//optional
	N         int `json:"n"`          //返回的object条目数，默认1
	MaxTokens int `json:"max_tokens"` //最大token数（prompt+返回的），默认16
}

type completionResponse struct {
	Id      string `json:"id"`
	Object  string `json:"object"`
	Created int    `json:"created"`
	Model   string `json:"model"`
	Choices []struct {
		Text         string      `json:"text"`
		Index        int         `json:"index"`
		Logprobs     interface{} `json:"logprobs"`
		FinishReason string      `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
}

type errorResponse struct {
	Error struct {
		Message string      `json:"message"`
		Type    string      `json:"type"`
		Param   interface{} `json:"param"`
		Code    interface{} `json:"code"`
	} `json:"error"`
}

func CreateCompletion() error {
	data := completionRequest{
		Models: "text-davinci-003",
		//Prompt: "解释一下openai的Chat模型和Completion模型区别",
		Prompt:    "translate the text '你个瓜批' to en、fr、jp，and split them by identifier |",
		N:         3,
		MaxTokens: 4000,
	}
	buf, err := json.Marshal(&data)
	if err != nil {
		log.Printf("Marshal failed, err:%v", err)
		return err
	}

	body := new(bytes.Buffer)
	body.Write(buf)
	req, err := http.NewRequest("POST", consts.OPENAI_API_V1+consts.OPENAI_MODEL_COMPLETIONS, body)
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

	log.Printf("result:%v", string(resultBuf))

	errorData := &errorResponse{}
	err = json.Unmarshal(resultBuf, errorData)
	if err == nil && errorData.Error.Type != "" {
		log.Printf("error type:%v, error msg:%v", errorData.Error.Type, errorData.Error.Message)
		return errors.New(errorData.Error.Message)
	}

	resultData := &completionResponse{}
	err = json.Unmarshal(resultBuf, resultData)
	if err != nil {
		log.Printf("Unmarshal error:%v", err)
		return err
	}

	log.Println(resultData)

	return nil
}
