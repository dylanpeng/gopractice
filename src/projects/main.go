package main

import (
	"fmt"
	"math/rand"
	"regexp"
	"strings"
	"time"
)

func main() {
	//str := "lang.ar.limit.10.offset.0.waFFopay@2021"
	//sign := fmt.Sprintf("%x", md5.Sum([]byte(str)))
	//fmt.Println(sign)
	//
	//for i := 0; i < 100; i++ {
	//	probabilityMap := map[int64]int{1: 20, 2: 80}
	//	winId := WinningLottery(probabilityMap)
	//	fmt.Printf("winid: %d\n", winId)
	//}

	//b := CheckIdFormat(1, "234345456")
	//fmt.Println(b)

	sdfj := []string{"1","2","3","4"}

	b := sdfj[1:5]


	fmt.Printf("%v \n", b)
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

func CheckIdFormat(idType int32, number string) (ok bool) {
	if idType == 1 {
		pattern := `^\d{9}$`
		reg := regexp.MustCompile(pattern)
		ok = reg.MatchString(number)
	} else if idType == 2 || idType == 3 {
		pattern := `^[A-Za-z0-9]{8}$`
		reg := regexp.MustCompile(pattern)
		ok = reg.MatchString(number)
	}

	return
}

func DesensitizationNormal(org string) (result string) {
	if len(org) == 0 {
		return
	}

	if len(org) == 1 {
		return "*"
	}

	orgLen := len(org)

	num := 4

	if orgLen <= 4 {
		num = orgLen / 2
	}

	result = strings.Repeat("*", orgLen-num) + org[orgLen-num:orgLen]
	return
}

func DesensitizationMobile(mobile string) (result string) {
	mobile = strings.TrimLeft(mobile, "+")
	result = "+" + DesensitizationNormal(mobile)
	return
}

func DesensitizationEmail(email string) (result string) {
	if len(email) == 0 {
		return
	}

	index := strings.Index(email, "@")

	if index == 0 {
		return
	}

	var preEmail, subEmail string

	if index == -1 {
		preEmail = email
		subEmail = ""
	} else {
		preEmail = email[:index]
		subEmail = email[index:]
	}

	preLen := len(preEmail)

	if preLen == 1 {
		return "*" + subEmail
	}

	if preLen <= 5 {
		result = strings.Repeat("*", 2) + preEmail[2:preLen] + subEmail
		return
	}

	result = preEmail[0:3] + strings.Repeat("*", preLen-3) + subEmail
	return
}
