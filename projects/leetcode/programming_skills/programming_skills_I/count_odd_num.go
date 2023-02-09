package programmingskillsi

func countOdds(low int, high int) []int {
	result := []int{}
	for i := low; i <= high; i++ {
		if i%2 == 0 {
			result = append(result, low)
		}
	}

	return result
}

func countOdd(low, high int) int {
	count := 0
	mid := high - low + 1
	count = mid / 2
	if mid%2 == 0 {
		return count
	}
	if low%2 == 1 && high%2 == 1 {
		count += 1
	}
	return count
}

func countOdd1(low, high int) int {
	return (high + 1) / 2 - low / 2
}

func countOdd2(low, high int) (count int) {
	for i := low; i < high; i++ {
		if i % 2 != 0 {
			count++
		}
	}
	return count
}

func countOdds3(low int, high int) int {
    if low % 2 != 0 || high % 2 != 0{
        return (high-low)/2 + 1
    }
    return (high-low)/2
}