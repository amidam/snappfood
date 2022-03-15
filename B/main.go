package main

import "fmt"

func main() {
	var arr = []int{2, 2, 5, 6, 5}
	fmt.Printf("%d\n", findExceptionNum(arr))
}

func findExceptionNum(s []int) (num int) {
	for _, v := range s {
		num ^= v
	}
	return
}
