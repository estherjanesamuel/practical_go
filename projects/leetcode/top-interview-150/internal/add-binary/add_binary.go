package math

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

/*
	Adds two binary strings.

	Args:
		a: The first binary string.
		b: The second binary string.

	Returns:
		The sum of two binary strings as a binary string.
*/
func addBinary(a, b string) string{
	
	fmt.Println(a)
	fmt.Println(b)

	// Convert the binary strings to list of integer
	tempA, _ := strconv.Atoi(a)
	tempB, _ := strconv.Atoi(b)


	// Reverse the lists so that the digits are in the same order
	a_list := []int{}
	b_list := []int{}

	for tempA > 0 {
		digit := tempA % 10
		a_list = append(a_list, digit)
		tempA /= 10
	}

	for tempB > 0 {
		digit := tempB % 10
		b_list = append(b_list, digit)
		tempB /= 10
	}

	// Initialize the result list
	result_list := []int{}
	carryOver := 0

	// Add the digits of the two lists, carrying over any overflows
	for i := 0; i < len(a_list); i++ {
		sumDigit := a_list[i] + b_list[i] 
		fmt.Println(a_list[i], b_list[i],sumDigit,result_list,carryOver)
		if sumDigit > 1 {
			result_list = append(result_list, sumDigit-2)
			carryOver = 1
		} else {
			result_list = append(result_list, sumDigit)
			carryOver = 0
		}

	}

	if carryOver > 0 {
		result_list = append(result_list, carryOver)
	}

	fmt.Println(result_list)

	// Reverse the lists so that the digits are in the same order
	start := 0
	end := len(result_list) -1
	for start < end {
		result_list[start], result_list[end] = result_list[end], result_list[start]
		start ++
		end --
	}
	// strings.Trim(strings.Join(strings.Split(fmt.Sprint(A), " "), delim), "[]")
	result := strings.Trim(strings.Join(strings.Split(fmt.Sprint(result_list), " "),""),"[]")
	return result
}

func addBinaryString(a, b string) string {
	// Convert the binary strings to slices of ints.
	aSlice := []int{}
	for _, c := range a {
	  aSlice = append(aSlice, int(c - '0'))
	}
	bSlice := []int{}
	for _, c := range b {
	  bSlice = append(bSlice, int(c - '0'))
	}
  
	// Reverse the slices so that the digits are in the same order.
	aSlice = reverse2(aSlice)
	bSlice = reverse2(bSlice)
  
	// Initialize the result slice.
	resultSlice := []int{}
  
	// Add the digits of the two slices, carrying over any overflows.
	carry := 0
	for i := 0; i < len(aSlice); i++ {
	  sum := aSlice[i] + bSlice[i] + carry
	  if sum > 1 {
		carry = 1
		resultSlice = append(resultSlice, sum-2)
	  } else {
		carry = 0
		resultSlice = append(resultSlice, sum)
	  }
	}
  
	// If there is any carry over from the last digit, add it to the result slice.
	if carry > 0 {
	  resultSlice = append(resultSlice, carry)
	}
  
	// Reverse the result slice so that the digits are in the correct order.
	resultSlice = reverse2(resultSlice)
  
	// Convert the result slice to a string.
	resultString := ""
	for _, digit := range resultSlice {
	  resultString += strconv.Itoa(digit)
	}
  
	// Return the result string.
	return resultString
  }
  
  func reverse2(slice []int) []int {
	for i := 0; i < len(slice)/2; i++ {
	  slice[i], slice[len(slice)-1-i] = slice[len(slice)-1-i], slice[i]
	}
	return slice
  }
  
  func addBinaryStrings(a, b string) string {
	// Convert the binary strings to slices of runes.
	aRuneSlice := []rune(a)
	bRuneSlice := []rune(b)
  
	// Reverse the slices so that the digits are in the same order.
	aRuneSlice = reverse(aRuneSlice)
	bRuneSlice = reverse(bRuneSlice)
  
	// Initialize the result slice.
	resultRuneSlice := []rune{}
  
	// Add the digits of the two slices, carrying over any overflows.
	carry := rune(0)
	for i := 0; i < len(aRuneSlice); i++ {
	  sum := aRuneSlice[i] + bRuneSlice[i] + carry
	  if sum > 1 {
		carry = 1
		resultRuneSlice = append(resultRuneSlice, rune(sum-2))
	  } else {
		carry = 0
		resultRuneSlice = append(resultRuneSlice, rune(sum))
	  }
	}
  
	// If there is any carry over from the last digit, add it to the result slice.
	if carry > 0 {
	  resultRuneSlice = append(resultRuneSlice, rune(carry))
	}
  
	// Reverse the result slice so that the digits are in the correct order.
	resultRuneSlice = reverse(resultRuneSlice)
  
	// Convert the result slice to a string.
	resultString := string(resultRuneSlice)
  
	// Return the result string.
	return resultString
  }
  
  func reverse(slice []rune) []rune {
	for i := 0; i < len(slice)/2; i++ {
	  slice[i], slice[len(slice)-1-i] = slice[len(slice)-1-i], slice[i]
	}
	return slice
  }
  

  func addBinaryStringsII(a, b string) string {
	// Convert the binary strings to slices of runes.
	aRunes := []rune(a)
	bRunes := []rune(b)
  
	// Check if the two strings are of the same length.
	if len(aRunes) != len(bRunes) {
	  // If they are not, pad the shorter string with 0s.
	  shorterRunes := aRunes
	  longerRunes := bRunes
	  if len(aRunes) < len(bRunes) {
		shorterRunes = bRunes
		longerRunes = aRunes
	  }
	  for i := 0; i < len(longerRunes)-len(shorterRunes); i++ {
		shorterRunes = append(shorterRunes, rune('0'))
	  }
	}
  
	// Initialize the result slice.
	resultRunes := []rune{}
  
	// Add the digits of the two slices, carrying over any overflows.
	carry := rune(0)
	for i := 0; i < len(aRunes); i++ {
	  sum := aRunes[i] + bRunes[i] + carry
	  if sum > 1 {
		carry = 1
		resultRunes = append(resultRunes, rune(sum-2))
	  } else {
		carry = 0
		resultRunes = append(resultRunes, rune(sum))
	  }
	}
  
	// If there is any carry over from the last digit, add it to the result slice.
	if carry > 0 {
	  resultRunes = append(resultRunes, rune(carry))
	}
  
	// Reverse the result slice so that the digits are in the correct order.
	resultRunes = reverse(resultRunes)
  
	// Convert the result slice to a string.
	resultString := string(resultRunes)
  
	// Return the result string.
	return resultString
  }
  
  func AdddBin(a, b string) string {
	  
	  i := len(a) - 1
	  j := len(b) - 1
	  ans := make([]byte, len(a)+1)	
	  var carry byte
	
	for i >= 0 || j >= 0 {
		var cA, cB byte
		if i >= 0 {
			cA = a[i]-'0'
		}
		if j >= 0 {
			cB = b[j] - '0'
		}
		sum := cA ^ cB ^ carry
		carry = cA & cB | carry & (cA ^ cB)
		ans[i+1] = sum + '0'
		i--
		j-- 

	}
	if carry == 1 {
		ans[0] = '1'
		return string(ans)
	}
	return string(ans[:])
  }
  
  func AddBin2(a,b string) string {
	aInt,bInt,sum:= new(big.Int),new(big.Int),new(big.Int)
	aInt.SetString(a,2)
	bInt.SetString(b,2)
	sum.Add(aInt,bInt)
	return sum.Text(2)
  }

  func AddBin3(a,b string) string {
	if len(a) < len(b) {
		a,b = b,a
	}
	indexB := len(b) - 1
	result := make([]byte, len(a))
	var (
		 shifter, sum byte
		)
		
	for i := len(a) - 1; i >= 0; i-- {
		sum = shifter + a[i]
		if indexB >= 0 {
			sum += b[indexB]
			indexB --
		}

		result[i] = sum % 2 + '0'
		shifter = sum >> 1 % 2
	}
	if shifter == 0 {
		return string(result)
	}

	return "1" + string(result)
  }
  