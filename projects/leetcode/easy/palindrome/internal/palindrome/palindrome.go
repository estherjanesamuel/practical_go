package palindrome

import (
	"strconv"
)

func Palindrome(n int) bool {
	s := strconv.Itoa(n)
	r := Reverse(s)
	return s == r
}

func Reverse(s string) (result string) {
	for _,v := range s {
	  result = string(v) + result
	}
	return 
}

func IsPalindrome(x int) bool{
	str := strconv.Itoa(x)
    for i := 0; i < len(str); i++ {
        j := len(str)-1-i
        if str[i] != str[j] {
            return false   
        }
    }
    return true
}

