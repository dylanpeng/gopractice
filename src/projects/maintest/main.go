package main

import (
	"fmt"
	"time"
)

func main() {
	chanA := make(chan bool)

	go func() {
		time.Sleep(3 * time.Second)
		close(chanA)
		fmt.Printf("send\n")
	}()

	b, ok := <-chanA

	if ok {
		fmt.Printf("ok: %s\n", b)
	} else {
		fmt.Printf("not ok %s\n", b)
	}

	return
}
