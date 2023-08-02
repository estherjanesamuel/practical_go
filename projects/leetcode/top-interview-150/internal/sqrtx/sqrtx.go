package math

import "fmt"

/*
	Given a non-negative integer x,
	return the square root of x rounded down to the nearest integer.
	The returned integer should be non-negative as well.

	You must not use any built-in exponent function or operator.
	For example, do not use pow(x, 0.5) in c++ or x ** 0.5 in python.
*/
func mySqrtx(x int) int {
	// base case: if x is 0 or 1, return x
	if x < 2 {
		return x
	}

	// binary seacrh
	left := 1
	right := x / 2

	for left <= right {
		mid := left + right/2
		square := mid * mid

		// if mid squared is equal to x, return mid
		if square == x {
			return mid
		} else if square < x { // if mid squared is less than x, update the left pointer
			left = mid + 1
		} else { // if mid square is greater than x, update the right pointer
			right = mid - 1
		}
	}
	// if no perfect square is found, return the value of the right pointer
	return right

	/*
		This algorithm starts with a left pointer at 1 and a right pointer at x // 2.
		It repeatedly calculates the square of the number at the middle
		of the current range and compares it with x.
		Based on the comparison, it updates the left or right pointer accordingly
		to narrow down the search range. The algorithm terminates when it finds a perfect square or when the left pointer becomes greater than the right pointer. In the latter case, it returns the value of the right pointer, which is the square root of x rounded down to the nearest integer.0
	*/
}

func sqrt(x int) int {
	return x * x
}

func pwr(x int, y float64) int {
	for i := 1; i < int(y); i++ {
		x *= x
	}
	return x
}

func square_root(x int) int {
	// 9,3
	if x < 2 {
		return x
	}
	left := 1
	right := x / 2
	for left <= right {
		mid := left + right/2
		fmt.Println(left, right, mid, x, mid == x, mid < x)
		if mid == x {
			return mid
		}

		if mid < x {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return right
}

func sqrx(x int) int {
	/*
		2^2 = 4
		y = sqrx(4)
		y = 2
	*/

	if x < 2 {
		return x
	}

	left := 1
	right := x / 2
	for left <= right {
		mid := left + right/2
		square := mid * mid
		if square == x {
			return mid
		} else if square < x {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return right
}

func mySqrx(x int) int {
	if x == 0 {
		return x
	}

	left, right := 1, x
	fmt.Println(right)
	for left <= right {
		mid := (left + right) / 2
		if mid*mid == x {
			return mid
		}
		if mid*mid > x {
			right = mid - 1
		} else {
			left = mid + 1
		}
		fmt.Println(left, right, mid, mid*mid > x, mid*mid)
	}

	return right
}

func squareRoot(x int) int {
	if x == 0 {
		return 0
	}

	left := 1
	right := x

	for left <= right {
		mid := (left + right) / 2
		if mid * mid > x {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return right
}
