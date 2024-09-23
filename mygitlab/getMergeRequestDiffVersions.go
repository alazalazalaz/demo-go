package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"
)

type GetMergeRequestDiffVersionsResponse struct {
	Id             int       `json:"id"`
	HeadCommitSha  string    `json:"head_commit_sha"`
	BaseCommitSha  string    `json:"base_commit_sha"`
	StartCommitSha string    `json:"start_commit_sha"`
	CreatedAt      time.Time `json:"created_at"`
	MergeRequestId int       `json:"merge_request_id"`
	State          string    `json:"state"`
	RealSize       string    `json:"real_size"`
}

//[
//{
//"id": 246172,
//"head_commit_sha": "8aac77d9d10a4fec1f23afc8837bfca9e8852d95",
//"base_commit_sha": "aac8dff52483b18204eb914747880b8359772d41",
//"start_commit_sha": "c58125b6bb1e5a98966505c30420ea089991bd16",
//"created_at": "2024-09-19T21:47:39.513+08:00",
//"merge_request_id": 152359,
//"state": "collected",
//"real_size": "1"
//}
//]

// GetMergeRequestDiffVersions list merge request diffs
func GetMergeRequestDiffVersions(projectId, iid int) ([]GetMergeRequestDiffVersionsResponse, error) {
	params := url.Values{
		"private_token": {PrivateToken},
	}

	addr := gitlabAddr + fmt.Sprintf("/projects/%d/merge_requests/%d/versions", projectId, iid)
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

	var respInfo []GetMergeRequestDiffVersionsResponse

	err = json.Unmarshal(body, &respInfo)
	if err != nil {
		fmt.Printf("json.Unmarshal failed, err:%v\n", err)
		return nil, err
	}

	return respInfo, nil
}
