package isomorphicstrings

import "fmt"


func isIsomorphic(s, t string) bool {
	if len(s) != len(t){
		return false
	}

	mapping := make(map[byte]byte)
	usedChars := make(map[byte]bool)

	for i := 0; i < len(s); i++ {
		if _, ok := mapping[s[i]]; ok {
			if mapping[s[i]] != t[i] {
				return false
			}
		} else {
			if usedChars[t[i]] {
				return false
			}
			mapping[s[i]] = t[i]
			usedChars[t[i]] = true
		}
	}

	return true
}

func is_isomorphic(s,t string) bool {
	if len(s) != len(t){
		return false
	}
	mapping := make(map[byte]byte)
	mapBefore := make(map[byte]bool)
	// var equivalent bool
	eq := true

	for i := 0; i <len(s); i++ {
		// ok: key is already exist
		fmt.Println(mapping)
		fmt.Println(mapBefore)
		fmt.Printf("s[%v]: %v | t[%v]: %v \n", i,s[i],i,t[i])
		_, exist := mapping[s[i]];
		fmt.Println(exist ," ", i)
		if exist {
			fmt.Printf("if block \n ====================================\n s[%v]: %v | t[%v]: %v \n", i,s[i],i,t[i])
			fmt.Printf("mapping: %v | i: %v %v %v \n",mapping, i,mapping[s[i]], t[i])
			fmt.Printf("%v != %v => %v \n",mapping[s[i]], t[i] ,mapping[s[i]] != t[i])
			if mapping[s[i]] != t[i] {
				eq = false
			}
		} else {
			fmt.Printf("else block \n ====================================\n s[%v]: %v | t[%v]: %v \n", i,s[i],i,t[i])
			
			fmt.Printf("mapBefore[t[i]] => %v \n", mapBefore[t[i]])
			if mapBefore[t[i]] {
				eq = false
			}
			mapping[s[i]] = t[i]
			mapBefore[t[i]] = true
		}
	}
	fmt.Println(eq)
	return eq
}

func Is_Isomorphic(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	dict := make(map[byte]byte)
	mapBefore := make(map[byte]bool)

	for i := 0; i < len(s); i++ {
		_, exist := dict[s[i]]
		if !exist && !mapBefore[t[i]] {
			dict[s[i]] = t[i]
			mapBefore[t[i]] = true
		} else if dict[s[i]] != t[i] {
			return false
		}

	}
	return true
}