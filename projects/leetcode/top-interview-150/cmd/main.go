package main

import (
	"fmt"
	// stack "top-interview-150/internal/valid-parentheses"
)

func main() {
	str := "aku suka join codeid bootcamp"
	midStar := replaceMiddleCharWithStar(str)
	firstStar := replaceFirstCharWithStar(str)
	endStart := replaceEndCharWithStar(str)
	firstEndStart := replaceFirstAndEndCharWithStar(str)
	fmt.Println(str)
	fmt.Println(midStar)
	fmt.Println(firstStar)
	fmt.Println(endStart)
	fmt.Println(firstEndStart)
	validBracket := IS_valid("[({[}])]")
	fmt.Println(validBracket)
}

func IS_valid(s string) bool {
	// defines pairs of bracket
	pairs := map[rune]rune{
		')':'(',
		']':'[',
		'}':'{',
	}

	// create Stack data structure (LIFO)
	stack := []rune{}

	// iterate over s to check the characters
	for _, char := range s {
		opening, exist := pairs[char]
		if exist { // if char is close bracket

			// check whether the stack is empty or the order of bracket doesn't match e.g ([)] return false
			if len(stack) == 0 || stack[len(stack)-1] != opening {
				return false
			}
			// pop the stack
			stack = stack[:len(stack)-1]
			
		} else { // else, char is open bracket
			// store or push into the stack
			stack = append(stack, char)
		}
	}
	return len(stack) == 0
}

func replaceMiddleCharWithStar(str string) string {
	runes := []rune(str)
	for i := 1; i < len(str)-1; i++ {

		if runes[i-1] == ' ' || runes[i+1] == ' ' {
			continue
		}
		if runes[i] == ' ' {
			continue
		}

		runes[i] = '*'
	}
	return string(runes)
}

func replaceFirstCharWithStar(str string) string {
	runes := []rune(str)
	for i := 1; i < len(str)-1; i++ {
		if i-1 == 0 {
			runes[0] = '*'
		}
		if runes[i-1] == ' ' { //|| runes[i+1] == ' ' {
			runes[i] = '*'
		}
		if runes[i] == ' ' {
			continue
		}

	}
	return string(runes)
}

func replaceEndCharWithStar(str string) string {
	runes := []rune(str)
	for i := 1; i < len(str)-1; i++ {
		// fmt.Println(string(runes[i]),i,len(str))
		if i+1 == len(str)-1 {
			runes[i+1] = '*'
		}
		if runes[i+1] == ' ' { //|| runes[i+1] == ' ' {
			runes[i] = '*'
		}
		if runes[i] == ' ' {
			continue
		}

	}
	return string(runes)
}

func replaceFirstAndEndCharWithStar(str string) string {
	runes := []rune(str)
	for i := 1; i < len(str)-1; i++ {
		// fmt.Println(string(runes[i]),i,len(str))
		if i-1 == 0 {
			runes[0] = '*'
		}
		if i+1 == len(str)-1 {
			runes[i+1] = '*'
		}
		if runes[i-1] == ' ' || runes[i+1] == ' ' {
			runes[i] = '*'
		}
		if runes[i] == ' ' {
			continue
		}

	}
	return string(runes)
}

/*IS_valid
fmt.Println(s)
	// defines pairs of bracket in map/dictionary
	pairs := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	// create stack data structure
	stack := []rune{}

	for _, char := range s {
		opening, exist := pairs[char]
		fmt.Println("`", string(stack), "`", exist)
		if exist { // char is close bracket
			// if start with close bracket and stack is nil, or the order of bracket is not valid e.g ([)]. then false
			fmt.Println(string(stack[len(stack)-1]),"vs",string(opening),"!=",stack[len(stack)-1] != opening, "char: ", string(char))
			// pop from stack
			if len(stack) == 0 || stack[len(stack)-1] != opening {
				return false
			}
			// pop from stack (LIFO)
			fmt.Println("`", string(stack), "`")
			stack = stack[:len(stack)-1]

		} else { // char is open bracket
			// push/adding open bracket to stack
			stack = append(stack, char)
		}
	}
	return len(stack) == 0

*/