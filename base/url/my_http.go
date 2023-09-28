package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {

	for i := 1; i < 1000; i++ {
		wg.Add(1)
		if i%5 == 0 {
			time.Sleep(time.Second)
		}
		go doHttp(i)
	}

	log.Println("waiting...")
	wg.Wait()
	log.Println("DONE...")
}

func doHttp(i int) error {
	pushClient := &http.Client{Timeout: time.Duration(30) * time.Second}

	PushUrl := "https://translate.beta.tap4hub.com/translate"

	req, err := http.NewRequest("GET", PushUrl, nil)
	if err != nil {
		log.Printf("send http err:%v", err)
		return err
	}

	params := req.URL.Query()
	params.Add("client_id", "p2:zx")
	params.Add("source", "zh")
	params.Add("target", "en")
	params.Add("serverid", "100")
	params.Add("q", fmt.Sprintf("同学们，跟我数数, %d 。", i))
	req.URL.RawQuery = params.Encode()

	req.Header.Set("authorization", "Bearer eyJhbGciOiJSUzI1NiIsImtpZCI6IktTVkNHTFhaUEQiLCJrdHkiOiJSU0EiLCJ0eXAiOiJKV1QifQ.eyJhY2NvdW50X2lkIjoiMTk3MDMyNDgzNjk3NDY1MSIsImF1ZCI6InBmLWN1c3RvbWVyIiwiZXhwIjoxNjk1NzQyMzYyNTc4LCJleHBpcmVfc3RyIjoiMTY5NTc0MjM2MjU3OCIsImdhbWVfaWQiOiJwMiIsImlwX3NlZ21lbnQiOiIwLjAuMC4wLzAsOjovMCIsImlzcyI6InBmLWF1dGgyIiwiaXNzdWVfdGltZSI6IjE2OTU3MjA3NjI1NzgiLCJyb2xlcyI6IjEiLCJzY29wZXMiOiJhdXRoLGF1dGhfZWRpdG9yLGNoYXQsY2hhdF9lZGl0b3IsbWFpbCxtYWlsX2VkaXRvcixtZXNzYWdlLG1lc3NhZ2VfZWRpdG9yLHBheSxwYXlfZWRpdG9yLHJhYmJpdG1xLHJhYmJpdG1xX2FkbWluaXN0cmF0b3Isc3lzdGVtLHRyYW5zbGF0ZSx0cmFuc2xhdGVfZWRpdG9yLGNvbnNvbGUiLCJzZXNzaW9uX2lkIjoiYTU2NTVjMDEtNWM0Zi0xMWVlLWI1NTctZmUzYWE0ZjNjMzRmIiwidG9rZW5fdHlwZSI6ImFjY2VzcyIsInVzZXIiOiJzeXN0ZW06emhhbmd4aW9uZyIsInZlcnNpb24iOiIwLjEuMSJ9.Zc-FXnAUeKmd4FNKVF5Jpt2rUaxTbg9ECLb7KCq0zWNfBoF1wLkq2mP3SLuC_jeeTGLyVFqB83pOZk-s9lDPzhOXr3YNh26w9oTACvDzZnnX_4HGPvD4R8EanM2LmzvgLz602vpSoydDm7qQVxCPNU9OiNxvhC2FyGSQ-KjcOiGCxwT0mMmNp9x-mGHLO7W1EDOJ-nOfl9xqbwrYBii6qc1T0arrEz-bNJi_o0bbGQiAqw7iUCKGa-VTU0U32tYm72JqD640SigDKMICWHznrTPZhnO1KeB7vxtEUV3eIGVnSU0FgbipH2wIaXCUFtX0TYTJJsRdvTeQP4XovnFxKg")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8")
	rep, err := pushClient.Do(req)
	if err != nil {
		log.Printf("do send err:%v", err)
		return err
	}
	defer rep.Body.Close()
	buf, err := ioutil.ReadAll(rep.Body)
	if err != nil {
		log.Printf("readall err:%v", err)
		return err
	}

	log.Println(string(buf))
	wg.Done()
	return nil
}
