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

type createNewMergeRequestDiscussionResponse struct {
	Id             string `json:"id"`
	IndividualNote bool   `json:"individual_note"`
	Notes          []struct {
		Id         int         `json:"id"`
		Type       string      `json:"type"`
		Body       string      `json:"body"`
		Attachment interface{} `json:"attachment"`
		Author     struct {
			Id        int         `json:"id"`
			Name      string      `json:"name"`
			Username  string      `json:"username"`
			State     string      `json:"state"`
			AvatarUrl interface{} `json:"avatar_url"`
			WebUrl    string      `json:"web_url"`
		} `json:"author"`
		CreatedAt       time.Time   `json:"created_at"`
		UpdatedAt       time.Time   `json:"updated_at"`
		System          bool        `json:"system"`
		NoteableId      int         `json:"noteable_id"`
		NoteableType    string      `json:"noteable_type"`
		Resolvable      bool        `json:"resolvable"`
		Resolved        bool        `json:"resolved"`
		ResolvedBy      interface{} `json:"resolved_by"`
		Confidential    interface{} `json:"confidential"`
		NoteableIid     int         `json:"noteable_iid"`
		CommandsChanges struct {
		} `json:"commands_changes"`
	} `json:"notes"`
}

//curl --request POST --header "PRIVATE-TOKEN: <your_access_token>"\
//--form 'position[position_type]=text'\
//--form 'position[base_sha]=<use base_commit_sha from the versions response>'\
//--form 'position[head_sha]=<use head_commit_sha from the versions response>'\
//--form 'position[start_sha]=<use start_commit_sha from the versions response>'\
//--form 'position[new_path]=file.js'\
//--form 'position[old_path]=file.js'\
//--form 'position[new_line]=18'\
//--form 'body=test comment body'\
//"https://gitlab.example.com/api/v4/projects/5/merge_requests/11/discussions"

// createNewMergeRequestDiscussion 给commit评论
func createNewMergeRequestDiscussion(projectId int, iid int, noteBody, baseSha, headSha, startSha, oldFile, newFile string, oldLine, newLine string) (*createNewMergeRequestDiscussionResponse, error) {
	params := map[string]string{
		"private_token":           PrivateToken,
		"body":                    noteBody,
		"position[position_type]": "text",
		"position[base_sha]":      baseSha,
		"position[head_sha]":      headSha,
		"position[start_sha]":     startSha,
		"position[old_path]":      oldFile,
		"position[new_path]":      newFile,
	}

	//position[new_line]和position[old_line]互斥
	if oldLine != "" && newLine == "" {
		params["position[old_line]"] = oldLine
	} else if oldLine == "" && newLine != "" {
		params["position[new_line]"] = newLine
	} else {
		return nil, fmt.Errorf("oldLine and newLine can't be empty at the same time")
	}

	jsonData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	addr := gitlabAddr + fmt.Sprintf("/projects/%d/merge_requests/%d/discussions", projectId, iid)
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

	var respInfo createNewMergeRequestDiscussionResponse

	err = json.Unmarshal(body, &respInfo)
	if err != nil {
		return nil, err
	}

	return &respInfo, nil
}
