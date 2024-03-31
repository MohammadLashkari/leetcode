package main

import (
	"slices"
)

// Time Complexity  O(n)
// Space Complexity O(1)
func countSubarrays(nums []int, k int) int64 {
	var (
		max                          = slices.Max(nums)
		max_freq, start, end, answer int
	)

	for end = 0; end < len(nums); end++ {
		if nums[end] == max {
			max_freq++
		}
		for max_freq >= k {
			if nums[start] == max {
				max_freq--
			}
			start++
		}
		answer += start
	}
	return int64(answer)
}
