package main

import (
	"bytes"
	"crypto/tls"
	"io/ioutil"
	"net/http"
	"unsafe"
)

func main() {
	convertByteToString()
}

func bubbleSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

// send a notification to apple APNS
// and fill the request with the given parameters
func sendNotificationToAppleAPNS(token string, payload []byte, sandbox bool) error {
	var err error
	var req *http.Request
	var resp *http.Response
	var client *http.Client
	var url string

	if sandbox {
		url = "https://api.development.push.apple.com/3/device/" + token
	} else {
		url = "https://api.push.apple.com/3/device/" + token
	}

	req, err = http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		return err
	}

	req.Header.Set("apns-topic", "com.abc.def")
	req.Header.Set("apns-priority", "10")
	req.Header.Set("apns-expiration", "0")

	client = &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	resp, err = client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func convertByteToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
