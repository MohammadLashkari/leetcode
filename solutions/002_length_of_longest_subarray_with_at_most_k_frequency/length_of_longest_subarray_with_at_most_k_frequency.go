package main

// Time Complexity  O(n)
// Space Complexity O(n)
func maxSubarrayLength(nums []int, k int) int {
	var (
		start, end, answer = 0, 0, 0
		freq               = make(map[int]int)
	)
	for end < len(nums) {
		freq[nums[end]]++
		for freq[nums[end]] > k {
			freq[nums[start]]--
			start++
		}
		answer = max(answer, end-start+1)
		end++
	}
	return answer

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
