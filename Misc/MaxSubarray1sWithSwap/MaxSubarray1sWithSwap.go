package main

import (
	"fmt"
	"math"
)

func MaxSubarray1sWithSwap(nums []int) int {
	t1s := 0
	for _, v := range nums {
		if v == 1 {
			t1s++
		}
	}
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
			if t1s > r+l+1 {
				if ans < r+l+1 {
					ans = r + l + 1
				}
			} else {
				if ans < r+l {
					ans = r + l
				}
			}
		}
	}
	if ans == math.MinInt {
		return len(nums)
	}
	return ans
}

func main() {
	nums := []int{1, 1, 0, 1, 0, 1, 0, 1, 1}
	fmt.Println(MaxSubarray1sWithSwap(nums))
}
