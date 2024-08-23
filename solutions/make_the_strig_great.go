package main

import "strings"

func makeGood(s string) string {
	for {
		idx := badNeighbor(s)
		if idx == -1 {
			break
		}
		s = s[:idx] + s[idx+2:]

	}
	return s
}

func badNeighbor(s string) int {
	for i := 0; i < len(s)-1; i++ {
		if check(string(s[i]), string(s[i+1])) {
			return i
		}
	}
	return -1
}

func check(a, b string) bool {
	if strings.ToLower(a) == a && strings.ToUpper(b) == b && strings.ToUpper(a) == b {
		return true
	}
	if strings.ToLower(b) == b && strings.ToUpper(a) == a && strings.ToUpper(b) == a {
		return true
	}
	return false
}
