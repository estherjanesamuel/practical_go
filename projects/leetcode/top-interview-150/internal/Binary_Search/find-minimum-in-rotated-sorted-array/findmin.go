package binarysearch

/*
	To find the minimum element in a sorted and rotated array of unique elements
	with an algorithm that runs in O(log n) you can use a modified binary
	search approach.
*/
func findMin(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := (right + left) / 2
		// if the mid element is greater than the rightmost element,
		// the minimum element lies to the right of mid.
		if nums[mid] > nums[right] {
			left = mid + 1
		} else { // if the mid element is less than or equal to the rihtmost element
			// the minimum element lies to the left of or at the mid
			right = mid
		}
	}
	// the loop will terminate when left == right, pointing to the minimum element
	return nums[left]
}
