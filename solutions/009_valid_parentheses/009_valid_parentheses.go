package main

// Time Complexity  O(n)
// Space Complexity O(n)
func isValid1(s string) bool {
	var (
		stack       []rune
		closeToOpen = map[rune]rune{')': '(', '}': '{', ']': '['}
	)
	for _, char := range s {
		if last, in := closeToOpen[char]; in {
			if len(stack) != 0 && stack[len(stack)-1] == last {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		} else {
			stack = append(stack, char)
		}
	}
	return len(stack) == 0
}

func isValid2(s string) bool {
	var stack []rune

	for _, char := range s {
		switch char {
		case '(':
			stack = append(stack, ')')
		case '{':
			stack = append(stack, '}')
		case '[':
			stack = append(stack, ']')
		default:
			if len(stack) == 0 || char != stack[len(stack)-1] {
				return false
			}
			stack = stack[:len(stack)-1]
		}

	}
	return len(stack) == 0
}
