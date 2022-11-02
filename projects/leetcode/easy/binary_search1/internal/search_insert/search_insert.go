package searchinsert

func searchinsert(nums []int, target int) (n int) {
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			n = i
			return
		}
		if nums[i] > target {
			return 
		}
		if i == len(nums)-1 {
			n = len(nums)
			return
		}
		if nums[i] < target && target < nums[i+1] {
			n = i + 1
			return
		}
	}
	return 
}

func searchInsert(nums []int, target int) int {
    var n []int
    for i:= 0; i < len(nums); i++ {
        if nums[i] == target {
            return i
        } 
        
        if target < nums[i] && nums[i] > target {
            n = append(n, i)
        }
    }
    if len(n) == 0 {
        return len(nums)
    }
    
    return n[0]
}