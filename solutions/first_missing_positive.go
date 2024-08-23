package main

import (
	"slices"
)

// Time Complexity  O(n)
// Space Complexity O(n)
func firstMissingPositive1(nums []int) int {
	hash := make(map[int]bool)
	for _, num := range nums {
		if num >= 1 && num <= len(nums) {
			hash[num] = true
		}
	}
	for i := 1; i <= len(nums); i++ {
		if !hash[i] {
			return i
		}
	}
	return len(nums) + 1
}

// Time Complexity  O(n)
// Space Complexity O(1)
func firstMissingPositive2(nums []int) int {
	if slices.Index(nums, 1) == -1 {
		return 1
	}
	for index, num := range nums {
		if num <= 0 || num > len(nums) {
			nums[index] = 1
		}
	}
	for _, num := range nums {
		index := abs(num) - 1
		if nums[index] > 0 {
			nums[index] *= -1
		}
	}
	for index, num := range nums {
		if num > 0 {
			return index + 1
		}
	}
	return len(nums) + 1
}

// Time Complexity  O(n)
// Space Complexity O(1)
func firstMissingPositive3(nums []int) int {
	i := 0
	for i < len(nums) {
		if nums[i] > len(nums) || nums[i] <= 0 {
			i++
			continue
		}
		idx := nums[i] - 1
		if nums[idx] != nums[i] {
			nums[i], nums[idx] = nums[idx], nums[i]
		} else {
			i++
		}
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return len(nums) + 1
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
