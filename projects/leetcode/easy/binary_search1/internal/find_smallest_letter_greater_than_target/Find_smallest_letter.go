package findsmallestlettergreaterthantarget

import (
	"strings"
)

// public function
func FindSmallestLetter(letters []string, target string) string {
	return nextGreaterLetter(letters, target)
}

/*
You are given an array of characters letters that is sorted in non-decreasing order,
and a character target. There are at least two different characters in letters.

Return the smallest character in letters that is lexicographically greater than target.
If such a character does not exist, return the first character in letters.
*/
func findSmallestLetter(letters []string, target string) string {
	left, right := 0, len(letters)

	for left < right {
		if letters[0] == target {
			return letters[1]
		}
		mid := left + (right-left)/2

		/*
			Compare strings lexicographically
			Compare returns an integer comparing two strings lexicographically. The result will be 0 if a==b,
			-1 if a < b, and +1 if a > b.
			Here is a go lang example that shows how to compare strings lexicographically.

			{[]string{"c","f","j"},"c","f"},
			{[]string{"x","x","y","y"},"z","x"},
		*/
		cek := strings.Compare(letters[mid], target)
		if cek == 0 {
			return letters[mid+1]
		} else if cek == -1 {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return letters[0]
}

// example : {[]string{"c","f","j"},"d","f"},
func nextGreaterLetter(letters []string, target string) string {
	left, right := 0, len(letters)
	for left < right {
		mid := left + (right-left)/2
		if letters[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	if left == 0 {
		return letters[left]
	}

	return letters[left % len(letters)]
}

func nextGreatestLetter(letters []byte, target byte) byte {
    num:=256
    var ans byte
    for i:=0;i<len(letters);i++{
        temp:=int(letters[i]-target)
        if(temp>0&&temp<num){
            num=temp
            ans=letters[i]
        }
    }
    return ans
}

func nextGreaterLetterBytes(letters []byte, target byte) byte {
	left, right := 0, len(letters)
	for left < right {
		mid := left + (right-left)/2
		if letters[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return letters[left % len(letters)]
}

func findNextGreatesLetter(letters []byte, target byte) byte {
	n := len(letters)
	for i := 0; i < n; i++ {
		if letters[i] > target {
			return letters[i]
		}
	}
	return letters[0]
}