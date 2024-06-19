package main

import (
	"fmt"
	"math"
)

func increasingTriplet(nums []int) bool {
	f := math.MaxInt
	s := math.MaxInt
	for _, v := range nums {
		if v <= f {
			f = v
		} else if v <= s {
			s = v
		} else {
			return true
		}
	}
	return false
}
func main() {
	nums := []int{5, 4, 3, 2, 1}
	fmt.Println(increasingTriplet(nums))
}
