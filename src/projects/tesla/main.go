package main

import "fmt"

func main() {
	Solution1(1234500)
	//fmt.Println(Solution2([]int{1, 0}))
	//fmt.Println(Solution3([]int{1, 2, 4, 5}, []int{2, 3, 5, 6}, 6))
}

func Solution1(N int) {
	var enable_print int
	enable_print = N % 10
	for N > 0 {
		if enable_print == 0 && N%10 != 0 {
			enable_print = 1
			fmt.Print(N % 10)
		} else if enable_print > 0 {
			fmt.Print(N % 10)
		}
		N = N / 10
	}
}

func Solution2(A []int) int {
	length, count := len(A), 0

	result, sum := 0, 0

	for i := 0; i < length-1; i++ {
		sum = A[i]
		if A[i] == 0 {
			result++
			count++
		}

		for j := i + 1; j < length; j++ {
			sum += A[j]
			if sum == 0 {
				result++
			}
		}
	}

	if A[length-1] == 0 {
		result++
	}

	if count == length-1 && A[length-1] == 0 {
		return -1
	}

	return result
}

func Solution3(A []int, B []int, N int) int {
	length, result := len(A), 0
	road := make(map[int]int)

	for i := 0; i < length; i++ {
		road[A[i]]++
		road[B[i]]++
	}

	for i := 0; i < length; i++ {
		if road[A[i]]+road[B[i]] > result {
			result = road[A[i]] + road[B[i]]
		}
	}

	return result - 1
}
