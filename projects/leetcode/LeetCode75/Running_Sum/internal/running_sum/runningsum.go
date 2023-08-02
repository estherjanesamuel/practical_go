package runningsum

func runningSum(nums []int) (result []int) {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		result = append(result, sum)
	}
    return
}

func runningSummary(nums []int) []int {
	output := []int{}

	for i := 0; i < len(nums); i++ {
		if i == 0 {
			output = append(output, nums[i])
		} else {
			output = append(output, nums[i] + nums[i-1])
		}
	}
	return output
}