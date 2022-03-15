package main

import "fmt"

func main() {
	var arr = []int{10, 10, 11, 15, 1, 11, 92, 15, 1, 92, 93}
	fmt.Printf("%d\n", findExceptionNum(arr))
}

func findExceptionNum(s []int) int {
	num := s[0]

	for i := 1; i < len(s); i++ {
		num ^= s[i]
	}

	return num
}
