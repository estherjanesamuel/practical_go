package Hashmap

import (
	"fmt"
)

func isIsomorphic(s, t string) bool {
	if len(s) != len(t) {
		return false
	}
	s_to_t := make(map[string]string) //make(map[byte]byte)
	t_to_s := make(map[string]string) //make(map[byte]byte)
	for i := 0; i < len(s); i++ {
		/*
				_, sExist := s_to_t[s[i]]
				_, tExist := t_to_s[t[i]]
				if sExist && s_to_t[s[i]] != t[i]{
					return false
				}

				if tExist && t_to_s[t[i]] != s[i]{
					return false
				}
			s_to_t[s[i]] = t[i]
			t_to_s[t[i]] = s[i]
		*/
		s_to_t[string(s[i])] = string(t[i])
		t_to_s[string(t[i])] = string(s[i])

		fmt.Println(s_to_t, " | ", t_to_s)

	}
	fmt.Println()
	return true
}

func Isomorphic(s, t string) bool {
	if len(s) != len(t) {
		return false
	}
	s_to_t := make(map[byte]byte)
	t_to_s := make(map[byte]byte)

	for i := 0; i < len(s); i++ {
		vals, sExist := s_to_t[s[i]]
		valt, tExist := t_to_s[t[i]]
		if sExist && vals != t[i] {
			return false
		}

		if tExist &&  valt != s[i] {
			return false
		}
		s_to_t[s[i]] = t[i]
		t_to_s[t[i]] = s[i]
	}
	return true
}
