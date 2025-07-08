package main

import "fmt"

const (
	minimumRandomClubID = 100000
)

func main() {
	randomID := FakeUniqueID(1000, minimumRandomClubID)
	r := CheckClubRandomID(int64(randomID))
	fmt.Printf("check result: %v, randomID: %d\n", r, randomID)
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

var fakeCvtMap map[uint64]uint64

func init() {
	fakeCvtMap = make(map[uint64]uint64)
	fakeCvtMap[1] = 9
	fakeCvtMap[2] = 6
	fakeCvtMap[3] = 3
	fakeCvtMap[4] = 8
	fakeCvtMap[5] = 4
	fakeCvtMap[6] = 5
	fakeCvtMap[7] = 2
	fakeCvtMap[8] = 0
	fakeCvtMap[9] = 1
	fakeCvtMap[0] = 7
}

func FakeUniqueID(uniqueID uint64, min ...uint64) uint64 {
	var num uint64
	if len(min) > 0 {
		uniqueID += min[0]
		num = getXORNumber(min[0])
	}
	return fakeUniqueID(uniqueID, num, 1)
}

// FakeUniqueID 转化一个唯一ID到另一钟唯一表示
func fakeUniqueID(uniqueID uint64, num uint64, l3 int) uint64 {
	l := 0
	vv := uniqueID
	r := []uint64{}
	r = append(r, vv%10)
	l++
	for vv >= 10 {
		vv /= 10
		r = append(r, vv%10)
		l++
	}
	i := 0
	for i < l/2 {
		r[i], r[l-2-i] = r[l-2-i], r[i]
		i++
	}
	var v uint64
	pow := pow10(l - 1)
	for i := 0; i < l; i++ {
		num := r[l-i-1]
		tp2 := fakeCvtMap[num]
		if i == 0 && tp2 == 0 {
			tp2 = fakeCvtMap[0]
		}
		v += tp2 * pow
		pow /= 10
	}
	l3--
	if l3 > 0 {
		return fakeUniqueID(v, num, l3)
	}
	return v ^ num
}

func pow10(l int) uint64 {
	var u uint64 = 1
	for l > 0 {
		u *= 10
		l--
	}
	return u
}

func getXORNumber(num uint64) uint64 {
	var l int
	vv := num
	l++
	for vv >= 10 {
		vv /= 10
		l++
	}
	size := l - 5
	if size > 0 {
		var n uint64
		pow := pow10(size - 1)
		for i := 0; i < size; i++ {
			ii := i%10 + 1
			n += uint64(ii) * pow
			pow /= 10
		}
		return n
	}
	return 0
}
