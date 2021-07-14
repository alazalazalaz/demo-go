package main

import (
	"fmt"
	sarama "gopkg.in/Shopify/sarama.v1"
)

func main(){
	fmt.Println("xxx")

	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true

	//构造一个消息
	msg := &sarama.ProducerMessage{}
	msg.Topic = "test_log"
	msg.Value = sarama.StringEncoder("哈哈哈啊this is a test log")

	//连接kafka
	client, err := sarama.NewSyncProducer([]string{"localhost:9092"}, config)
	if err != nil {
		fmt.Println("连接失败, err: ", err)
		return
	}

	defer client.Close()

	//发送消息
	pid, offset, err := client.SendMessage(msg)
	if err != nil {
		fmt.Println("发送失败, err : ", err)
		return
	}

	fmt.Printf("写入成功， pid:%d, offset:%d\n", pid, offset)
}


