package main

import "fmt"

//https://leetcode.com/problems/two-sum/description/
func TwoSum(a []int, t int) []int {
	m := make(map[int]int)
	for i, n := range a {
		if j, f := m[n]; f {
			return []int{i, j}
		}
		m[t-n] = i
	}
	return []int{-1, -1}
}

//https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/
func TwoSumSorted(a []int, t int) []int {
	i := 0
	j := len(a) - 1
	for i < j {
		if a[i]+a[j] == t {
			return []int{i, j}
		}
		if a[i]+a[j] > t {
			j--
		} else {
			i++
		}
	}
	return []int{-1, -1}
}

func main() {
	a := []int{-2, 1, 3, 4, 6, 8}
	k := 9
	fmt.Println(TwoSumSorted(a, k))
}
