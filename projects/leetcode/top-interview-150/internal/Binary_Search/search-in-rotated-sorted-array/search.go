package binarysearch


func search(nums []int, target int) int {
	// check null value and sanity check 
	if len(nums) <= 0 {
		return -1
	}

	// initialize two pointers left and right, pointing to the start and end of the array respectively.
	left, right := 0, len(nums) -1
	
	for left <= right {
		// calculate the middle index
		mid := (left + right) / 2

		// check if nums[mid] is equal to the target, return mid
		if nums[mid] == target {
			return mid
		}

		// check if the left half of the array is sorted (from left to mid)
		if nums[left] <= nums[mid] {
			// if the target is within within the range nums[left] to nums[mid] update right = mid - 1
			if nums[left] <= target && target <= nums[mid]{
				right = mid - 1
			} else { // otherwise update left = mid + 1
				left = mid + 1
			}
		} else { // otherwise (the right half of the array is sorted, from mid to right)
			// if the target is within the range nums[mid] to nums[right] update left = mid + 1
			if target <= nums[mid] && target <= nums[right] {
				left = mid + 1
			} else { // otherwise update right = mid - 1
				right = mid - 1
			}
		}
	}
	return -1
}