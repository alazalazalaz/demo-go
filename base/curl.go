package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	//rawCurl()

	httpCurl()
}

type pushBody struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func httpCurl() error {
	data := pushBody{
		Id:   1,
		Name: "test",
	}
	buf, err := json.Marshal(&data)
	if err != nil {
		log.Printf("Marshal failed, err:%v", err)
		return err
	}

	body := new(bytes.Buffer)
	body.Write(buf)
	req, err := http.NewRequest("POST", "https://pf-chat-en2en.tap4fun.com/k1d2-beta/3048/permanent/avatar", body)
	if err != nil {
		return err
	}
	//req.Header.Add("Authorization", "Bearer xxx")
	rep, err := http.DefaultClient.Do(req) //这个http.DefaultClient可以自定义一个，避免每次都去初始化一个client
	if err != nil {
		log.Printf("http do failed err=%s", err.Error())
		return err
	}
	defer rep.Body.Close()
	buf, err = ioutil.ReadAll(rep.Body)
	if err != nil {
		log.Printf("read failed:%v", err)
		return err
	}

	log.Println(string(buf))
	return nil
}

func rawCurl() {
	resp, err := http.Get("https://pf-chat-en2en.tap4fun.com/k1d2-beta/3048/permanent/avatar")
	if err != nil {
		fmt.Println("err")
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	fmt.Println(resp.StatusCode)
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("err:%v\n", err)
		return
	}

	fmt.Println(data)
}
