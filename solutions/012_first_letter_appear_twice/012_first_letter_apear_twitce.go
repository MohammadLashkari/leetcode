package main

func repeatedCharacter(s string) byte {
	m := make(map[byte]struct{})
	for i := 0; i < len(s); i++ {
		if _, ok := m[s[i]]; ok {
			return s[i]
		}
		m[s[i]] = struct{}{}
	}
	return ' '
}
