package main

import "fmt"

func totalWateTrapped(height []int) int {
	lm := make([]int, len(height))
	lm[0] = height[0]
	rm := make([]int, len(height))
	rm[len(rm)-1] = height[len(height)-1]
	for i := 1; i < len(height); i++ {
		lm[i] = max(lm[i-1], height[i])
		rm[len(height)-1-i] = max(rm[len(height)-i], height[len(height)-1-i])
	}
	ans := 0
	for i := 0; i < len(height); i++ {
		ans += min(lm[i], rm[i]) - height[i]
	}
	return ans
}

func main() {
	h := []int{4, 2, 5, 7, 4, 2, 3, 6, 8, 2, 3}
	fmt.Println(totalWateTrapped(h))
}

//https://leetcode.com/problems/trapping-rain-water/description/
