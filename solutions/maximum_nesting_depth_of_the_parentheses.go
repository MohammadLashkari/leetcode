package main

// Time Complexity  O(n)
// Space Complexity O(1)
func maxDepth(s string) int {
	var open, answer int
	for _, char := range s {
		if char == '(' {
			open++
			if answer < open {
				answer = open
			}
		}
		if char == ')' {
			open--
		}
	}
	return answer
}
