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

type CreateNewMergeRequestNoteResponse struct {
	Id         int         `json:"id"`
	Type       interface{} `json:"type"`
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
	Confidential    interface{} `json:"confidential"`
	NoteableIid     int         `json:"noteable_iid"`
	CommandsChanges struct {
	} `json:"commands_changes"`
}

// CreateNewMergeRequestNote create merge request note 普通评论
func CreateNewMergeRequestNote(projectId, iid int, noteBody string) (*CreateNewMergeRequestNoteResponse, error) {
	params := map[string]string{
		"private_token": PrivateToken,
		"body":          noteBody,
	}

	jsonData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	addr := gitlabAddr + fmt.Sprintf("/projects/%d/merge_requests/%d/notes", projectId, iid)
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

	var respInfo CreateNewMergeRequestNoteResponse

	err = json.NewDecoder(resp.Body).Decode(&respInfo)
	if err != nil {
		return nil, err
	}

	return &respInfo, nil
}
