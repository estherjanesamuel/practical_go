package maximumsubarray


func maximumSubarray(nums []int) int {
	sum := 0
	minSum := 0
	var ans int
	for _, n := range nums {
		sum += n
		
		ans = max(ans, sum-minSum)
		minSum = min(minSum, sum)
	}
	if ans == 0 {
		return minSum
	}
	return ans
}


func maximumSubarrayII(nums []int) int {
	maxSum := nums[0]
	currentSum := 0
	for _, n := range nums {
		currentSum += n
		if currentSum > maxSum {
			maxSum = currentSum	
		}
		if currentSum < 0 {
			currentSum = 0
		}
	}
	return maxSum
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}
