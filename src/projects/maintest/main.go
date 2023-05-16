package main

import (
	"fmt"
	"time"
)

var chanA chan int

func main() {
	chanA = make(chan int)

	//go GoMethod(2)
	//go GoMethod(3)
	//go GoMethod(4)
	//go GoMethod(5)

	go func() { chanA <- 1 }()
	go func() { chanA <- 2 }()
	go func() { chanA <- 3 }()
	go func() { chanA <- 4 }()
	go func() { chanA <- 5 }()
	time.Sleep(1000 * time.Millisecond)

	go GoMethod(1)

	time.Sleep(3 * time.Second)
	close(chanA)

	time.Sleep(time.Hour)
	return
}

func GoMethod(index int) {
	//for item := range chanA {
	//	fmt.Printf("item: %d |index: %d\n", item, index)
	//}

	for {
		out := false
		select {
		case item, ok := <-chanA:
			if ok {
				fmt.Printf("item: %d |index: %d\n", item, index)
			} else {
				out = true
			}
		}

		if out {
			break
		}
	}

	fmt.Printf("index done: %d\n", index)
}
