package main

import "fmt"

func main() {
	var nums = []int{12, 54, 89, 21, 66, 47, 14, 285, 96}

	c := make(chan int)
	go cal(nums[:len(nums)/2], c)
	go cal(nums[len(nums)/2:], c)
	x, y := <-c, <-c

	fmt.Println(x+y)
}

func cal(nums []int, c chan int) {
	sum := 0
	for _, v := range nums {
		sum += v * v
	}
	c <- sum
}
