package main

import "fmt"

//https://leetcode.com/problems/range-sum-query-mutable/
func rangeSum(nums []int, q [][]int) []int {
	for i := range nums {
		if i == 0 {
			continue
		}
		nums[i] += nums[i-1]
	}
	ans := make([]int, len(q))
	for i, n := range q {
		if n[0] != 0 {
			ans[i] = nums[n[1]] - nums[n[0]-1]
		} else {
			ans[i] = nums[n[1]]
		}
	}
	return ans
}
func main() {
	nums := []int{-3, 6, 2, 4, 5, 2, 8, -9, 3, 1}
	q := [][]int{{1, 3}, {2, 7}, {4, 8}, {0, 3}}
	fmt.Println(rangeSum(nums, q))
}
