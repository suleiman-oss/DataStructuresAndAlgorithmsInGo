package main

import "fmt"

func NoOfGoodTriplets(nums []int, a, b, c int) int {
	ans := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] <= b {
			l := 0
			for j := i - 1; j >= 0 && nums[j] <= a; j-- {
				l++
			}
			r := 0
			for j := i + 1; j < len(nums) && nums[j] <= c; j++ {
				r++
			}
			ans += (r * l)
		}
	}
	return ans
}

func main() {
	nums := []int{3, 0, 1, 1, 9, 7}
	a, b, c := 7, 2, 3
	fmt.Println(NoOfGoodTriplets(nums, a, b, c))
}
