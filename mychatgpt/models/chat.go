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
	"regexp"
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

func CreateChatCompletion(source string, target []string, q string) error {
	prompt := "我想开一个工商个体户，帮我取一个中文名字，主要经营日用品、小物件、装饰类销售。"
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

	//maxToken := 45 + utf8.RuneCountInString(text)*3
	data := chatRequest{
		Models: "gpt-3.5-turbo",
		Message: []chatMessage{
			{
				Role:    "system",
				Content: "",
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

	resultBuf = []byte(`{  "id": "chatcmpl-7xtStsepNIOAvOx6LHjJ4PjcroDNn",  "object": "chat.completion",  "created": 1694508915,  "model": "gpt-3.5-turbo-0613",  "choices": [    {      "index": 0,      "message": {        "role": "assistant",        "content": "SourceLanguage: Russian\nTranslation->en: Stupid fucking"      },      "finish_reason": "stop"    }  ],  "usage": {    "prompt_tokens": 58,    "completion_tokens": 12,    "total_tokens": 70  }}`)
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
	convertOpenaiResultToStruct(source, target, resultData.Choices[0].Message.Content)

	fmt.Println("debug")

	log.Println(resultData)

	return nil
}

func convertOpenaiResultToStruct(source string, targets []string, resultText string) (error, error) {

	if source == "" && len(targets) == 1 {
		//无source，一对一，大部分是这种情况
		return convertNoSourceOneToOne(targets[0], resultText)
	}

	if source == "" && len(targets) > 1 {
		//无source，一度多
		return convertNoSourceOneToMul(targets, resultText)
	}

	return nil, errors.New("convertOpenaiResultToStruct error target len")
}

//source为空的情况：
//一对一：Language: Chinese\nTranslated: こんにちは，
func convertNoSourceOneToOne(target string, resultText string) (error, error) {
	detectedLanguage, translatedText := "", ""
	matchString := fmt.Sprintf("SourceLanguage:([\\s\\S]*)Translated->[ ]*%s[ ]*[ text]*:([\\s\\S]*)", target)
	result := regexp.MustCompile(matchString).FindStringSubmatch(resultText)
	if len(result) < 3 {
		result = regexp.MustCompile("SourceLanguage" + target + "[ text]*:([\\s\\S]*)").FindStringSubmatch(resultText)
		if len(result) < 2 {
			err := errors.New("convertNoSourceOneToOne api result struct err")
			//pfctx.Warningf("%v", err)
			return nil, err
		}
		detectedLanguage, translatedText = "", result[1]
	} else {
		detectedLanguage, translatedText = result[1], result[2]
	}

	detectedLanguage = strings.ReplaceAll(detectedLanguage, "\n", "")
	translatedText = strings.ReplaceAll(translatedText, "\n", "")

	trimLetters := []string{" ", "```", "\""}
	for _, v := range trimLetters {
		detectedLanguage = strings.TrimSuffix(strings.TrimPrefix(detectedLanguage, v), v)
		translatedText = strings.TrimSuffix(strings.TrimPrefix(translatedText, v), v)
	}

	//resultJson := &typdef.OpenaiApiChatTranslateResult{}
	//item := typdef.OpenaiApiResultItem{}
	//item.DetectedSourceLanguage = detectedLanguage
	//item.Language = target[0]
	//item.TranslatedText = translatedText
	//resultJson.Result = append(resultJson.Result, item)

	return nil, nil
}

//source为空的情况：
//一对多：Language: zh (Chinese)\n\nTranslated (ja): 「あなたは幸せですか？」\nTranslated (ru): «Вы счастливы?»\nTranslated (en): \"Are you happy?\
func convertNoSourceOneToMul(target []string, resultText string) (error, error) {
	resultText = `SourceLanguage: 
fr
Translated->ja: 
前の白黒のジェネリック・シークエンスでは、ジェームズ・ボンドはMI6によって00エージェントに任命され、自己の判断に基づいて殺すことが許されるようになります。
Translated->en: 
In the pre-title black and white sequence, James Bond is going to be appointed as the 00 agent by the MI6, and thus will be allowed to kill according to his own judgment.
Translated->ar: 
في التسلسل التمهيدي بالأبيض والأسود ، سيتم تعيين جيمس بوند كعميل 00 من قبل المخابرات البريطانية ، وسيسمح له بالقتل وفقًا لحكمه الشخصي.
Translated->zh-cn: 
在黑白色的前导片序列中，詹姆斯·邦德将被MI6任命为00特工，因此将被允许根据他自己的判断进行杀戮。`
	resultText = strings.ReplaceAll(resultText, "\n\n", "\n")
	reArr := strings.Split(resultText, "\n")
	if len(reArr) < 2 {
		err := errors.New("convertNoSourceOneToMul api result struct err")
		return nil, err
	}

	detectedLanguage := ""
	pregString := "SourceLanguage:([\\s\\S]*)"
	for _, v := range target {
		pregString += "Translated->" + v + ":([\\s\\S]*)"
	}
	targetResult := regexp.MustCompile(pregString).FindStringSubmatch(resultText)
	fmt.Println(targetResult)

	for k, v := range reArr {
		if strings.HasPrefix(v, "SourceLanguage: ") {
			detectedLanguage = strings.TrimPrefix(v, "SourceLanguage: ")
			fmt.Println(detectedLanguage)
			continue
		}

		targetResult := regexp.MustCompile("Translated->([\\s\\S]*):([\\s\\S]*)").FindStringSubmatch(v)

		fmt.Println(targetResult)
		if strings.HasPrefix(v, "Translated") && k > 0 && k <= len(target) {
			//item := typdef.OpenaiApiResultItem{}
			//item.DetectedSourceLanguage = detectedLanguage
			//item.Language = target[k-1]
			//translatedArr := strings.Split(v, ": ")
			//if len(translatedArr) < 2 {
			//	err := errors.New("convertNoSourceOneToMul api result struct arr err")
			//	pfctx.Warningf("%s", err)
			//	return nil, err
			//}
			//item.TranslatedText = strings.TrimSuffix(strings.TrimPrefix(translatedArr[1], `"`), `"`)
			//
			//resultJson.Result = append(resultJson.Result, item)
		}
	}

	return nil, nil
}
