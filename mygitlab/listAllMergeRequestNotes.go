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

type AllMergeRequestNotes struct {
	Id         int         `json:"id"`
	Type       *string     `json:"type"` //为nil时，是普通评论，为DiffNote时，是diff评论（也就是附有代码片段的评论）
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
	Position struct { //只有type为DiffNote时才有
		BaseSha      string      `json:"base_sha"`
		StartSha     string      `json:"start_sha"`
		HeadSha      string      `json:"head_sha"`
		OldPath      string      `json:"old_path"`
		NewPath      string      `json:"new_path"`
		PositionType string      `json:"position_type"`
		OldLine      int         `json:"old_line"`
		NewLine      interface{} `json:"new_line"`
		LineRange    interface{} `json:"line_range"`
	} `json:"position,omitempty"`
	Resolved   bool        `json:"resolved,omitempty"`
	ResolvedBy interface{} `json:"resolved_by"`
}

// ListAllMergeRequestNotes list merge request diffs
func ListAllMergeRequestNotes(projectId, iid int) ([]AllMergeRequestNotes, error) {
	params := url.Values{
		"private_token": {PrivateToken},
	}

	addr := gitlabAddr + fmt.Sprintf("/projects/%d/merge_requests/%d/notes", projectId, iid)
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

	var respInfo struct {
		Changes []AllMergeRequestNotes
	}

	err = json.NewDecoder(resp.Body).Decode(&respInfo)
	if err != nil {
		return nil, err
	}

	return respInfo.Changes, nil
}
