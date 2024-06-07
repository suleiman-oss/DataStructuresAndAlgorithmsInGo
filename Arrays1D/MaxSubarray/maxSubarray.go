package main

import (
	"fmt"
	"math"
)

func maxSubarray(nums []int) int {
	maxEle := math.MinInt
	minEle := math.MaxInt
	for _, v := range nums {
		if v > maxEle {
			maxEle = v
		}
		if v < minEle {
			minEle = v
		}
	}
	lastMax, lastMin := -1, -1
	ans := math.MaxInt
	for i, v := range nums {
		if v == maxEle {
			lastMax = i
			if lastMin != -1 {
				ans = min(ans, lastMax-lastMin+1)
			}
		}
		if v == minEle {
			lastMin = i
			if lastMax != -1 {
				ans = min(ans, lastMin-lastMax+1)
			}
		}
	}
	return ans
}

func main() {
	nums := []int{1, 2, 3, 1, 3, 4, 6, 4, 6, 3}
	fmt.Println(maxSubarray(nums))
}
