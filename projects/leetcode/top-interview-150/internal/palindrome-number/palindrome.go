package math

import (
	"fmt"
	"strconv"
)

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	str := strconv.Itoa(x)
	start := 0
	end := len(str)-1
	for start < end {
		if str[start] != str[end] {
			return false
		}
		start++
		end--
	}
	return true
}

func is_palindrome(x int) bool {
	if x < 0 {
		return false
	}

	// extract the digits from the integer
	// we extract the digits from the integer x by performing modulo to get the rightmost digit,
	// adding it to the 'digits' slice and the dividing x
	// by 10 to remove the rightmost digit we repeat this process untill x becomes 0
	

	temp := x
	digs := []int{}
	for temp > 0 {
		digit := temp % 10
		fmt.Println(digit,temp)
		digs = append(digs, digit)
		temp /= 10
	}

	// fmt.Println(digs)
	start := 0
	end := len(digs) -1
	for start < end {
		if digs[start] != digs[end] {
			return false
		}
		start ++
		end --
	}
	return true
}