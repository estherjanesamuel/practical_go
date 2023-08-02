package Hashmap

import "fmt"



func canConstruct(ransomNote, magazine string) bool{
	if len(ransomNote) > len(magazine)  {
		return false
	}
	alphabets_counter := make([]int, 26)

	char := []rune(ransomNote)
	runes := []rune(magazine) //"aa", "aab"
	fmt.Println(alphabets_counter)

	for j := 0; j < len(magazine); j++ {
		fmt.Println(alphabets_counter[runes[j]-'a'], " | ", j)
		alphabets_counter[runes[j]-'a']++
		fmt.Println(alphabets_counter,j)

	}
	// fmt.Println(alphabets_counter)
	for i := 0; i < len(ransomNote); i++ {
		// fmt.Println(alphabets_counter,i, " | ", char[i]-'a', char[i], 'a', " | ", alphabets_counter[char[i]-'a'])
		if alphabets_counter[char[i]-'a'] == 0 {
			fmt.Println(i)
			return false
		}
		alphabets_counter[char[i]-'a']--
		// fmt.Println(alphabets_counter,i, " | ", char[i]-'a', char[i], 'a', " | ", alphabets_counter[char[i]-'a'])

	}
	return true
}

/*
return true if ransomNote can be constructed by using the letters from magazine and false otherwise.
*/
func construct(ransomNote, magazine string) bool{
	if len(ransomNote) > len(magazine){
		return false
	}
	
	ransomCount := make(map[byte]int)
	for i := 0; i < len(ransomNote); i++ {
		_, exist := ransomCount[ransomNote[i]]
		if exist  {
			ransomCount[ransomNote[i]]++
		} else {
			ransomCount[ransomNote[i]] = 1
		}
	}
	
	for i := 0; i < len(magazine); i++ {
		_, exist := ransomCount[magazine[i]]
		if exist  {
			ransomCount[magazine[i]]--
		}
	}

	for _, val := range ransomCount {
		if val > 0 {
			return false
		}
	}
	fmt.Println("ransomCount: ", ransomCount)
	return true
}