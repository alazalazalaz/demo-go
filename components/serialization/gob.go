package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"github.com/go-redis/redis"
	"log"
)

func main(){

	var data interface{}
	redisClient := GetRedis()
	key := "mail_brief_n1_1_72_1"
	val, err := redisClient.Get(key).Result()
	if err != nil {
		if err != redis.Nil {
			log.Println("redis error")
		}
	}

	data =	[]byte(val)

	var u UserCache
	msg := data.([]byte)
	Decode(&msg, &u)
	log.Printf("decode success: u : \n%v", u)
	log.Printf("len(mails):%d", len(u.Mails))
}

// 用户cache
type UserCache struct {
	MaxId int64             `json:"max_id"` // 最大的邮件id
	Mails []*DbMailBrief `json:"mails"`  // 邮件列表
}

type DbMailBrief struct {
	Id               int64        `json:"id" gorm:"column:id;type:bigint(20) unsigned auto_increment;not null;primary_key"`                     // 自增ID
	GameCd           string       `json:"game_cd" gorm:"column:game_cd;type:varchar(255) default \"\";not null"`                                // BI用于表示游戏的唯一ID
	GameId           string       `json:"gameid" gorm:"column:game_id;type:varchar(255) default \"\";not null;unique_index:idx_user_mail"`      // 游戏在平台服务中的代号, 业务id
	RegionId         string       `json:"region_id" gorm:"column:region_id;type:varchar(255) default \"\";not null;unique_index:idx_user_mail"` // 大区id
	ActualType       int8         `json:"actual_type" gorm:"column:actual_type;type:tinyint default 0;not null;unique_index:idx_user_mail"`     // 发送者的ID
	ActualID         int64        `json:"actual_id" gorm:"column:actual_id;type:bigint(20) default 0;not null;unique_index:idx_user_mail"`      // 接受者ID
	MailId           int64        `json:"mail_id" gorm:"column:mail_id;type:bigint(20) default 0;not null;unique_index:idx_user_mail"`          // 邮件唯一ID
	ReceiverBIUserId string       `json:"receiver_bi_user_id" gorm:"column:receiver_bi_user_id;type:varchar(255) default \"\";not null"`        // BI使用的接收者userId,实际就是RegionId+Id
	SenderID         int64        `json:"sender_id" gorm:"column:sender_id;type:bigint(20) default 0;not null"`                                 // 发送者的ID
	SenderType       int8         `json:"sender_type" gorm:"column:sender_type;type:tinyint default 0;not null"`                                // 发送者的类型
	SenderBIUserId   string       `json:"sender_bi_user_id" gorm:"column:sender_bi_user_id;type:varchar(255) default \"\";not null"`            // BI使用的发送者userId,实际就是RegionId+Id
	Head             Header `json:"head" gorm:"column:head;type:text;not null"`                                                           // 邮件头部信息
	ContId           int64        `json:"cont_id" gorm:"column:cont_id;type:bigint(20) default 0;not null"`                                     // 内容id
	Flags            int16        `json:"flags" gorm:"column:flags;type:int(11) default 0;not null"`                                            // 邮件状态(已读,加星,分享,领取附件)
	BoxId            string       `json:"box_id" gorm:"column:box_id;type:varchar(255) default \"\";not null"`                                  // 邮件箱
	Type             int8         `json:"type" gorm:"column:type;type:tinyint default 0;not null"`                                              // 邮件类型, 1玩家,2联盟,3地图
	GroupId          string       `json:"group_id" gorm:"column:group_id;type:varchar(255) default \"\";not null"`                              // 分组ID, 可选，2表示收件箱, 3表示发件箱, 4表示收藏夹, 5表示系统报告
	Relation         int64        `json:"relation" gorm:"column:relation;type:bigint(20) default 0;not null"`                                   // 关系
	CreatTime        int64        `json:"creat_time" gorm:"column:creat_time;type:bigint(20) default 0;not null"`                               // 创建时间
	ExpireAfter      int64        `json:"expire_after" gorm:"column:expire_after;type:bigint(20) default 0;not null;index"`                     // 过期时间戳(用于在内存中使用)
	MessageId        string       `json:"message_id" gorm:"column:message_id;type:varchar(255) character set utf8 default \"\";not null;index"` // 通过sqs队列的邮件会有一个mesageid用于处理重复消息（亚马逊sqs不保证消息只能获取一次，需要自己去重）
	HasAttachment    int16        `json:"has_attachment" gorm:"column:has_attachment;type:int(11) default 9;not null"`                          // 邮件是否有摘要, 1没有,2有,9未设置
	MailContentId    int64        `json:"mail_content_id" gorm:"column:mail_content_id;type:bigint(20) default 0;not null"`                     // 关联的邮件内容ID
	DumpStatu        int8         `gorm:"-"`
}

type Header struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Subject string `json:"subject"`
}

func Decode(raw *[]byte, v interface{}) bool {

	enc := gob.NewDecoder(bytes.NewReader(*raw))
	err := enc.Decode(v)

	if err != nil {
		log.Fatalln("Decode error")
		return false
	}

	return true
}


func GetRedis() *redis.Client{
	c := redis.NewClient(setOptions())

	if _, err := c.Ping().Result(); err != nil{
		return nil
	}

	return c
}

func setOptions() *redis.Options{
	return &redis.Options{
		Addr:     fmt.Sprintf("%s:%d", "redis-dev.pf.tap4fun.com", 6379),
		Password: "", // no password set
		DB:       1819,  // use default DB
	}
}
