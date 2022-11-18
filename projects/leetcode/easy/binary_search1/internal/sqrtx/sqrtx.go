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


func mysqrt(x int) int {
	/*
	if x == 1 {
        return 1
    }
    left := 0
	right := x / 2
    var pivot int

	for left <= right {

		pivot = left + (right-left)/2
        pp := pivot*pivot

        if pp == x || ( pp < x && (pivot+1)*(pivot+1) > x ) {
            return pivot
        } else if pp > x {
			right = pivot - 1
		} else {
			left = pivot + 1
        }

	}

	return pivot
	*/
	if x == 1 {
		return 1
	}

	left, right, mid := 0, x / 2, 0
	for left <= right {
		mid = left + (right - left)/2
		sqrt := mid * mid
		if sqrt == x || (sqrt < x && (mid + 1) * (mid + 1) > x) {
			return mid
		}
		if sqrt > x {
			right = mid - 1
		}
		if sqrt < x {
			left = mid + 1
		}
	}
	return mid
}