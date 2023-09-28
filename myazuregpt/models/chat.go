package models

import (
	"bytes"
	"demo/myazuregpt/consts"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type chatRequest struct {
	//required
	Models  string        `json:"model"`
	Message []chatMessage `json:"messages"`

	//optional
	N int `json:"n"` //返回的object条目数，默认1
	//MaxTokens int `json:"max_tokens"` //返回的最大token数，默认16
}

type chatMessage struct {
	//required
	Role string `json:"role"`
	//optional
	Content string `json:"content"`
}

type chatResponse struct {
	Id      string `json:"id"`
	Object  string `json:"object"`
	Created int    `json:"created"`
	Model   string `json:"model"`
	Choices []struct {
		Text         string      `json:"text"`
		Index        int         `json:"index"`
		Logprobs     interface{} `json:"logprobs"`
		FinishReason string      `json:"finish_reason"`
		Message      chatMessage `json:"message"`
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

func Chat(source string, target []string, q string) error {
	prompt := "我是卖美甲饰品的，帮我写点推销术语"
	targetString := strings.Join(target, "|")
	localizationString := ""

	if source == "" {
		multiPrompt, targetPrompt := "", ""
		if len(target) > 1 {
			multiPrompt = "the target language is multi and separate with |."
		}
		for _, v := range target {
			targetPrompt += fmt.Sprintf("Translated->%s:\n", v)
		}
		prompt = fmt.Sprintf("%s Translate the text ```%s``` to target language %s and detect the source language.%s Returns only results like the format below, no interpretation needed.\n"+
			"SourceLanguage:\n"+
			"%s", localizationString, q, targetString, multiPrompt, targetPrompt)
	} else {
		prompt = fmt.Sprintf("%s Translate the following text from %s to %s.Replay me with english. Returns only results, no interpretation needed. \n ```%s```\n", localizationString, source, targetString, q)
	}

	prompt = "Translate the text ```操你妈``` to target language en and detect the source language. Returns only results like the format below, no interpretation needed."

	url := fmt.Sprintf("%s/openai/deployments/gpt-35-turbo/chat/completions?api-version=2023-05-15", consts.AZURE_ENDPOINT)
	data := chatRequest{
		Message: []chatMessage{
			{
				Role:    "system",
				Content: "you are a helpful translator",
			},
			{
				Role:    "user",
				Content: prompt,
			},
		},
		N: 1,
	}

	buf, err := json.Marshal(&data)
	if err != nil {
		log.Printf("Marshal failed, err:%v", err)
		return err
	}

	body := new(bytes.Buffer)
	body.Write(buf)
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return err
	}

	req.Header.Add("api-key", consts.AZURE_KEY)
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

	resultData := &chatResponse{}
	err = json.Unmarshal(resultBuf, resultData)
	if err != nil {
		log.Printf("Unmarshal error:%v", err)
		return err
	}

	if len(resultData.Choices) <= 0 {
		err = errors.New("len err")
		log.Printf("err:%v", err)
		return err
	}

	return nil
}
