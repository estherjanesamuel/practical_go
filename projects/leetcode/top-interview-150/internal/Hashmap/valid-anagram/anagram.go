package validanagram


func isValidAnagram(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	charCount := make(map[byte]int, 0)
	for i := 0; i < len(s); i++ {
		_, exist := charCount[s[i]]
		if exist {
			charCount[s[i]]++
		} else {
			charCount[s[i]] = 1
		}
	}

	for i := 0; i < len(t); i++ {
		_, exist := charCount[t[i]]
		if exist {
			charCount[t[i]]--
		}
	}

	for _, val := range charCount {
		if val > 0 {
			return false
		}
	}
	return true
}