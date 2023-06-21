package main

import (
	"fmt"
	"gopractice/projects/gormshardingdemo/global"
	"gopractice/projects/gormshardingdemo/logic"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	global.InitDB()

	//logic.Logic.AddSharding()
	//logic.Logic.GetSharding()
	logic.Logic.Update()

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

	return
}
