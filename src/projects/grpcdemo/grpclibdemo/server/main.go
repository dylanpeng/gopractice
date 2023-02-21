package main

import (
	"gopractice/projects/grpcdemo/grpclibdemo"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	config := &grpclibdemo.Config{
		Host: "localhost",
		Port: "23000",
	}

	router := &grpclibdemo.GRouter{}

	server := grpclibdemo.NewServer(config, router)
	_ = server.Start()

	config2 := &grpclibdemo.Config{
		Host: "localhost",
		Port: "23001",
	}

	router2 := &grpclibdemo.GRouter{}

	server2 := grpclibdemo.NewServer(config2, router2)
	_ = server2.Start()

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
	<-exit

	return
}
