package searchinrotatedsortedarray

/*
There is an integer array nums sorted in ascending order (with distinct values).
Prior to being passed to your function, 
nums is possibly rotated at an unknown pivot index k (1 <= k < nums.length) 
such that the resulting array is [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]] 
(0-indexed). For example, [0,1,2,4,5,6,7] might be rotated at pivot index 3 and become [4,5,6,7,0,1,2].

Given the array nums after the possible rotation and an integer target, 
return the index of target if it is in nums, or -1 if it is not in nums.

You must write an algorithm with O(log n) runtime complexity.

Example 1:
Input: nums = [4,5,6,7,0,1,2], target = 0
Output: 4

Example 2:
Input: nums = [4,5,6,7,0,1,2], target = 3
Output: -1

Example 3:
Input: nums = [1], target = 0
Output: -1

Constraints:

1 <= nums.length <= 5000
-104 <= nums[i] <= 104
All values of nums are unique.
nums is an ascending array that is possibly rotated.
-104 <= target <= 104
*/
func search(nums []int, target int) int {
	pivot := findMin(nums)

	if nums[pivot] == target {
		return pivot
	}

	if pivot == 0 {
		return binarySearch(nums, target, 0, len(nums)-1)
	}
	if nums[pivot] < target && nums[len(nums) - 1] >= target {
		return binarySearch(nums,target,pivot,len(nums)-1)
	} else {
		return binarySearch(nums,target, 0, pivot-1)
	}
}

func binarySearch(nums []int, target, left, right int) int {
	for left <= right {
		// mid := left + (right - left) / 2
		mid := left + (right - left)
		if nums[right] == target {
			return right
		} 
		if nums[left] == target {
			return left
		} 

		switch {
		case nums[mid] < target:
			left = mid + 1
		case nums[mid] > target:
			right = mid - 1
		default:
			return mid
		}
	}
	return -1
}

func findMin(nums []int) int {
	l, r :=0, len(nums) -1

	if nums[l] < nums[r] {
		return 0
	}
	for r - l > -1 {
		m := r - l
		
		switch  {
		case m > 0 && nums[m] < nums[m-1]:
			return m
		case nums[m] > nums[r]:
			l = m + 1
		case nums[m] < nums[l]:
			r = m - 1
		default:
			return l
		}
	}
	return 0
}