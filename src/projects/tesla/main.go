package main

import "fmt"

func main() {
	//Solution1(1234500)
	fmt.Println(Solution2([]int{0, 0, 1, 0, -1}))
	//fmt.Println(Solution3([]int{1, 2, 4, 5}, []int{2, 3, 5, 6}, 6))
	//t := t1()
	//fmt.Println(*t)

}

func t1() (a *int) {
	a = new(int)
	*a = 9
	defer func() {
		*a++
		fmt.Printf("one: %d\n", *a)
	}()

	defer func() {
		*a++
		fmt.Printf("two: %d\n", *a)
	}()

	return a
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

	for start := 0; start < length; start++ {
		sum = 0
		if A[start] == 0 {
			count++
		}
		for end := start; end >= 0; end-- {
			sum += A[end]
			if sum == 0 {
				result++
			}
		}
	}

	if count == length {
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
