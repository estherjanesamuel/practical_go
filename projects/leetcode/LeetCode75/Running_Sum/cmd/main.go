package main

import (
	"fmt"
)

func main()  {
	mapping := make(map[byte]byte)
	usedChars := make(map[byte]bool)
	s := "add"
	t := "egg"
	ok := true
	for i := 0; i < len(s); i++ {
		fmt.Println(mapping[s[i]])
		if _, yes := mapping[s[i]]; yes {
			if mapping[s[i]] != t[i] {
				ok = false
				break
			}
		} else {
			if usedChars[t[i]] {
				ok = false
				break
			}
			mapping[s[i]] = t[i]
			usedChars[t[i]] = true
		}
	}
	if ok {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}
}