package validpalindrome

import (
	"fmt"
	"regexp"
	"strings"
)

/*
	A phrase is a palindrome if, after converting all uppercase letters into lowercase letters
	and removing all non-alphanumeric characters, it reads the same forward and backward.
	Alphanumeric characters include letters and numbers.

	Given a string s, return true if it is a palindrome, or false otherwise.
*/

var nonAlphanumericRegex = regexp.MustCompile(`[^a-zA-Z0-9 ]+`)

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	s = strings.Join(strings.Fields(s), "")
	s = nonAlphanumericRegex.ReplaceAllString(s, "")
	// r := reverse(s)
	// r := rvrs(s, 0, len(s)-1)
	r := reverseII(s)
	fmt.Println(s, r)

	return s == r
}

/* TODO: reverse string , byte, or rune
func reverse(str string, start, end int) string {
	if start < end {
		// str[start], str[end] = str[end], str[start]
		fmt.Println(str[start])//, " | ", str[end])
		start++
		end--
	}
	return str
}
*/

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
		// fmt.Println(i,j)
	}
	return string(runes)
}

func rvrs(str string, start, end int) string {
	nums := []rune(str)
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
	return string(nums)
}

func reverseII(s string) string {
	start, end := 0, len(s)-1
	runes := []rune(s)
	if start < end {
		runes[start], runes[end] = runes[end], runes[start]
		start++
		end++
	}
	 fmt.Println(string(runes))

	return string(runes)
}

/* https://stackoverflow.com/questions/32081808/strip-all-whitespace-from-a-string
package main_test

import (
    "strings"
    "unicode"
    "testing"
)

func SpaceMap(str string) string {
    return strings.Map(func(r rune) rune {
        if unicode.IsSpace(r) {
            return -1
        }
        return r
    }, str)
}

func SpaceFieldsJoin(str string) string {
    return strings.Join(strings.Fields(str), "")
}

func SpaceStringsBuilder(str string) string {
    var b strings.Builder
    b.Grow(len(str))
    for _, ch := range str {
        if !unicode.IsSpace(ch) {
            b.WriteRune(ch)
        }
    }
    return b.String()
}

func BenchmarkSpaceMap(b *testing.B) {
    for n := 0; n < b.N; n++ {
        SpaceMap(data)
    }
}

func BenchmarkSpaceFieldsJoin(b *testing.B) {
    for n := 0; n < b.N; n++ {
        SpaceFieldsJoin(data)
    }
}

func BenchmarkSpaceStringsBuilder(b *testing.B) {
    for n := 0; n < b.N; n++ {
        SpaceStringsBuilder(data)
    }
}

randomString := "  hello      this is a test"
fmt.Println(strings.ReplaceAll(randomString, " ", ""))

>hellothisisatest
*/

/*	https://gosamples.dev/remove-non-alphanumeric/#:~:text=ReplaceAllString()%20method%20to%20replace,)%20and%20numbers%20(%D9%A6).

package main

import (
    "fmt"
    "regexp"
)

var nonAlphanumericRegex = regexp.MustCompile(`[^a-zA-Z0-9 ]+`)

func clearString(str string) string {
    return nonAlphanumericRegex.ReplaceAllString(str, "")
}

func main() {
    str := "Test@%String#321gosamples.dev ـا ą ٦"
    fmt.Println(clearString(str))
}

*/
