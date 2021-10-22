package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

func main(){

	users := []int{12874004,12992464,13017873,12928199,12059472,13017873,12841418,12205331,12205331,12839046,12579412,11059014,12688920,13112912,13112912,13114010,12992464,11766766,12042803,12118805,12846645,13017873,12594620,12948786,12857178,12205331,12360169,12973652,12059472,12205331,12052695,12115512,11766766,12059472,12958289,11995767,11766766,13114063,13070107,12042803,12992464,13070107,11037325,12954879,12968411,12992464,12954879,13022703,12957834,12992464,12928199,12596954,12839046,13059151,12973652,12957834,12968411,11972249,12053432,13017873,12992464,11670081,12839046,11766766,12688920,508362,12846645,12992464,12992464,13112912,11037325,12591541,13119713,13119737,11578078,12561623,12401309,11975157,12027457,12975636,11983914,13119737,11455051,10582765,11455051,508313,10079967,11039391,13059151,11818368,12059472,10649536,11455051,11795951,12329771,12817772,12591541,10582765,10678763,12029012,11818368,12118867,12528084,12966348,12528084,10264435,13080997,10477717,12618104,12802262,10882343,12618104,11642138,10280689,11567313,11059014,11078455,12305135,11455051,10678763,13090928,12368273,10678763,12528084,10264435,12305135,10252951,10942925,12234715,10942925,10942925,10655644,11818368,10655644,12428866,11818368,10641380,12454405}
	rightUsers := RemoveRepeatedElement(users)
	for _, v := range rightUsers {
		delete(v)
	}

	log.Printf("total num : %d\r\n", totalNum)
}

var totalNum int = 0

func delete(userId int){
	fileName := "del_archive_record.log"
	f, err1 := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0666)
	if err1 != nil {
		log.Fatalf("open file err : %v", err1)
	}
	defer f.Close()

	searchUrl := "https://cproxyapi.tgs.tap4fun.com/archive/history"
	characterId := strconv.Itoa(userId)
	pageNum := 1
	pageSize := 100

	for  {

		q := url.Values{}
		q.Set("game_id", "b2")
		q.Set("server_id", "0")
		q.Set("character_id", characterId)
		q.Set("room_id", "")
		q.Set("page_size", strconv.Itoa(pageSize))
		q.Set("page_number", strconv.Itoa(pageNum))
		q.Set("start_time", "1634054400000")// 10-13 00:00:00
		q.Set("end_time", "1634659200000")// 10-20 00:00:00
		q.Set("client_id", "b2:tgs.1.0.0")

		fullUrl := searchUrl + "?" + q.Encode()
		log.Println(fullUrl)
		req, err := http.NewRequest(http.MethodGet, fullUrl, nil)
		if err != nil {
			log.Fatalf("new err:%v", err)
		}

		req.Header.Set("authorization", "Bearer eyJhbGciOiJSUzI1NiIsImtpZCI6IlJIR1VZWFRMVVYiLCJrdHkiOiJSU0EiLCJ0eXAiOiJKV1QifQ.eyJhY2NvdW50X2lkIjoiMTk3MDMyNDgzNjk3NDYzNiIsImF1ZCI6InBmLWN1c3RvbWVyIiwiZXhwIjoxNjM0OTAyMTUwNzk1LCJnYW1lX2lkIjoiYjIiLCJpcF9zZWdtZW50IjoiMC4wLjAuMC8wLDo6LzAiLCJpc3MiOiJwZi1hdXRoMiIsImlzc3VlX3RpbWUiOiIxNjM0NjQyOTQ5NDYyIiwic2NvcGVzIjoiYXV0aCxhdXRoX2VkaXRvcixhdXRoX293bmVyLGNoYXQsY2hhdF9lZGl0b3IsY2hhdF9vd25lcixtYWlsLG1haWxfZWRpdG9yLG1haWxfb3duZXIsbWVzc2FnZSxtZXNzYWdlX2VkaXRvcixtZXNzYWdlX293bmVyLHBheSxwYXlfZWRpdG9yLHBheV9vd25lcixwYXlfd3hwYXlfc2lnbixyYWJiaXRtcSxyYWJiaXRtcV9hZG1pbmlzdHJhdG9yLHN5c3RlbSx0Z3MsdHJhbnNsYXRlIiwidXNlciI6InN5c3RlbTp6aGFuZ3hpb25nIiwidmVyc2lvbiI6IjAuMS4wIn0.gkxiTem-FJ-qJ9wUPz6L4WhKt777rlQNQN85EVNchUxfNk9wUo3JJX2PRaVuQjz7O_6kEB4VH2hPexbupMGwohDYZpKFJ6w6h8OPl_P_vSGhSKod3N5cR2dYpgSFWeo1UyIsaTA69RjeitAZ0vQcv26z2TEZPvyZgMzUTzeg8CmtfEYVsMw4aGvRFb1ODYoJtaj9LeW5mFQTzGFqmR1W25-pZk7VqpK4bskLkmCWKfoRS8aePRgzJprqU_dKElRcrFfWmXoQuHJiXczEDPc5cMqIZM9DTqQJQlNjnfBw3AGkKZpSGyfUcjirDSiIUlZlEa-ulhiEUdzmoY2B6Qa6aA")
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			log.Fatalf("do err:%v", err)
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalf("readall err:%v", err)
		}

		//log.Println(string(body))
		data := &ArchiveHistoryGet{}
		if err := json.Unmarshal(body, data); err != nil {
			log.Fatalf("unmarshal err:%v", err)
		}

		currentPageCount := len(data.Histories)
		if currentPageCount <= 0 {
			log.Printf("user:%d over\r\n", userId)
			return
		}

		log.Printf("currentPageCount:%d", currentPageCount)
		log.Printf("totalCount:%d", totalNum)

		for _, v := range data.Histories {
			logString := fmt.Sprintf("id:%d,from_jid:%s, body:%s\r\n", v.Id, v.From, v.Body)
			n, err := io.WriteString(f, logString)
			if err != nil {
				log.Fatalf("write file err : %v", err)
			}

			log.Printf("success write %d bytes", n)
			totalNum++
		}

		pageNum++
	}
}

type ArchiveHistoryGet struct {
	TotalCount int32                   `json:"total_count"`
	Histories  []ArchiveHistoryGetItem `json:"histories"`
}

type ArchiveHistoryGetItem struct {
	Id        int64  `json:"id"`
	From      string `json:"from"`
	To        string `json:"to"`
	Timestamp int64  `json:"timestamp"`
	Subject   string `json:"subject"`
	Body      string `json:"body"`
	UserData  string `json:"user_data"`
}



func RemoveRepeatedElement(arr []int) (newArr []int) {
	newArr = make([]int, 0)
	for i := 0; i < len(arr); i++ {
		repeat := false
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				repeat = true
				break
			}
		}
		if !repeat {
			newArr = append(newArr, arr[i])
		}
	}
	return
}