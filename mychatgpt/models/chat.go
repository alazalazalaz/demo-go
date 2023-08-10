package models

import (
	"bytes"
	"demo/mychatgpt/consts"
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

type OpenaiApiChatTranslateResult struct {
	Result []OpenaiApiResultItem `json:"result"`
}

type OpenaiApiResultItem struct {
	DetectedSourceLanguage string `json:"detectedSourceLanguage"`
	Language               string `json:"language"`
	TranslatedText         string `json:"translatedText"`
}

func CreateChatCompletion(source string, target []string, text string) error {
	prompt := "我想开一个工商个体户，帮我取一个中文名字，主要经营日用品、小物件、装饰类销售。"
	targetString := strings.Join(target, "|")

	//if source == "" {
	//	prompt = fmt.Sprintf("Translate the following text to %s and detect what language it is in.Returns only results, no interpretation needed.\n "+
	//		"```%s```\n"+
	//		"Language:\n"+
	//		"Translated:\n", targetString, text)
	//} else {
	//	prompt = fmt.Sprintf("Translate the following text from %s to %s. Returns only results, no interpretation needed.\n "+
	//		"```%s```\n", source, targetString, text)
	//}

	//maxToken := 45 + utf8.RuneCountInString(text)*3
	data := chatRequest{
		Models: "gpt-3.5-turbo",
		Message: []chatMessage{
			{
				Role:    "system",
				Content: "帮我取名",
			},
			{
				Role:    "user",
				Content: prompt,
			},
		},
		N: 1,
		//MaxTokens: maxToken,
	}

	buf, err := json.Marshal(&data)
	if err != nil {
		log.Printf("Marshal failed, err:%v", err)
		return err
	}

	body := new(bytes.Buffer)
	body.Write(buf)
	req, err := http.NewRequest("POST", consts.OPENAI_API_V1+consts.OPENAI_MODEL_CHAT, body)
	if err != nil {
		return err
	}

	req.Header.Add("Authorization", "Bearer "+consts.OPENAI_API_KEY)
	req.Header.Add("OpenAI-Organization", consts.OPENAI_ORG_ID)
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

	//处理返回的内容，返回结果会根据是否有source来决定
	//source为空的情况：
	//一对一：Language: Chinese\nTranslated: こんにちは，
	//一对多：Language: zh (Chinese)\n\nTranslated (ja): 「あなたは幸せですか？」\nTranslated (ru): «Вы счастливы?»\nTranslated (en): \"Are you happy?\
	//有source的情况：
	//一对一：あなたは幸せですか。
	//一对多：ja: 「あなたは幸せですか」\nru: \"Вы счастливы?\"\nen: \"Are you happy?\""
	resultJson := &OpenaiApiChatTranslateResult{}
	resultData.Choices[0].Message.Content = strings.Replace(resultData.Choices[0].Message.Content, "\n\n", "\n", -1)
	if source == "" {
		reArr := strings.Split(resultData.Choices[0].Message.Content, "\n")
		if len(reArr) < 2 {
			err := errors.New("api result struct err")
			fmt.Errorf("%v", err)
			return err
		}

		detectedLanguage := ""
		for k, v := range reArr {
			if strings.HasPrefix(v, "Language: ") {
				detectedLanguage = strings.TrimPrefix(v, "Language: ")
			}
			if strings.HasPrefix(v, "Translated") && k > 0 && k <= len(target) {
				item := OpenaiApiResultItem{}
				item.DetectedSourceLanguage = detectedLanguage
				item.Language = target[k-1]
				index := strings.Index(v, ": ")
				if index < 0 {
					fmt.Errorf("api result struct index err")
					continue
				}
				item.TranslatedText = v[index+2:]
				resultJson.Result = append(resultJson.Result, item)
			}
		}

		fmt.Println("debug")
		return nil
	}

	if len(target) == 1 {
		item := OpenaiApiResultItem{}
		item.DetectedSourceLanguage = source
		item.Language = targetString
		item.TranslatedText = resultData.Choices[0].Message.Content
		resultJson.Result = append(resultJson.Result, item)
		return nil
	}

	reArr := strings.Split(resultData.Choices[0].Message.Content, "\n")
	if len(reArr) < 2 {
		err := errors.New("api result struct err")
		fmt.Errorf("%v", err)
		return err
	}

	for k, v := range reArr {
		item := OpenaiApiResultItem{}
		item.DetectedSourceLanguage = source
		item.Language = target[k]
		index := strings.Index(v, ": ")
		if index < 0 {
			fmt.Errorf("api result struct index err")
			continue
		}
		item.TranslatedText = v[index+2:]
		resultJson.Result = append(resultJson.Result, item)
	}

	fmt.Println("debug")

	log.Println(resultData)

	return nil
}
