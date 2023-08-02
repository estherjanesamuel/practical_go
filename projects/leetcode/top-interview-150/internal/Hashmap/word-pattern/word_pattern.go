package wordpattern

import (
	"fmt"
	"strings"

	"golang.org/x/exp/slices"
)

func wordPattern(pattern, s string) bool {
	words := strings.SplitN(s, " ", 10)

	fmt.Println(words)
	if len(pattern) != len(words) {
		return false
	}

	word := make(map[byte]string)
	for i := 0; i < len(words); i++ {
		word[pattern[i]] = words[i]
	}
	diff := []string{}
	for i := 0; i < len(pattern); i++ {
		diff = append(diff, word[pattern[i]])
	}
	fmt.Println(word)
	return slices.Equal(words, diff)
}

/*
	 find if s follows the same pattern.
	 Here follow means a full match,
	 such that there is a bijection between a letter in pattern and a non-empty word in s.
	 Input: pattern = "abba", s = "dog cat cat dog"
	Output: true
*/
func hasSamePattern(pattern, s string) bool {
	words := strings.SplitN(s, " ", 10)
	if len(words) > len(pattern) {
		return false
	}

	fmt.Println(words)

	patternToWord := make(map[byte]string, 0) // to store the mapping from pattern characters to words,
	wordToPattern := make(map[string]byte, 0) // to store the mapping from words to pattern characters
	for i := 0; i < len(pattern); i++ {
		ch := pattern[i]
		_, exist := patternToWord[ch]
		if exist {
			if patternToWord[ch] != words[i] {
				return false
			}
		}

		word := words[i]
		if _, exist = wordToPattern[word]; exist {
			if wordToPattern[word] != ch {
				return false
			}
		}
		patternToWord[ch] = word
		wordToPattern[word] = ch
	}
	return true
}

/*
To determine if a string s follows the same pattern as a given pattern, you need to establish a bijection (a one-to-one mapping) between the characters in the pattern and the words in the string s. Here's a step-by-step algorithm to solve this problem:

Split the string s into a list of words. You can use the strings.Split function in Go to split the string based on spaces.

Check if the length of the pattern is equal to the number of words in s. If they are not equal, return false since the pattern cannot be followed.

Create two empty maps: patternToWord to store the mapping from pattern characters to words, and wordToPattern to store the mapping from words to pattern characters.

Iterate through the characters in the pattern using a for loop with an index variable i:

Get the current character ch from the pattern.

Check if ch exists as a key in patternToWord. If it does, compare the corresponding word stored in patternToWord[ch] with the word at index i in the list of words. If they are not the same, return false since the pattern is violated.

Check if the word at index i exists as a key in wordToPattern. If it does, compare the corresponding character stored in wordToPattern[word] with ch. If they are not the same, return false since the pattern is violated.

If both ch and the word at index i are not present in their respective maps, add them as key-value pairs to patternToWord and wordToPattern maps.

After the loop, if you reach this point, it means that the pattern is followed correctly. Return true to indicate that s follows the pattern.

Here's the algorithm implemented in Go:
*/
func samePattern(pattern, s string) bool {
	words := strings.Split(s, " ")
	// to store the mapping from pattern characters to words,
	patternToWord := make(map[byte]string)
	// to store the mapping from words to pattern characters
	wordtoPattern := make(map[string]byte)

	/*
			map[97:dog]  |  map[dog:97]
		map[97:dog 98:cat]  |  map[cat:98 dog:97]
		map[97:dog 98:cat]  |  map[cat:98 dog:97]
		map[97:fish 98:cat]  |  map[cat:98 dog:97 fish:97]

				// a b b a | dog_cat_cat_dog | true
				// map[97:dog 98:cat]  |  map[cat:98 dog:97]

				words [dog_cat_cat_dog]
						0   1	2

				2. patternToWord[a:dog b:cat]
				2. wordtoPattern[dog:a cat:b]
				3. patternToWord[a:fish b:cat]
				2. wordtoPattern[fish:a cat:b dog:a]


				patternToWord 	| 	wordtoPattern
				0. a : dog				dog : a
				1. b : cat				cat : b
				2. b : cat				cat : b
				3. a : fish				fish : a

				cek apakah exist? jika tidak tambahkan ke map.
				jika ya cek:
				- patternToWord[i] == words[i] jika ya continue, jika tidak return false
				- wordtoPattern[i] == pattern[i] jika ya continue, jika tidak return false

				| dog_cat_cat_fish | false
				// map[97:fish 98:cat]  |  map[cat:98 dog:97 fish:97]
	*/
	for i := 0; i < len(pattern); i++ {
		ch := pattern[i]
		word := words[i]
		if val, exist := patternToWord[ch]; exist {
			if val != word {
				return false
			}
		}

		if val, exist := wordtoPattern[word]; exist {
			if val != ch {
				return false
			}
		}
		patternToWord[pattern[i]] = words[i]
		wordtoPattern[words[i]] = pattern[i]
		fmt.Println(patternToWord, " | ", wordtoPattern)
	}
	return true
}
