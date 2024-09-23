package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type postCommentToCommitResponse struct {
	Note     string `json:"note"`
	Path     string `json:"path"`
	Line     int    `json:"line"`
	LineType string `json:"line_type"`
	Author   struct {
		Id        int         `json:"id"`
		Name      string      `json:"name"`
		Username  string      `json:"username"`
		State     string      `json:"state"`
		AvatarUrl interface{} `json:"avatar_url"`
		WebUrl    string      `json:"web_url"`
	} `json:"author"`
	CreatedAt time.Time `json:"created_at"`
}

// postCommentToCommit 给commit评论
func postCommentToCommit(projectId int, commitSha string, noteBody string, path string, line string, lineType string) (*postCommentToCommitResponse, error) {
	params := map[string]string{
		"private_token": PrivateToken,
		"note":          noteBody,
		"path":          path,
		"line":          line,
		"line_type":     lineType, //new表示是新增的代码，old表示是原有的代码
	}

	jsonData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	addr := gitlabAddr + fmt.Sprintf("/projects/%d/repository/commits/%s/comments", projectId, commitSha)
	req, err := http.NewRequest("POST", addr, bytes.NewBuffer(jsonData))

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	log.Println(string(body))

	var respInfo postCommentToCommitResponse

	err = json.NewDecoder(resp.Body).Decode(&respInfo)
	if err != nil {
		return nil, err
	}

	return &respInfo, nil
}
