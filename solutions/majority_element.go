package main

// Time Complexity  O(n)
// Space Complexity O(n)
func majorityElement(nums []int) int {
	m := make(map[int]int)
	half := len(nums) / 2
	for _, num := range nums {
		m[num]++
		if m[num] > half {
			return num
		}
	}
	return -1
}

// Boyer Moore Voting Algo
// Time Complexity  O(n)
// Space Complexity O(1)
func majorityElement2(nums []int) int {
	count, res := 0, 0
	for _, num := range nums {
		if count == 0 {
			res = num
		}
		if num == res {
			count++
		} else {
			count--
		}
	}
	return res
}
