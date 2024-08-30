package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	cardInit := "9,42,24,39,54,"
	cardArray := []int{9, 42, 24, 39, 54}
	for _, v := range cardArray {
		fmt.Printf("%d,%d|", v%15, v/15)
		fmt.Println()
	}
	cardEnd := CardSorting(cardInit)
	fmt.Printf("%s", cardEnd)
}

type CardSlice []string

func (a CardSlice) Len() int {
	return len(a)
}

func (a CardSlice) Swap(i, j int) { // 重写 Swap() 方法
	a[i], a[j] = a[j], a[i]
}

// 数字从大到小，花色黑红梅芳
func (a CardSlice) Less(i, j int) bool { // 重写 Less() 方法， 从大到小排序
	iInt, _ := strconv.Atoi(a[i])
	jInt, _ := strconv.Atoi(a[j])
	iNum := iInt % 15
	iType := iInt / 15
	jNum := jInt % 15
	jType := jInt / 15
	if iNum > jNum {
		return true
	} else if iNum < jNum {
		return false
	} else {
		if iType < jType {
			return true
		} else {
			return false
		}
	}
}

func CardSorting(cards string) string {
	if cards == "" {
		return ""
	}
	cardArray := strings.Split(strings.Trim(cards, ","), ",")
	sort.Sort(CardSlice(cardArray))
	return strings.Join(cardArray, ",")
}
