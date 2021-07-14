package main 

import (
	"fmt"
	"encoding/json"
	"github.com/astaxie/beego/httplib"
)

type dataInfo struct{
	server_id  string
	role_id	string
	time 	string
	vip_level int 
	level 	int
	game_id 	int
	role_name string
	message string
	type int
}


func main(){
	req := httplib.Post("http://spro.bbsdev.yingxiong.com/fenghao/chat/index")

	var data dataInfo
	data.server_id = "111"
	data.role_id = "111"
	data.time = "2020-03-19 08:00:38"
	data.vip_level = 1
	data.level = 1
	data.game_id = 164
	data.role_name = "万条数据测试"
	data.message = "一二三哈哈哈哈"
	data.type = 1
	dataString, _ := json.Marshal(data)

	req.Param("game_id", 164)
	req.Param("sign", "7c1f464c8bb2f0b437764ab1e64640db")
	req.Param("time", 1584604838)
	req.Param("data", string(dataString))
	
	var dat map[string]interface{}
    req.ToJSON(&dat)

    fmt.Println(dat)
}

