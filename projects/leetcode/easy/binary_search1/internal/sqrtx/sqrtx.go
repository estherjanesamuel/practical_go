package sqrtx

/*
Given a non-negative integer x, return the square root of x rounded down to the nearest integer. 
The returned integer should be non-negative as well.
You must not use any built-in exponent function or operator.
For example, do not use pow(x, 0.5) in c++ or x ** 0.5 in python.
*/
func mySqrt(x int) int {
	left, right, n := 1, x, 0
	for left <= right {
		mid := left + (right - left) / 2
		s := mid * mid
		if s <= x {
			n = mid
			left = mid + 1
		} else {
			right = mid -1
		}
	}
	return n
}

func sqrt(x int) int {
	return -1
}