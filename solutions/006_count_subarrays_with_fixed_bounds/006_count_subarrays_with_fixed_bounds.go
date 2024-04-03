package main

// Time Complexity  O(n)
// Space Complexity O(1)
func countSubarrays(nums []int, minK int, maxK int) int64 {
	var start, maxFreq, minFreq, currCount, answer int

	for end := 0; end < len(nums); end++ {
		if nums[end] < minK || nums[end] > maxK {
			for end < len(nums)-1 && (nums[end] < minK || nums[end] > maxK) {
				end++
			}
			start = end
			maxFreq, minFreq, currCount = 0, 0, 0
		}
		if nums[end] == minK {
			minFreq++
		}
		if nums[end] == maxK {
			maxFreq++
		}
		for minFreq >= 1 && maxFreq >= 1 {
			if nums[start] == minK {
				minFreq--
			}
			if nums[start] == maxK {
				maxFreq--
			}
			start++
			currCount++
		}
		answer += currCount
	}
	return int64(answer)
}
