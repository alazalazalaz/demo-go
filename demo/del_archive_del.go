package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main(){

	ids := readFromFile()

	//ids := []int{370592049,370592401,370592549,370592811}

	num := 0
	var putIds []int
	for _, id := range ids {
		num++
		putIds = append(putIds, id)
		if num == 100 {
			del(putIds)
			fmt.Printf("delTotalNum:%d\n", delTotalNum)
			num = 0
			putIds = nil
		}
	}

	del(putIds)

	fmt.Printf("delTotalNum:%d\n", delTotalNum)
}

var delTotalNum, delSuccessNum, delFailedNum = 0, 0, 0

func readFromFile() []int{
	fileName := "del_archive_record.log"
	fi, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return nil
	}
	defer fi.Close()

	br := bufio.NewReader(fi)
	var delNum []int

	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		lineString := string(a)

		rr := regexp.MustCompile(`[0-9]+,from`)
		aa := rr.FindAllString(lineString, -1)
		for _, v := range aa {
			stringArr := strings.Split(v, ",")
			if len(stringArr) > 1 {
				num, err := strconv.Atoi(stringArr[0])
				if err == nil && num > 1000000 {
					delNum = append(delNum, num)
				}
			}
		}
	}

	fmt.Println("total del : ")
	fmt.Println(len(delNum))
	return delNum
}


func del(ids []int){
	if len(ids) <= 0 {
		return
	}
	var idsString []string
	for _, v := range ids {
		idsString = append(idsString, strconv.Itoa(v))
	}
	stringIds := strings.Join(idsString, ",")

	searchUrl := "https://cproxyapi.tgs.tap4fun.com/archive/byids"

	q := url.Values{}
	q.Set("ids", stringIds)
	q.Set("client_id", "b2:tgs.1.0.0")

	fullUrl := searchUrl + "?" + q.Encode()

	log.Println(fullUrl)

	req, err := http.NewRequest(http.MethodDelete, fullUrl, nil)
	if err != nil {
		log.Fatalf("del() new err:%v", err)
	}

	req.Header.Set("authorization", "Bearer eyJhbGciOiJSUzI1NiIsImtpZCI6IlJIR1VZWFRMVVYiLCJrdHkiOiJSU0EiLCJ0eXAiOiJKV1QifQ.eyJhY2NvdW50X2lkIjoiMTk3MDMyNDgzNjk3NDYzNiIsImF1ZCI6InBmLWN1c3RvbWVyIiwiZXhwIjoxNjM0OTAyMTUwNzk1LCJnYW1lX2lkIjoiYjIiLCJpcF9zZWdtZW50IjoiMC4wLjAuMC8wLDo6LzAiLCJpc3MiOiJwZi1hdXRoMiIsImlzc3VlX3RpbWUiOiIxNjM0NjQyOTQ5NDYyIiwic2NvcGVzIjoiYXV0aCxhdXRoX2VkaXRvcixhdXRoX293bmVyLGNoYXQsY2hhdF9lZGl0b3IsY2hhdF9vd25lcixtYWlsLG1haWxfZWRpdG9yLG1haWxfb3duZXIsbWVzc2FnZSxtZXNzYWdlX2VkaXRvcixtZXNzYWdlX293bmVyLHBheSxwYXlfZWRpdG9yLHBheV9vd25lcixwYXlfd3hwYXlfc2lnbixyYWJiaXRtcSxyYWJiaXRtcV9hZG1pbmlzdHJhdG9yLHN5c3RlbSx0Z3MsdHJhbnNsYXRlIiwidXNlciI6InN5c3RlbTp6aGFuZ3hpb25nIiwidmVyc2lvbiI6IjAuMS4wIn0.gkxiTem-FJ-qJ9wUPz6L4WhKt777rlQNQN85EVNchUxfNk9wUo3JJX2PRaVuQjz7O_6kEB4VH2hPexbupMGwohDYZpKFJ6w6h8OPl_P_vSGhSKod3N5cR2dYpgSFWeo1UyIsaTA69RjeitAZ0vQcv26z2TEZPvyZgMzUTzeg8CmtfEYVsMw4aGvRFb1ODYoJtaj9LeW5mFQTzGFqmR1W25-pZk7VqpK4bskLkmCWKfoRS8aePRgzJprqU_dKElRcrFfWmXoQuHJiXczEDPc5cMqIZM9DTqQJQlNjnfBw3AGkKZpSGyfUcjirDSiIUlZlEa-ulhiEUdzmoY2B6Qa6aA")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("del() do err:%v", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("del() readall err:%v", err)
	}

	log.Println(string(body))
	data := &ArchiveDelete{}
	if err := json.Unmarshal(body, data); err != nil {
		log.Fatalf("del()  unmarshal err:%v", err)
	}

	if data.Success == true {
		delTotalNum += data.DeletedNum
		delSuccessNum += data.DeletedNum
	}
}

type ArchiveDelete struct {
	Success    bool `json:"success"`
	CountNum   int  `json:"count_num"`
	DeletedNum int  `json:"deleted_num"`
}