package monotonicarray

func isMonotonic(nums []int) bool {
	count := 0
	for i := 1; i < len(nums); i++ {
		if count == 0 {
			if (nums[i] > nums[i - 1]) {
				count = 1
			}
			if (nums[i] < nums[i - 1]) {
				count = -1
			}
		} else {
			if count * ( nums[i] - nums[i - 1]) < 0 {
				return false
			}
		}
	}
	return true
}