package main

import (
	"bufio"
	"fmt"
	"github.com/Shopify/sarama"
	"os"
	"strings"
)

func main() {
	fmt.Printf("kafka producer begin, please enter a word. \n")

	config := sarama.NewConfig()
	// 异步生产者不建议把 Errors 和 Successes 都开启，一般开启 Errors 就行
	// 同步生产者就必须都开启，因为会同步返回发送成功或者失败
	config.Producer.Return.Errors = true    // 设定是否需要返回错误信息
	config.Producer.Return.Successes = true // 设定是否需要返回成功信息
	// ack等级
	config.Producer.RequiredAcks = sarama.WaitForAll
	// 分区选择器
	config.Producer.Partitioner = sarama.NewRandomPartitioner

	producer, err := sarama.NewSyncProducer([]string{"localhost:9092"}, config)

	if err != nil {
		panic(err)
	}
	defer producer.Close()

	msg := &sarama.ProducerMessage{
		Topic: "testgo",
		//相同的key会进入相同的分区中
		//Key: sarama.StringEncoder("key"),
	}

	var value string
	//var i int32
	for {
		inputReader := bufio.NewReader(os.Stdin)
		value, err = inputReader.ReadString('\n')

		if err != nil {
			panic(err)
		}

		value = strings.Replace(value, "\n", "", -1)
		msg.Value = sarama.ByteEncoder(value)
		msg.Partition = 0
		partition, offset, err := producer.SendMessage(msg)

		if err != nil {
			fmt.Println("Send Message Fail!")
		}

		fmt.Printf("Partion = %d, offset = %d\n", partition, offset)

	}

}
