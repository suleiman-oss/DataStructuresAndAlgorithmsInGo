package main

import "fmt"

func sumOfAllSubArrays(nums []int) int {
	ans := 0
	for i, v := range nums {
		ans = ans + v*(i+1)*(len(nums)-i)
	}
	return ans
}

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(sumOfAllSubArrays(nums))
}
