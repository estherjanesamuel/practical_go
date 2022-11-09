package firstbadversion

/*
You are a product manager and currently leading a team to develop a new product.
Unfortunately, the latest version of your product fails the quality check.
Since each version is developed based on the previous version,
all the versions after a bad version are also bad.

Suppose you have n versions [1, 2, ..., n] and you want to find out the first bad one,
which causes all the following ones to be bad.

You are given an API bool isBadVersion(version) which returns whether version is bad.
Implement a function to find the first bad version. You should minimize the number of calls to the API.
*/
func isBadVersion(version int) bool {
	return version >= 50
}

func firstBadVersionMine(n int) int {
	left, right := 0, n
	for left <= right {
		mid := left + (right-left)/2
		if isBadVersion(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func firstBadVer(n int) int {
	i, j := 1, n
	for i+1 < j {
		mid := (i + j) / 2
		if isBadVersion(mid) {
			j = mid
		} else {
			i = mid
		}
	}
	if isBadVersion(i) {
		return i
	}
	return j
}

func firstBadVersion(n int) int {
	start := 0
	end := n

	for mid := (end + start) / 2; start < end; mid = (end + start) / 2 {
		if isBadVersion(mid) {
			end = mid
		} else {
			start = mid + 1
		}
	}

	return start
}

func firstBadVersionV2(n int) int {
	maybeBad := 1

	for mid := (n + maybeBad) / 2; maybeBad < n; mid = (n + maybeBad) / 2 {
		if isBadVersion(mid) {
			n = mid
		} else {
			maybeBad = mid + 1
		}
	}

	return n
}

func firstBadVersionV3(n int) int {
	maybeBad := 1

	for maybeBad < n {
		if isBadVersion((n + maybeBad) / 2) {
			n = (n + maybeBad) / 2
		} else {
			maybeBad = (n+maybeBad)/2 + 1
		}
	}

	return n
}

func firstBadVersionV4(n int) int {
	start := 1
	end := n

	for mid := (end + start) / 2; start < end; mid = (end + start) / 2 {
		if isBadVersion(mid) {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}

	return start
}

func firstBadVersionDiscuss1756768(n int) int {
	left, right := 1, n
	var mid int

	for right >= left {
		mid = left + (right-left)/2
		if isBadVersion(mid) {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}
