package happynumber

import (
	"fmt"
	"strconv"
)

/*
	an algorithm to determine if a number n is happy.

	A happy number is a number defined by the following process:

	Starting with any positive integer, replace the number by the sum of the squares of its digits.
	Repeat the process until the number equals 1 (where it will stay), or it loops endlessly in a cycle which does not include 1.
	Those numbers for which this process ends in 1 are happy.
	Return true if n is a happy number, and false if not.



	Example 1:

	Input: n = 19
	Output: true
	Explanation:
	1^2 + 9^2 = 82
	8^2 + 2^2 = 68
	6^2 + 8^2 = 100
	1^2 + 0^2 + 0^2 = 1
	Example 2:

	Input: n = 2
	Output: false
*/
func isHappy(n int) bool {
	inLoop := make(map[int]bool, 0)
	for _, exist := inLoop[n]; !exist; {
		squareSum, remain := 0, 0
		fmt.Println(n)
		for n > 0 {
			remain = n % 10
			squareSum += remain * remain
			n /= 10
			fmt.Println(n,remain,squareSum,inLoop)
		}
		n = squareSum
		if n == 1 {
			return true
		}
		if inLoop[n] {
			return false
		} else {
			inLoop[n] = true
		}
	}
	return true
}

func IsHappy(n int) bool {
	seen := make(map[int]bool)
	for n != 1 {
		sum := 0
		for n != 0 {
			sum += (n % 10) * (n % 10)
			n = n / 10
		}
		n = sum

		if seen[n] {
			return false
		} else {
			seen[n] = true
		}
	}
	return true
}

func areHappy(n int) bool {

	inLoop := make(map[int]bool, 0)
	for n != 1 {
		squareSum, remain := 0, 0
		for n != 0 {
			remain = n % 10
			squareSum += remain * remain
			n /= 10
		}
		n = squareSum
		if inLoop[n] {
			return false
		} else {
			inLoop[n] = true
		}
	}
	return true
}

func is_happy(n int) bool {
	// create a set/map called visited to keep track of all the numbers we have encountered 
	// while calculating the sum of squares of digits.
	visited := make(map[int]bool, 0)

	// initialize a variable "current" with the value of n.
	current := n

	// create a loop that continues until current is either 1(happy nuber)
	// or number already visited (unhappy number)
	for current != 1 || !visited[current]{
		// convert the current number to a string
		current_str := strconv.Itoa(current)

		// initialize a variable 'sum' to 0
		sum := 0

		// iterate through each digit in the string representation
		// of the current number
		for _, r := range current_str {
			// for each digit, square it and add it to the sum
			val, _ := strconv.Atoi(string(r))
			sum += (val * val)
		}

		// update the value of the current number with the sum
		current = sum

		// check if the current number is 1, 
		// if true return True since it is a happy number
		if current == 1 {
			return true
		}

		// check if the current number is already in the visited set,
		// if true return false since it is an unhappy number.
		if visited[current] {
			return false
		}

		// add the current number to visited set 
		visited[current]= true

	}

	// return false since the loop has ended without finding 1,
	// indicating the the number is an unhappy number
	return false
}
