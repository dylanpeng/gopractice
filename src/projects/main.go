package main

import (
	"crypto/md5"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	str := "lang.ar.limit.10.offset.0.waFFopay@2021"
	sign := fmt.Sprintf("%x", md5.Sum([]byte(str)))
	fmt.Println(sign)

	for i := 0; i < 100; i++ {
		probabilityMap := map[int64]int{1: 20, 2: 80}
		winId := WinningLottery(probabilityMap)
		fmt.Printf("winid: %d\n", winId)
	}
}

func WinningLottery(probabilityMap map[int64]int) (winId int64) {
	rand.Seed(time.Now().UnixNano())
	var total, n int

	for _, v := range probabilityMap {
		total += v
	}

	rNum := rand.Intn(total)

	for k, v := range probabilityMap {
		if rNum >= n && rNum < n+v {
			return k
		}
		n += v
	}

	return
}
