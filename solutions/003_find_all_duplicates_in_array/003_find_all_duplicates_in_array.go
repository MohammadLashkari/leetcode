package main

func findDuplicates1(nums []int) []int {
	freq := make(map[int]bool)
	answer := make([]int, 0, len(nums))
	for _, num := range nums {
		if freq[num] {
			answer = append(answer, num)
		} else {
			freq[num] = true
		}
	}
	return answer
}

func findDuplicates2(nums []int) []int {
	answer := make([]int, 0, len(nums))

	for _, num := range nums {
		index := abs(num) - 1
		if nums[index] < 0 {
			answer = append(answer, abs(num))
		} else {
			nums[index] *= -1
		}
	}
	return answer
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
