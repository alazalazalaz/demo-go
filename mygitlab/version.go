package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

type VersionResponse struct {
	Version  string `json:"version"`
	Revision string `json:"revision"`
}

// Version {"version":"12.10.14","revision":"fe3e5d62b3e"}
func Version() ([]VersionResponse, error) {
	params := url.Values{
		"private_token": {PrivateToken},
	}

	addr := gitlabAddr + fmt.Sprintf("/version")
	addr += "?" + params.Encode()

	resp, err := http.Get(addr)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	log.Println(string(body))

	var respInfo []VersionResponse

	err = json.Unmarshal(body, &respInfo)
	if err != nil {
		fmt.Printf("json.Unmarshal failed, err:%v\n", err)
		return nil, err
	}

	return respInfo, nil
}
