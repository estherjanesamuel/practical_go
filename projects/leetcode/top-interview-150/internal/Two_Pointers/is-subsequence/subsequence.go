package issubsequence

/*
	Given two strings s and t,return true if s is a subsequence of t,or false otherwise.

	A subsequence of a string is a new string that is formed from the original string 
	by deleting some (can be none) of the characters without disturbing 
	the relative positions of the remaining characters. 
	(i.e., "ace" is a subsequence of "abcde" while "aec" is not).
*/
func isSubsequence(s, t string) bool {
	count := 0
	idx := 0
	for i := 0; i < len(t); i++ {
		if count == len(s) {
			return true
		}
		if s[idx] == t[i] {
			idx ++
			count ++
		}
	}
	return count == len(s)
}