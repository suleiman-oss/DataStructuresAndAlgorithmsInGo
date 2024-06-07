package main

import "fmt"

// https://leetcode.com/problems/find-pivot-index/
func EqInd(a []int) int {
	p := make([]int, len(a))
	p[0] = a[0]
	for i, v := range a {
		if i == 0 {
			continue
		}
		p[i] = p[i-1] + v
	}
	//fmt.Println(p)
	for i := range p {
		s1, s2 := p[i]-a[i], p[len(p)-1]-p[i]
		if s1 == s2 {
			return i
		}
	}
	return -1
}

func EqIndEfficient(nums []int) int {
	rs := 0
	for _, v := range nums {
		rs += v
	}
	ls := 0
	for i := range nums {
		rs = rs - nums[i]
		if i != 0 {
			ls += nums[i-1]
		}
		if ls == rs {
			return i
		}
	}
	return -1
}

func main() {
	a := []int{-7, 1, 5, 2, -4, 3}
	fmt.Println(EqIndEfficient(a))
}
