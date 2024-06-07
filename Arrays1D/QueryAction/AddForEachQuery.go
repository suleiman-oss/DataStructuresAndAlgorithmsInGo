package main

import "fmt"

// Given an integer array such that for every i A[i] = 0
// Return the final away after performing multiple queries
// Query (i ,j,x) Add x to all elements from A[i] to A[j]
func addForEachQueryInRange(A []int, Q [][]int) {
	for _, a := range Q {
		i := a[0]
		j := a[1]
		k := a[2]
		A[i] += k
		if j+1 < len(A) {
			A[j+1] += -k
		}
	}
	for i := range A {
		if i == 0 {
			continue
		}
		A[i] += A[i-1]
	}
}

func main() {
	Q := [][]int{{2, 3, 5}, {6, 9, 3}, {3, 4, 2}}
	A := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	addForEachQueryInRange(A, Q)
	fmt.Println(A)
}
