package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// 1. 启动100个任务，但是同一时间最多运行10个
	// 2. 需要调用到 doSomething
	pool := make(chan int, 10)

	i := 0
	for ; i < 100; i++ {
		pool <- i
		go func(a int) {
			doSomething(a)
			<-pool
		}(i)
	}

	// waitting for exit signal
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

func doSomething(i int) {
	time.Sleep(time.Second)
	fmt.Println(i)
}
