package main

import "log"

const (
	gitlabAddr   = "https://git.tap4fun.com/api/v4"
	PrivateToken = "xx"
)

type MergeRequestChange struct{}

func main() {
	projectId := 2932
	iid := 9 //https://git.tap4fun.com/zhangxiong/zx_kadmin/-/merge_requests/9
	log.Printf("projectId: %d, iid: %d\n", projectId, iid)
	//Version()
	//ListMergeRequestChanges(projectId, iid)
	//ListAllMergeRequestNotes(projectId, iid)
	//CreateNewMergeRequestNote(projectId, iid, "test note from demo-go")

	//测试给commit评论 begin
	//commitSha := "4e1fff49c24719643119ac9bb62310e89313efda" //https://git.tap4fun.com/zhangxiong/zx_kadmin/-/commit/4e1fff49c24719643119ac9bb62310e89313efda
	//commitPath := "Dockerfile"
	//commitLine := "9"
	//commitLineType := "new"
	//postCommentToCommit(projectId, commitSha, "test commit note from demo-go", commitPath, commitLine, commitLineType)
	//测试给commit评论 end

	//测试创建discussion begin
	//iid = 10 //https://git.tap4fun.com/zhangxiong/zx_kadmin/-/merge_requests/10
	//
	//versions, err := GetMergeRequestDiffVersions(projectId, iid)
	//if err != nil || len(versions) <= 0 {
	//	log.Printf("GetMergeRequestDiffVersions failed: %v\n", err)
	//	return
	//}
	//mergeOldPath := "newfile.go"
	//mergeNewPath := "newfile.go"
	//mergeOldLine := ""
	//mergeNewLine := "7"
	//createNewMergeRequestDiscussion(projectId, iid, "createNewMergeRequestDiscussion mergeNewLine=7", versions[0].BaseCommitSha,
	//	versions[0].HeadCommitSha, versions[0].StartCommitSha, mergeOldPath, mergeNewPath, mergeOldLine, mergeNewLine)
	//测试创建discussion end

	postCommentToMR(projectId, "postCommentToMR 18", "newfile.go", "18", "new")
}
