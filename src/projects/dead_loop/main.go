package main

import (
	"fmt"
	"time"
)

func main() {
	i := 1
	for {
		fmt.Printf("%d, %s\n", i, time.Now())
		i++
		time.Sleep(1 * time.Minute)
	}
}
