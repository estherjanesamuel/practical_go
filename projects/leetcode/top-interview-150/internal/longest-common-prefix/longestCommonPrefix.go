package longestcommonprefix

import (
	"fmt"
	"sort"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	sort.Strings(strs)
	s1 := strs[0]
	s2 := strs[len(strs)-1]
	idx := 0
	for idx < len(s1) && idx < len(s2) {
		if s1[idx] == s2[idx] {
			idx++
		} else {
			break
		}
	}
	return substr(s1, 0, idx)
}

// NOTE: this isn't multi-Unicode-codepoint aware, like specifying skintone or
//       gender of an emoji: https://unicode.org/emoji/charts/full-emoji-modifiers.html
func substr(input string, start int, length int) string {
	asRunes := []rune(input)

	if start >= len(asRunes) {
		return ""
	}

	if start+length > len(asRunes) {
		length = len(asRunes) - start
	}

	return string(asRunes[start : start+length])
}

func longestCmonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	minLength := len(strs[0])
	for i := 1; i < len(strs); i++ {
		fmt.Println(strs[i-1], minLength)
		if len(strs[i]) > minLength {
			minLength = len(strs[i])
		}
	}
	fmt.Println(minLength)
	for i := 0; i < minLength; i++ {
		char := strs[0][i]
		fmt.Println(string(char), minLength, i)
		for j := 1; j < len(strs); j++ {
			if strs[j][i] != char {
				return strs[0][:i]
			}
		}
	}
	return strs[0][:minLength]
}

/*
	A prefix is a collection of characters at the beginning of a string.
	For instance, “mi” is a prefix of “mint” and the longest common prefix between “mint”, “mini”, and “mineral” is “min”.
	In order to find the longest common prefix, we sort the array of strings alphabetically.
	Then, we compare the characters in the first and last strings in the array.
	If the character in first is in last at the corresponding index,
	the character must be in the remaining words at the corresponding index as well,
	because the array of strings have already been sorted.
*/
func longestCommonPre(strs []string) string {
	sort.Strings(strs)
	firstStr := strs[0]
	lastStrs := strs[len(strs)-1]
	longest := ""
	for i := 0; i < len(firstStr); i++ {
		if firstStr[i] == lastStrs[i] {
			longest += string(firstStr[i])
		} else {
			break
		}
	}
	return longest
}

// 592665
