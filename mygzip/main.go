package main

import (
	"fmt"
	"log"
	"net/http"
)

// 测试gzip的压缩效率
func main() {
	testGzip()
}

// 流程：
// 1、访问同一个地址，通过控制传入不同参数，接口响应不同大小的内容，以及是否使用gzip压缩
// 2、使用client_id=n1:1中的1来控制response大小，值越大，response越大
// 3、使用request中的header Accept-Encoding:gzip来控制是否使用gzip
func testGzip() {
	batchSize := []int{10}
	//batchSize := []int{0, 10, 20, 30, 40, 70, 80, 100, 110, 160, 320, 640, 970, 1970}

	for _, v := range batchSize {
		totalSize, headerSize, bodySize := doOnce(v, false)
		outputString := fmt.Sprintf("|withGzip:false |%v |%v |%v |", totalSize, headerSize, bodySize)
		totalSizeWithGzip, headerSizeWithGzip, bodySizeWithGzip := doOnce(v, true)
		outputString += fmt.Sprintf("withGzip:true |%v |%v |%v |", totalSizeWithGzip, headerSizeWithGzip, bodySizeWithGzip)
		outputString += fmt.Sprintf("(%d-%d)/%d=%v%%|", bodySize, bodySizeWithGzip, bodySize, (bodySize-bodySizeWithGzip)*100/bodySize)
		fmt.Println(outputString)
	}
}

func doOnce(batchSize int, withGzip bool) (int, int, int) {
	myurl := fmt.Sprintf("http://127.0.0.1:32854/tgs/message/message/push/aws?client_id=n1:xx&multi_lang_data=%v&user_id=1223&message={\"MessageID\":\"64dc638a0818000a\",\"Title\":\"Training+complete\",\"Message\":\"Training+700+Shooters+II+complete\",\"Badge\":0,\"Sound\":\"\",\"URL\":\"\",\"CfgID\":22111004,\"Extension\":{\"analytics_label\":\"test\"}}&use_ntf=true", batchSize)
	httpClient := http.Client{}
	req, err := http.NewRequest("POST", myurl, nil)
	if err != nil {
		log.Fatalf("err:%v", err)
		return 0, 0, 0
	}

	req.Header.Set("Authorization", "Bearer xx")
	if withGzip {
		req.Header.Set("Accept-Encoding", "gzip")
	} else {
		req.Header.Set("Accept-Encoding", "identify")
	}
	resp, err := httpClient.Do(req)
	if err != nil {
		log.Fatalf("response err:%v", err)
		return 0, 0, 0
	}
	headerSize, bodySize, totalSize := 0, int(resp.ContentLength), 0
	for headerK, headerV := range resp.Header {
		headerSize += len(headerK)
		for _, v := range headerV {
			headerSize += len(v)
		}
	}

	totalSize = headerSize + bodySize

	defer resp.Body.Close()
	//buf, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	log.Printf("read failed:%v", err)
	//	return 0, 0, 0
	//}
	//
	//log.Println(string(buf))
	//fmt.Println(resp)
	return totalSize, headerSize, bodySize
}

//放message里
//batchNum, err := strconv.Atoi(multi_lang_data)
//if err != nil {
//pfctx.ELog("err:%v", err)
//}
//dd := "1234"
//randomList := "qazwsxedcrfvtgbyhnujmiklop"
//for i := 0; i < batchNum; i++ {
//dd += string(randomList[rand.Intn(10)])
//}
//return &typdef.MessagePushRsp{Status: true, Data: dd}, nil
