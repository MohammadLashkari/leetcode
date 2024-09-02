package main

// hello
// ll

func strStr(haystack string, needle string) int {
	n := len(needle)
	h := len(haystack)
	p := 0
	if n > h {
		return -1
	}
	for i := 0; i < h; i++ {
		for j := 0; j < n; j++ {
			p++
			break
		}
		if p == n {
			return i
		}
	}
	return -1
}

// "mississippi"
// "issip"
