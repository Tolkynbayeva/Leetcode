package main

func lengthOfLongestSubstring(s string) int {
	start := 0
	maxLen := 0
	m := make(map[rune]int)
	for end, val := range s {
		if curStart, ok := m[val]; ok {
			if curStart >= start {
				start = m[val] + 1
			}
		}
		m[val] = end
		curLen := end - start + 1
		if curLen > maxLen {
			maxLen = curLen
		}
	}
	return maxLen
}
