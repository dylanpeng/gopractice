package main

import (
	"context"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"os"
	"strconv"
	"time"
)

func main() {
	go StartConsumer("1")
	go StartConsumer("2")
	//go StartOrderConsumer()
	go StartProducer()
	//time.Sleep(time.Minute)
	//go StartProducer()
	time.Sleep(time.Hour)
}

func StartConsumer(num string) {
	c, _ := rocketmq.NewPushConsumer(
		consumer.WithGroupName("testGroup"),
		consumer.WithNameServer([]string{"127.0.0.1:9876"}),
		//consumer.WithInstance(num),
	)

	err := c.Subscribe("test", consumer.MessageSelector{}, func(ctx context.Context,
		msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
		for i := range msgs {
			fmt.Printf("------------------consumer:%s--------------------------\n", num)
			fmt.Printf("consumer subscribe callback: %v \n", msgs[i])
		}

		return consumer.ConsumeSuccess, nil
	})
	if err != nil {
		fmt.Println(err.Error())
	}
	// Note: start after subscribe
	err = c.Start()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}
	time.Sleep(time.Hour)
	err = c.Shutdown()
	if err != nil {
		fmt.Printf("shutdown Consumer error: %s", err.Error())
	}
}

func StartProducer() {
	p, _ := rocketmq.NewProducer(
		producer.WithNameServer([]string{"127.0.0.1:9876"}),
		producer.WithRetry(2),
		//producer.WithQueueSelector(producer.NewManualQueueSelector()),
	)
	err := p.Start()
	if err != nil {
		fmt.Printf("start producer error: %s", err.Error())
		os.Exit(1)
	}
	topic := "test"
	tags := []string{"TagA", "TagB", "TagC"}

	for i := 0; i < 20; i++ {
		tag := tags[i%3]
		if i > 10 {
			time.Sleep(time.Second * 10)
		}

		msg := &primitive.Message{
			Topic: topic,
			Body:  []byte("Hello RocketMQ Go Client! " + strconv.Itoa(i)),
			//Queue: &primitive.MessageQueue{
			//	Topic:      topic,
			//	BrokerName: "dylan-own-mac.local",
			//	QueueId:    i % 4,
			//},
		}
		msg.WithTag(tag)

		res, err := p.SendSync(context.Background(), msg)
		fmt.Printf("------------------producer--------------------------\n")

		if err != nil {
			fmt.Printf("send message error: %s\n", err)
		} else {
			fmt.Printf("producer send message success: result=%s\n", res.String())
		}
	}
	err = p.Shutdown()
	if err != nil {
		fmt.Printf("shutdown producer error: %s", err.Error())
	}
}

func StartOrderConsumer() {
	c, _ := rocketmq.NewPushConsumer(
		consumer.WithGroupName("testGroup"),
		consumer.WithNameServer([]string{"127.0.0.1:9876"}),
		consumer.WithConsumerModel(consumer.Clustering),
		consumer.WithConsumeFromWhere(consumer.ConsumeFromFirstOffset),
		consumer.WithConsumerOrder(true),
	)
	err := c.Subscribe("test", consumer.MessageSelector{Type: consumer.TAG, Expression: "TagA"}, func(ctx context.Context,
		msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
		for i := range msgs {
			fmt.Printf("------------------consumer--------------------------\n")
			fmt.Printf("consumer subscribe callback: %v \n", msgs[i])
		}
		return consumer.ConsumeSuccess, nil
	})
	if err != nil {
		fmt.Println(err.Error())
	}
	// Note: start after subscribe
	err = c.Start()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}
	time.Sleep(time.Hour)
	err = c.Shutdown()
	if err != nil {
		fmt.Printf("Shutdown Consumer error: %s", err.Error())
	}
}
