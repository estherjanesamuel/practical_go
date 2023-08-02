package firstoccurrence

import "fmt"

func firstOccurrence(haystack, needle string) int {
	/*
		need := hitung len dari needle
		cek array per by need e.g sadbutsad (array of char),
		need = 3 maka char dengan index 0 dengan len =3 adalah sad == dengan needle
		index += len cek

	*/
	idx := []int{}
	// need := len(needle)
	// fmt.Println(need, idx)
	for i := 0; i < len(haystack); i++ {
		// fmt.Println(heystack[i:], len(heystack),i)
		hay := haystack[i:]
		if len(hay) >= len(needle) {
			fmt.Println(haystack[i:])
			if hay[:len(needle)] == needle {
				idx = append(idx, i)
			}
		}

	}
	if len(idx) > 0 {
		return idx[0]
	}
	return -1
}
