package main

import (
	"log"
	"os"
	"time"
)

// TOTAL_MB 总体写入多少MB，至少1MB起步。
const TOTAL_MB = 10

// IS_DEL_FILE_AFTER_WRITE 写完后是否删除文件
const IS_DEL_FILE_AFTER_WRITE = true

//这个脚本的作用：简单的写入文件，可以大概估算出磁盘的iops。
func main(){
	fileName := "temp_test_file.log"
	//拼凑一个1MB的字符
	content := make([]byte, 0)
	for i:= 0; i< 1024; i++{
		for j:= 0; j< 1024; j++{
			content = append(content, '1')
		}
	}
	var f *os.File

	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		//文件不存在
	}else{
		os.Remove(fileName)
	}

	f = createFile(fileName)
	defer f.Close()

	log.Println("begin write...")
	totalBytes := 0
	beginAt := time.Now().UnixNano()
	for i := 0; i< TOTAL_MB; i++{
		n, err := f.Write(content)
		if err != nil {
			log.Fatalln("write error")
		}
		//log.Printf("wrote %d bytes", n)
		totalBytes += n
	}
	endAt := time.Now().UnixNano()
	durationNS := endAt - beginAt

	totalMB := totalBytes/(1024*1024)
	log.Printf("begin at : %d\n", beginAt)
	log.Printf("end at : %d\n", endAt)
	log.Printf("duration : %dns, %dms\n", durationNS, durationNS/1e6)
	log.Printf("total wrote: %dB, %dMB", totalBytes, totalMB)
	log.Printf("IOPS : %0.2fMB/s\n", (float64(totalMB)/float64(durationNS)) * 1e9)
	if IS_DEL_FILE_AFTER_WRITE {
		os.Remove(fileName)
		log.Println("deleted file")
	}else{
		log.Println("the file is not deleted,please delete it manually")
	}
	log.Println("write over")
}

func createFile(fileName string) *os.File{
	f, err := os.Create(fileName)
	if err != nil {
		log.Fatalln("create file failed")
	}
	return f
}