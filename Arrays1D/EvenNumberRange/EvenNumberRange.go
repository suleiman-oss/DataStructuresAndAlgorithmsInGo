package main

import "fmt"

func evenNum(a []int, q [][]int) []int {
	p := make([]int, len(a))
	if a[0]%2 == 0 {
		p[0] = 1
	} else {
		p[0] = 0
	}
	for i, v := range a {
		if i == 0 {
			continue
		}
		if v%2 == 0 {
			p[i] = p[i-1] + 1
		} else {
			p[i] = p[i-1]
		}
	}
	ans := make([]int, len(q))
	for i, v := range q {
		if v[0] == 0 {
			ans[i] = p[v[1]]
		} else {
			ans[i] = p[v[1]] - p[v[0]-1]
		}
	}
	return ans
}

func main() {
	a := []int{2, 4, 3, 7, 9, 8, 6, 3, 4, 9}
	q := [][]int{{4, 8}, {3, 9}, {0, 3}, {0, 2}}
	fmt.Println(evenNum(a, q))
}
