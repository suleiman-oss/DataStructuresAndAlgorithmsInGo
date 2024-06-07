package main

import (
	"fmt"
	"math"
)

func maxContinuousSum(arr []int) int {
	s := 0
	m := math.MinInt
	for _, i := range arr {
		s = max(i, s+i)
		m = max(s, m)
	}
	return m
}

func main() {
	arr := []int{10, -20, -12, 6, 5, -3, 8, -2}
	fmt.Println(maxContinuousSum(arr))
}

//https://leetcode.com/problems/maximum-subarray/description/
