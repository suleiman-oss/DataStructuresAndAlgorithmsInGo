package main

import "fmt"

func minSwapsK(nums []int, k int) (minSwaps int) {
	kfreq := 0
	for _, v := range nums {
		if v == k {
			kfreq++
		}
	}
	kinSubarray := 0
	for i := 0; i < kfreq; i++ {
		if nums[i] == k {
			kinSubarray++
		}
	}
	maxKinSubarray := kinSubarray
	s, e := 1, kfreq
	for e < len(nums) {
		if nums[s-1] == k {
			kinSubarray--
		}
		if nums[e] == k {
			kinSubarray++
		}
		maxKinSubarray = max(maxKinSubarray, kinSubarray)
		s++
		e++
	}
	minSwaps = kfreq - maxKinSubarray
	return
}

func main() {
	nums := []int{10, 4, 8, 7, 8, 3, 8, -1, 8, 8}
	k := 8
	fmt.Println(minSwapsK(nums, k))

}
