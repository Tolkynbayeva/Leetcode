func nextGreatestLetter(letters []byte, target byte) byte {
	start := 0
	end := len(letters) - 1
	for start <= end {
		mid := (start + end) / 2
		if letters[mid] <= target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}

	if start == len(letters) {
		return letters[0]
	} else {
		return letters[start]
	}
}