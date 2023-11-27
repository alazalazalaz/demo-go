package main

import (
	"errors"
	"log"
	"os/exec"
	"regexp"
	"strings"
	"time"
)

func main() {
	//printCmd("")
	//printCmd("--depth=1")
	//GitClone()
	GitLog()
}

func printCmd(cloneDepthOption string) {
	commitId := "xxx"
	var gitFetchCommand *exec.Cmd
	if cloneDepthOption == "" {
		gitFetchCommand = exec.Command("git", "fetch", "--force", "origin", commitId, "/tmp/x")
	} else {
		gitFetchCommand = exec.Command("git", "fetch", cloneDepthOption, "--force", "origin", commitId, "/tmp/x")
	}
	gitFetchCommand.Dir = "/Users/zhangxiong/Downloads/aaatmp/zxkadmin_clone"
	cmdString := gitFetchCommand.String()
	log.Println(cmdString)
}

func GitClone() {
	var gitCloneCmdOut []byte
	var err error
	root := "/Users/zhangxiong/Downloads/aaatmp/zxkadmin_clone/emptydir/"
	gitCloneCmdOut, err = exec.Command("git", "clone", "--depth=3", "git@git.tap4fun.com:zhangxiong/zx_kadmin.git", root).CombinedOutput()
	if err != nil {
		log.Printf("err:%v", err)
		return
	} else {
		log.Printf("success:%v", string(gitCloneCmdOut))
		return
	}
}

func GitLog() {
	var gitLogCmdOut []byte
	var err error
	root := "/Users/zhangxiong/Downloads/aaatmp/zxkadmin_clone/emptydir/"
	gitLogCmd := exec.Command("git", "log", "--date=unix")
	gitLogCmd.Dir = root
	gitLogCmdOut, err = gitLogCmd.CombinedOutput()
	if err != nil {
		log.Printf("err:%v", err)
		return
	} else {
		log.Printf("success:%v", string(gitLogCmdOut))
	}

	//putLogToArray(gitLogCmdOut)
	getFirstDate(gitLogCmdOut)

	return
}

func getFirstDate(gitLogCmdOut []byte) (int64, error) {
	regString := `Date:[ ]*(1[\d]{9})`
	result := regexp.MustCompile(regString).FindStringSubmatch(string(gitLogCmdOut))
	if len(result) < 2 {
		return 0, errors.New("get date failed")
	}

	loc, err := time.LoadLocation("UTC")
	if err != nil {
		return 0, err
	}

	theTime, err := time.ParseInLocation("Mon Jan _2 15:04:05 2006 -0700", result[1], loc)
	if err != nil {
		return 0, err
	}
	return theTime.Unix(), nil
}

func putLogToArray(gitLogCmdOut []byte) map[string]int64 {
	logAndDate := make(map[string]int64, 0)
	logArr := strings.Split(string(gitLogCmdOut), "\n\n")
	for _, v := range logArr {
		itemsArr := strings.Split(v, "\n")
		commitId := ""
		dateString := ""
		for _, vv := range itemsArr {
			if strings.HasPrefix(vv, "commit ") {
				commitId = strings.TrimPrefix(vv, "commit ")
				continue
			}
			if strings.HasPrefix(vv, "Date: ") {
				dateString = strings.TrimPrefix(vv, "Date:   ")
				continue
			}
		}
		if commitId == "" || dateString == "" {
			continue
		}
		theDate, err := convertStringToTimestamp(dateString)
		if err != nil {
			log.Printf("err:%v", err)
			continue
		}
		logAndDate[commitId] = theDate.Unix()
	}

	return logAndDate
}

func convertStringToTimestamp(s1 string) (*time.Time, error) {
	loc, err := time.LoadLocation("UTC")
	if err != nil {
		return nil, err
	}

	theTime, err := time.ParseInLocation("Mon Jan _2 15:04:05 2006 -0700", s1, loc)
	if err != nil {
		return nil, err
	}
	return &theTime, nil
}
