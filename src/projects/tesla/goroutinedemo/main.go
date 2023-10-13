package main

import (
	"fmt"
	"sync"
)

var messageQ chan int
var puducerQ chan int
var w sync.WaitGroup

/*尝试用goroutine和channel实现生产者-消费者业务，15个生产者，10个消费者，生产消费100次，产生0-99数字，打印出来，要求不重复*/
func main() {
	messageQ = make(chan int, 10)
	puducerQ = make(chan int, 10)

	w = sync.WaitGroup{}
	w.Add(15)
	for j := 0; j < 10; j++ {
		go consumer()
	}

	for m := 0; m < 15; m++ {
		go puducer()
	}

	for m := 0; m <= 99; m++ {
		puduce(m)
	}

	close(messageQ)
	close(puducerQ)

}

func puduce(m int) {
	puducerQ <- m
}

func puducer() {
	for q := range puducerQ {
		messageQ <- q
	}

	w.Done()
}

func consumer() {
	for m := range messageQ {
		fmt.Printf("mesage : %d\n", m)
	}

	w.Done()
}
