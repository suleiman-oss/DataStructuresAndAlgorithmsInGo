package main

import "fmt"

func NoOfLeaders(nums []int) int {
	m := nums[0]
	ans := 1
	for _, v := range nums {
		if v > m {
			ans++
			m = v
		}
	}
	return ans
}

func main() {
	nums := []int{2, 5, 3, 4, 17, 16}
	fmt.Println(NoOfLeaders(nums))
}
