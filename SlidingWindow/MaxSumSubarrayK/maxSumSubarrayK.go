package main

import "fmt"

func maxSumSubarrayK(nums []int, k int) int {
	s, e := 0, k
	ans := 0
	for i := s; i < e; i++ {
		ans += nums[i]
	}
	l := ans
	for e < len(nums) {
		l = l - nums[s] + nums[e]
		ans = max(ans, l)
		s++
		e++
	}
	return ans
}

func main() {
	nums := []int{-3, 4, -2, 5, 3, -2, 8, 2, -1, 4}
	k := 5
	fmt.Println(maxSumSubarrayK(nums, k))
}
