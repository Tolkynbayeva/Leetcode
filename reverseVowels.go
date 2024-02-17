func reverseVowels(s string) string {
	runes := []rune(s)
	start := 0
	end := len(runes) - 1
	for start < end {
		if !isVowel(runes[start]) && start < end {
			start++
			continue
		}
		for !isVowel(runes[end]) && start < end {
			end--
			continue
		}

		runes[start], runes[end] = runes[end], runes[start]
		start++
		end--

	}
	return string(runes)
}

func isVowel(r rune) bool {
	switch r {
	case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
		return true
	}
	return false
}

