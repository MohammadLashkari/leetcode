package main

// Time Complexity  O(n)
// Space Complexity O(n)
func twoSum(nums []int, target int) []int {
	hash := make(map[int]int, len(nums))
	for idx, num := range nums {
		if idx2, ok := hash[target-num]; ok {
			return []int{idx, idx2}
		}
		hash[num] = idx

	}
	return []int{}
}
