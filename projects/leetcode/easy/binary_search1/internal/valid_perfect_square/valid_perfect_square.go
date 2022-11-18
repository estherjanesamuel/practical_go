package validperfectsquare

func isPerfectSquare(num int) bool {
	left, right := 1, num/2
	for i := 0; i < num; i++ {
		mid := left + (right-left)/2
		if mid*mid == num {
			return true
		}
		if mid*mid < num {
			left = mid + 1
		}

		if mid*mid > num {
			right = mid - 1
		}
	}
	return false
}

func perfectSquare(num int) bool {
	left, right := 1, num
	for left <= right {
		// mid := (right + left) / 2
		mid := left + (right-left)/2
		sqrt := mid * mid
		if sqrt == num {
			return true
		} else if sqrt < num {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return false
}

func isPerfectSquare2(num int) bool {
	left, right := 1, num
	for left <= right {
		mid := left + (right-left)/2
		sqrt := mid * mid
		if sqrt == num {
			return true
		} else if sqrt < num {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return false
}

func isPerfectSquare3(num int) bool {
	var s int
	i := 1
	for s < num {
		s += i
		i += 2
	}

	if s == num {
		return true
	}
	return false
}

func isPerfectSquare4(num int) bool {
	left, right := 1, num
	for i := 0; i < num; i++ {
		mid := left + (right-left)/2
		sqrt := mid * mid
		if sqrt == num {
			return true
		}
		if sqrt < num {
			left = mid + 1
		}

		if sqrt > num {
			right = mid - 1
		}
	}
	return false
}
