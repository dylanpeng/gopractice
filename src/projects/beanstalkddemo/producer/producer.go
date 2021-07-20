package main

import (
	"bufio"
	"fmt"
	"github.com/kr/beanstalk"
	"os"
	"time"
)

func main() {
	//连接beanstalkd服务器
	c, err := beanstalk.Dial("tcp", "127.0.0.1:11300")

	if err != nil {
		fmt.Printf("beanstalk init failed err:%s\n", err)
		return
	}

	defer c.Close()

	fmt.Printf("beanstalkd producer begin, please enter a word. \n")

	//设置tube队列
	c.Tube.Name = "test"
	c.TubeSet.Name["test"] = true

	var value string
	for {
		inputReader := bufio.NewReader(os.Stdin)
		value, err = inputReader.ReadString('\n')

		if err != nil {
			panic(err)
		}

		//发送消息
		id, err := c.Put([]byte(value), 30, 0, 120*time.Second)

		if err != nil {
			fmt.Printf("put message failed id:%d | err: %s\n", id, err)
		} else {
			fmt.Printf("put message success id:%d\n", id)
		}
	}
}
