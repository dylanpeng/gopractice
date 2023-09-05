package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"math/rand"
	"regexp"
	"strings"
	"time"
)

func main() {

	fmt.Println(ToValidateAddress("0x5b7f33e9f0b24465cbd575d3cb354d286a9df576"))
}

func ToValidateAddress(address string) string {
	addrLowerStr := strings.ToLower(address)
	if strings.HasPrefix(addrLowerStr, "0x") {
		addrLowerStr = addrLowerStr[2:]
		address = address[2:]
	}
	var binaryStr string
	addrBytes := []byte(addrLowerStr)
	hash256 := crypto.Keccak256Hash([]byte(addrLowerStr)) //注意，这里是直接对字符串转换成byte切片然后哈希

	for i, e := range addrLowerStr {
		//如果是数字则跳过
		if e >= '0' && e <= '9' {
			continue
		} else {
			binaryStr = fmt.Sprintf("%08b", hash256[i/2]) //注意，这里一定要填充0
			if binaryStr[4*(i%2)] == '1' {
				addrBytes[i] -= 32
			}
		}
	}

	return "0x" + string(addrBytes)
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

func oldcode() {
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

	//enCode := common.MysqlAESEncryptString("sk_abd4940222cc45fe9dc8e5279c986c6esdsk_abd4940222cc45fe9dc8e5279c986c6esdsk_abd4940222cc45fe9dc8e5279c986c6esdsdfe3rjyuiojkaet", "QPMsNI1NaYsKnevjvmurTVbfc3IDITI4")
	//
	//fmt.Printf("hexCode:%s \n", enCode)
	//
	//deCode, _ := common.MysqlAESDecryptString(enCode, "QPMsNI1NaYsKnevjvmurTVbfc3IDITI4")
	//
	//fmt.Printf("%s \n", string(deCode))

	//var err error
	//fmt.Print(err.Error())

	//now := time.Date(2022, 7, 31, 11, 0, 0, 0, time.Local)
	//pastMonth := now.AddDate(0, -1, 0)
	//
	//a, b := now.Format("0601"), pastMonth.Format("0601")
	//
	//fmt.Printf("a: %s, b: %s, a: %s, b: %s \n", a, b, now, pastMonth)

	//update2  1    333
}
