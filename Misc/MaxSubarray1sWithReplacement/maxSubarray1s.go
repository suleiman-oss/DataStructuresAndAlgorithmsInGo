package main

import (
	"fmt"
	"math"
)

func maxSubarray(nums []int) int {
	l, r := 0, 0
	ans := math.MinInt
	for i, v := range nums {
		if v == 0 {
			l = 0
			for j := i - 1; j >= 0 && nums[j] == 1; j-- {
				l++
			}
			r = 0
			for j := i + 1; j < len(nums) && nums[j] == 1; j++ {
				r++
			}
			if r+l+1 > ans {
				ans = r + l + 1
			}
		}
	}
	if ans == math.MinInt {
		return len(nums)
	}
	return ans
}

func main() {
	nums := []int{0, 1, 1, 1, 0, 1, 1, 0, 1, 1, 0}
	fmt.Println(maxSubarray(nums))
}
