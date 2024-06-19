package main

import "fmt"

func OddEvenSumQueries(nums []int, q [][]int) []int {
	ans := make([]int, len(q))
	pe := make([]int, len(nums))
	po := make([]int, len(nums))
	pe[0] = nums[0]
	po[0] = 0
	for k, v := range nums {
		if k == 0 {
			continue
		}
		if k%2 == 0 {
			pe[k] = pe[k-1] + v
			po[k] = po[k-1]
		} else {
			po[k] = po[k-1] + v
			pe[k] = pe[k-1]
		}
	}
	for i := range q {
		s := q[i][0]
		e := q[i][1]
		oe := q[i][2]
		if oe == 0 {
			ans[i] = pe[e]
			if s != 0 {
				ans[i] -= pe[s-1]
			}
		} else {
			ans[i] = po[e]
			if s != 0 {
				ans[i] -= po[s-1]
			}
		}
	}
	return ans
}

func main() {
	nums := []int{4, 3, 2, 7, 6, -2, -3, 4, 2, 7}
	q := [][]int{{2, 5, 0}, {1, 8, 1}, {1, 2, 0}}
	fmt.Println(OddEvenSumQueries(nums, q))
}
