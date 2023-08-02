package issubsequence

import "fmt"

func IsSubsequence(s, t string) bool {
	found := 0
	// fmt.Printf("len(t): %v  len(s) : %v, len(t) && len(s)\n", len(t), len(s) )
	for i,j := 0,0; i < len(t) && j<len(s); i++ {
		fmt.Printf("i: %v, j:%v | s[j]:%s == t[i]:%s = %v \n",i,j,string(s[j]), string(t[i]),s[j] == t[i] )
		if s[j] == t[i] {
			j++
			found++
		}
		fmt.Printf("i:%d \t j:%d \n", i,j)
	}
	return found == len(s);
}


func Is_Subsequence(s, t string) bool {
	found, j := 0,0

	for i := 0; i < len(t); i++ {
		if s[j] == t[i] {
			j++
			found++
		}
	}

	return len(s) == found
}

func is_subsequence(s,t string) bool {
	var index int
	for _index := range t {
		if index == len(s) {
			return true
		}

		if t[_index] == s[index] {
			index ++	
		}
	}
	return index == len(s)
}

func Subsequence(s,t string) bool {
	var index int
	for i := 0; i < len(t); i++ {
		if index == len(s) {
			return true
		}

		if t[i] == s[index]{
			index ++
		}
	}
	return index == len(s)
}