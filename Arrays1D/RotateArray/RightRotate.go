package main

import "fmt"

func ReverseArray(s int, e int, a []int) {
	i, j := s, e
	var t int
	for i < j {
		t = a[i]
		a[i] = a[j]
		a[j] = t
		i++
		j--
	}
}

/*
3 2 9 6 5 8 k=3
6 5 8 3 2 9
*/
//https://leetcode.com/problems/rotate-array/
func RightRotate(a []int, k int) {
	k = k % len(a)
	if k == 0 {
		return
	}
	ReverseArray(0, len(a)-1, a)
	ReverseArray(0, k-1, a)
	ReverseArray(k, len(a)-1, a)
}

/*
3 2 9 6 5 8 k=2
9 6 5 8 3 2
*/
func LeftRotate(a []int, k int) {
	k = k % len(a)
	if k == 0 {
		return
	}
	ReverseArray(0, len(a)-1, a)
	ReverseArray(0, len(a)-k-1, a)
	ReverseArray(len(a)-k, len(a)-1, a)
}
func main() {
	a := []int{3, 2, 9, 6, 5, 8}
	k := 2
	LeftRotate(a, k)
	fmt.Println(a)
}
