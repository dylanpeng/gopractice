package main

import (
	"fmt"
	"github.com/nsqio/go-nsq"
	"log"
)

func doConsumerTask() {
	// 1. 创建消费者
	config := nsq.NewConfig()
	config.MaxInFlight = 1
	config.MaxAttempts = 3
	//config.MsgTimeout = 1 * time.Second
	q, errNewCsmr := nsq.NewConsumer("one-test", "ch-one-test", config)
	if errNewCsmr != nil {
		fmt.Printf("fail to new consumer!, topic=%s, channel=%s", "one-test", "ch-one-test")
	}

	// 2. 添加处理消息的方法
	q.AddConcurrentHandlers(nsq.HandlerFunc(func(message *nsq.Message) error {
		log.Printf("message: %v | id: %s", string(message.Body), message.ID)
		//time.Sleep(10 * time.Second)
		message.Finish()
		return nil
	}), 1)

	// 3. 通过http请求来发现nsqd生产者和配置的topic（推荐使用这种方式）
	lookupAddr := []string{
		"127.0.0.1:4161",
	}
	err := q.ConnectToNSQLookupds(lookupAddr)
	if err != nil {
		log.Panic("[ConnectToNSQLookupds] Could not find nsqd!")
	}

	// 4. 接收消费者停止通知
	<-q.StopChan

	// 5. 获取统计结果
	stats := q.Stats()
	fmt.Sprintf("message received %d, finished %d, requeued:%s, connections:%s",
		stats.MessagesReceived, stats.MessagesFinished, stats.MessagesRequeued, stats.Connections)
}

func main() {
	doConsumerTask()
}
