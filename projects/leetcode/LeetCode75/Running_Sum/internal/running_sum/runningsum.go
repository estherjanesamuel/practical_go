package runningsum

func runningSum(nums []int) (result []int) {
	
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			result = append(result, nums[i])
		} else {
			result = append(result, nums[i] + nums[i-1])
		}
	}
    return
}