package stack

import (
	"fmt"
	"strings"
)

/*
The first line of the function uses a guard statement to check if the length of the input string is even.
If the length is odd, the function returns false because it means that the brackets are not properly balanced.

The next line creates an empty stack of characters to store the opening brackets.
The function then loops through each character item in the input string s.
For each character item, the code checks whether it is an opening bracket: (, [, or {.
If it is, the corresponding closing bracket is pushed onto the stack.
For example, if the character is (, the character ) is pushed onto the stack.
If the character is a closing bracket: ), ], or },
the code checks whether the stack is empty. If the stack is empty,
it means that there is no opening bracket to match the closing bracket,
so the function returns false.
If the stack is not empty, the last opening bracket is removed from the stack
and checked against the current closing bracket. If the brackets do not match,
the function returns false.
After looping through all the characters in the input string,
the function checks whether the stack is empty.
If it is, it means that all the opening brackets have been matched with their corresponding closing brackets, so the function returns true. If the stack is not empty, it means that there are unmatched opening brackets, so the function returns false.
*/
func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}

	openDict := make(map[byte]int, 0)
	closeDict := make(map[byte]int, 0)
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			_, exist := openDict[s[i]]
			if exist {
				openDict[s[i]]++
			} else {
				openDict[s[i]] = 1
			}
		} else if s[i] == ')' || s[i] == ']' || s[i] == '}' {
			_, exist := closeDict[s[i]]
			if exist {
				closeDict[s[i]]++
			} else {
				closeDict[s[i]] = 1
			}
		} else {
			return false
		}

	}

	for key, val := range closeDict {
		if key == ')' {
			if openDict['('] != val {
				return false
			}
		} else if key == ']' {
			if openDict['['] != val {
				return false
			}
		} else if key == '}' {
			if openDict['{'] != val {
				return false
			}
		} else {
			return false
		}
	}

	return true
}

func isValid2(s string) bool {
	// s = strings.Replace(s, "{}", "",10)
	// s = strings.Replace(s, "()", "",10)
	// s = strings.Replace(s, "[]", "",10)

	// s = strings.Replace(s, "()", "",10)
	// s = strings.Replace(s, "[]", "",10)
	// s = strings.Replace(s, "{}", "",10)

	// s = strings.Replace(s, "[]", "",10)
	// s = strings.Replace(s, "{}", "",10)
	// s = strings.Replace(s, "()", "",10)

	// s = strings.Replace(s, "{}", "",10)
	// s = strings.Replace(s, "[]", "",10)
	// s = strings.Replace(s, "()", "",10)
	if len(s)%2 != 0 {
		return false
	}

	for i := 0; i < len(s); i++ {
		s = strings.Replace(s, "{}", "", 10)
		s = strings.Replace(s, "[]", "", 10)
		s = strings.Replace(s, "()", "", 10)
	}

	fmt.Println(s)
	return s == ""
}

// ch := s[i]
// if _, exist := dict[ch]; exist {
// 	dict[ch]++
// } else {
// 	dict[ch] = 1
// }

/* Java Solutions
		public boolean isValid(String s) {
	Stack<Character> stack = new Stack<Character>();
	for (char c : s.toCharArray()) {
		if (c == '(')
			stack.push(')');
		else if (c == '{')
			stack.push('}');
		else if (c == '[')
			stack.push(']');
		else if (stack.isEmpty() || stack.pop() != c)
			return false;
	}
	return stack.isEmpty();
}
*/

/*
	To solve the problem we need to match pairs of parentheses taking into account

	- the number of parentheses,
	- their direction (open or closed) and
	- their order.

	To be able to do it we have to reverse the order of opening parenthesis temporarily to match them to their closing peers. A data structure that would able to "reverse" the parentheses' order satisfying this requirement is stack.

	We traverse the string from left to right and check:

	If we see an opening parenthesis we put it to stack
	If it’s the closing one, then it must correspond to the parenthesis at the top of our stack, so we check if it is true and we remove this pair of parentheses if it's so. If not, we return false and stop the function’s execution immediately.
	At the end, iff we have empty stack, we have valid string thus return true.
*/
func Is_Valid(s string) bool {
	pairs := map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
	}

	stack := []rune{}

	for _, r := range s {
		if _, ok := pairs[r]; ok {
			stack = append(stack, r)
		} else if len(stack) == 0 || pairs[stack[len(stack)-1]] != r {
			return false
		} else {
			stack = stack[:len(stack)-1]
		}
		fmt.Println(string(stack))
	}
	return len(stack) == 0
}

/*
	checks if the given string s contains valid brackets
	according to the specified rules.
	It uses a stack to keep track of the opening brackets encountered.

	returns true if the string is valid and false otherwise.
*/
func IS_valid(s string) bool {
	// defines the corresponding opening brackets for each closing bracket.
	pairs := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	stack := []rune{}

	for _, char := range s {
		opening, ok := pairs[char]
		fmt.Println("char: ", string(char), " | ok: ", ok, " | stack: ", string(stack),stack)
		if ok { // closing bracket encountered

			// if len(stack) == 0 || stack[len(stack)-1] != opening {
			// 	return false
			// }
			// stack = stack[:len(stack)-1] // pop from stack

			fmt.Println("opening: ", string(opening), string(stack[:len(stack)-1]),stack[:len(stack)-1], len(stack), string(stack[0]), stack[:0], stack[len(stack)-1] != opening)

		} else { // opening bracket encountered
			stack = append(stack, char)
		}
	}
	// if stack is empty, all brackets are mathce
	return len(stack) == 0
}
