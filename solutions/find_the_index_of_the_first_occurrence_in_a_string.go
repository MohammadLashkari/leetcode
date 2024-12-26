package main

func strStr0(haystack string, needle string) int {
	for i := 0; i < len(haystack); i++ {
		c := 0
		for j := 0; j < len(needle); j++ {
			if j+i >= len(haystack) {
				break
			}
			if haystack[i+j] == needle[j] {
				c++
			}
		}
		if c == len(needle) {
			return i
		}
	}
	return -1
}

func strStr1(haystack string, needle string) int {
	for i := 0; i <= len(haystack)-len(needle); i++ {
		match := true
		for j := 0; j < len(needle); j++ {
			if haystack[i+j] != needle[j] {
				match = false
			}
		}
		if match {
			return i
		}
	}
	return -1
}

func strStr2(haystack string, needle string) int {
	for i := 0; i <= len(haystack)-len(needle); i++ {
		if haystack[i:i+len(needle)] == needle {
			return i
		}
	}
	return -1
}
