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

type MergeRequestChanges struct {
	Id             int         `json:"id"`
	Iid            int         `json:"iid"`
	ProjectId      int         `json:"project_id"`
	Title          string      `json:"title"`
	Description    string      `json:"description"`
	State          string      `json:"state"`
	CreatedAt      time.Time   `json:"created_at"`
	UpdatedAt      time.Time   `json:"updated_at"`
	MergedBy       interface{} `json:"merged_by"`
	MergedAt       interface{} `json:"merged_at"`
	ClosedBy       interface{} `json:"closed_by"`
	ClosedAt       interface{} `json:"closed_at"`
	TargetBranch   string      `json:"target_branch"`
	SourceBranch   string      `json:"source_branch"`
	UserNotesCount int         `json:"user_notes_count"`
	Upvotes        int         `json:"upvotes"`
	Downvotes      int         `json:"downvotes"`
	Assignee       interface{} `json:"assignee"`
	Author         struct {
		Id        int         `json:"id"`
		Name      string      `json:"name"`
		Username  string      `json:"username"`
		State     string      `json:"state"`
		AvatarUrl interface{} `json:"avatar_url"`
		WebUrl    string      `json:"web_url"`
	} `json:"author"`
	Assignees                 []interface{} `json:"assignees"`
	SourceProjectId           int           `json:"source_project_id"`
	TargetProjectId           int           `json:"target_project_id"`
	Labels                    []interface{} `json:"labels"`
	WorkInProgress            bool          `json:"work_in_progress"`
	Milestone                 interface{}   `json:"milestone"`
	MergeWhenPipelineSucceeds bool          `json:"merge_when_pipeline_succeeds"`
	MergeStatus               string        `json:"merge_status"`
	Sha                       string        `json:"sha"`
	MergeCommitSha            interface{}   `json:"merge_commit_sha"`
	SquashCommitSha           interface{}   `json:"squash_commit_sha"`
	DiscussionLocked          interface{}   `json:"discussion_locked"`
	ShouldRemoveSourceBranch  interface{}   `json:"should_remove_source_branch"`
	ForceRemoveSourceBranch   bool          `json:"force_remove_source_branch"`
	Reference                 string        `json:"reference"`
	References                struct {
		Short    string `json:"short"`
		Relative string `json:"relative"`
		Full     string `json:"full"`
	} `json:"references"`
	WebUrl    string `json:"web_url"`
	TimeStats struct {
		TimeEstimate        int         `json:"time_estimate"`
		TotalTimeSpent      int         `json:"total_time_spent"`
		HumanTimeEstimate   interface{} `json:"human_time_estimate"`
		HumanTotalTimeSpent interface{} `json:"human_total_time_spent"`
	} `json:"time_stats"`
	Squash               bool `json:"squash"`
	TaskCompletionStatus struct {
		Count          int `json:"count"`
		CompletedCount int `json:"completed_count"`
	} `json:"task_completion_status"`
	HasConflicts                bool        `json:"has_conflicts"`
	BlockingDiscussionsResolved bool        `json:"blocking_discussions_resolved"`
	Subscribed                  bool        `json:"subscribed"`
	ChangesCount                string      `json:"changes_count"`
	HeadPipeline                interface{} `json:"head_pipeline"`
	DiffRefs                    struct {
		BaseSha  string `json:"base_sha"`
		HeadSha  string `json:"head_sha"`
		StartSha string `json:"start_sha"`
	} `json:"diff_refs"`
	MergeError interface{} `json:"merge_error"`
	User       struct {
		CanMerge bool `json:"can_merge"`
	} `json:"user"`
	Changes []struct {
		OldPath     string `json:"old_path"`
		NewPath     string `json:"new_path"`
		AMode       string `json:"a_mode"`
		BMode       string `json:"b_mode"`
		NewFile     bool   `json:"new_file"`
		RenamedFile bool   `json:"renamed_file"`
		DeletedFile bool   `json:"deleted_file"`
		Diff        string `json:"diff"`
	} `json:"changes"`
}

// ListMergeRequestChanges list merge request diffs
func ListMergeRequestChanges(projectId, iid int) ([]MergeRequestChange, error) {
	params := url.Values{
		"private_token": {PrivateToken},
	}

	addr := gitlabAddr + fmt.Sprintf("/projects/%d/merge_requests/%d/changes", projectId, iid)
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
		Changes []MergeRequestChange
	}

	err = json.NewDecoder(resp.Body).Decode(&respInfo)
	if err != nil {
		return nil, err
	}

	return respInfo.Changes, nil
}
