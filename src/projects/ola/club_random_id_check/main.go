package main

import "fmt"

func main() {
	r := CheckClubRandomID(11243213)
	fmt.Printf("check result: %v\n", r)
	//r := CheckSameNum([]int64{3, 1, 1, 1, 3, 1}, 3)
	//fmt.Printf("check result: %v\n", r)
	//r := CheckContinuousNum([]int64{7, 6, 5, 6, 7, 4, 7}, 3)
	//fmt.Printf("check result: %v\n", r)
}

func CheckClubRandomID(rid int64) bool {
	idArry := make([]int64, 0)
	idArryReverse := make([]int64, 0)
	temp := rid
	for ; temp > 0; temp = temp / 10 {
		idArry = append(idArry, temp%10)
		idArryReverse = append(idArryReverse, temp%10)
	}
	for i := 0; i < len(idArry)/2; i++ {
		idArry[i], idArry[len(idArry)-i-1] = idArry[len(idArry)-i-1], idArry[i]
	}
	// 不要出现豹子号，三个或者以上挨着相同的号码，可以是989291，不能是999821
	if CheckSameNum(idArry, 3) {
		return false
	}
	// 顺子号，3个或以上挨着连起来的号码，可以是132465，不能是123645
	if CheckContinuousNum(idArry, 3) {
		return false
	}
	// 顺子号，3个或以上挨着连起来的号码，可以是132465，不能是123645
	if CheckContinuousNum(idArryReverse, 3) {
		return false
	}
	return true
}

// 检查连续相同数字
func CheckSameNum(arry []int64, minLen int) bool {
	l := len(arry)
	if l < minLen || minLen <= 1 {
		return false
	}
	// groupNum 连续几个一样
	for groupNum := minLen; groupNum <= l; groupNum++ {
		for i := 0; i <= l-groupNum; i++ {
			firstNum := arry[i]
			allSame := true
			for j := i + 1; j <= groupNum+i-1; j++ {
				if arry[j] != firstNum {
					allSame = false
					break
				}
			}
			if allSame {
				return true
			}
		}
	}
	return false
}

// 检查连续相同数字
func CheckContinuousNum(arry []int64, minLen int) bool {
	l := len(arry)
	if l < minLen || minLen <= 1 {
		return false
	}
	// groupNum 连续几个顺子
	for groupNum := minLen; groupNum <= l; groupNum++ {
		for i := 0; i <= l-groupNum; i++ {
			allContinuous := true
			for j := i + 1; j <= groupNum+i-1; j++ {
				if arry[j] != arry[j-1]+1 {
					allContinuous = false
					break
				}
			}
			if allContinuous {
				return true
			}
		}
	}
	return false
}
