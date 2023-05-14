package main

import (
	"context"
	"fmt"
	"github.com/Shopify/sarama"
	"os"
	"os/signal"
	"syscall"
)

type exampleConsumerGroupHandler struct {
	Id string
}

func (exampleConsumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (exampleConsumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }
func (h exampleConsumerGroupHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		fmt.Printf("Message topic:%q partition:%d offset:%d msg: %s Id: %s\n", msg.Topic, msg.Partition, msg.Offset, string(msg.Value), h.Id)
		sess.MarkMessage(msg, "")
	}
	return nil
}

func main() {
	go StartConsumerGroup("1")
	go StartConsumerGroup("2")
	go StartConsumerGroup("3")
	go StartConsumerGroup("4")

	exit := make(chan os.Signal)
	stopSigs := []os.Signal{
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGQUIT,
		syscall.SIGABRT,
		syscall.SIGKILL,
		syscall.SIGTERM,
	}
	signal.Notify(exit, stopSigs...)

	// catch exit signal
	sign := <-exit
	fmt.Printf("stop by exit signal '%s'", sign)
}

func StartConsumerGroup(id string) {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config.Consumer.Offsets.Initial = sarama.OffsetNewest //初始从最新的offset开始

	group, err := sarama.NewConsumerGroup([]string{"localhost:9092"}, "consumer_group_1", config)
	if err != nil {
		panic(err)
	}
	defer func() { _ = group.Close() }()

	// Track errors
	go func() {
		for err := range group.Errors() {
			fmt.Println("ERROR", err)
		}
	}()

	// Iterate over consumer sessions.
	ctx := context.Background()
	for {
		topics := []string{"testgo"}
		handler := exampleConsumerGroupHandler{
			Id: id,
		}

		// `Consume` should be called inside an infinite loop, when a
		// server-side rebalance happens, the consumer session will need to be
		// recreated to get the new claims
		err := group.Consume(ctx, topics, handler)
		if err != nil {
			panic(err)
		}
	}
}
