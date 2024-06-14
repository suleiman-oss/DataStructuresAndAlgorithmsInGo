package main

import "fmt"

//https://leetcode.com/problems/contiguous-array/
func MaxSubarrayWithEqual1s0s(nums []int) int {
	ans := 0

	c0, c1 := 0, 0
	m := make(map[int]int)
	m[0] = -1
	for i, v := range nums {
		if v == 0 {
			c0++
		}
		if v == 1 {
			c1++
		}
		d := c1 - c0

		if j, f := m[d]; f {
			if i-j > ans {
				ans = i - j
			}
		} else {
			m[d] = i
		}
	}
	return ans
}

func main() {
	nums := []int{0, 1, 0}
	fmt.Println(MaxSubarrayWithEqual1s0s(nums))
}
