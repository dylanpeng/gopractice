package main

import (
	"fmt"
	"gopractice/lib/ticker"
	"time"
)

func main() {
	//beginTickerDemo()
	beginWorksetDemo()
}

type workStruct struct {
}

func (w *workStruct) DoWork() {
	fmt.Printf("workStruct do work. time: %s \n", time.Now())
}

type workStruct2 struct {
}

func (w *workStruct2) DoWork() {
	fmt.Printf("workStruct2 do work. time: %s \n", time.Now())
}

type workStruct3 struct {
}

func (w *workStruct3) DoWork() {
	fmt.Printf("workStruct3 do work. time: %s \n", time.Now())
}

func beginTickerDemo() {
	t := ticker.NewTicker(time.Second*5, &workStruct{})
	t.Start()

	time.Sleep(time.Second * 30)

	t.Stop()
}

func beginWorksetDemo() {
	workset := &ticker.WorkSet{}
	workset.AddWork(time.Second*2, &workStruct2{})
	workset.AddWork(time.Second*3, &workStruct3{})
	workset.Start()

	time.Sleep(time.Second * 30)

	workset.Stop()
}
