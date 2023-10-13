package main

import "fmt"

/*
打印全排列数组
*/
func main() {
	a := "1234"
	CalcALLPermutation(&a, 0, 3)

	return
}

func CalcALLPermutation(str *string, from int, to int) {
	strr := []rune(*str)
	if to <= 1 {
		return
	}
	if from == to {
		fmt.Println("最终结果：", string(strr))
	} else {
		for j := from; j <= to; j++ {
			strr[j], strr[from] = strr[from], strr[j]
			*str = string(strr)
			CalcALLPermutation(str, from+1, to)
		}
	}
}
