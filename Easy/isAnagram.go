package main

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	var arr [26]int
	for i := 0; i < len(s); i++ {
		arr[s[i]-'a']++
		arr[t[i]-'a']--
	}
	for _, v := range arr {
		if v != 0 {
			return false
		}
	}
	return true
}

// func main() {
// 	fmt.Println(isAnagram("anagram", "nagaram")) // true
// 	fmt.Println(isAnagram("rat", "car")) // false
// }
