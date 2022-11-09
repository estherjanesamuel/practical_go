package findfirstandlastpositionofelementinsortedarray

/*
Given an array of integers nums sorted in non-decreasing order,
find the starting and ending position of a given target value.

If target is not found in the array, return [-1, -1].

You must write an algorithm with O(log n) runtime complexity.

e 1:

Input: nums = [5,7,7,8,8,10], target = 8
Output: [3,4]
Example 2:

Input: nums = [5,7,7,8,8,10], target = 6
Output: [-1,-1]
Example 3:

Input: nums = [], target = 0
Output: [-1,-1]


Constraints:

0 <= nums.length <= 105
-109 <= nums[i] <= 109
nums is a non-decreasing array.
-109 <= target <= 109
*/
func searchRangeTLE(nums []int, target int) []int {
	left, right := 0, len(nums)
	var i, j int

	for left < right {

		mid := left + (right-left)/2
		if nums[mid] == target {
			if i > 0 {
				j = mid
			} else {
				i = mid
			}
			left =mid
		} else if nums[mid] < target {
			left = mid + 1 
		} else {
			right = mid - 1
		}
	}
	if i > 0 {
		if j > 0 {
			return []int{i, j}
		}
		return []int{i, -1}
	}
	return []int{-1, -1}
}

func searchRange(nums []int, target int) []int {
	start, end := 0, len(nums)-1
	ans := []int{-1, -1}
	for start <= end {
		m := start + (end-start)/2
		if nums[m] == target {
			start = m - 1
			end = m + 1
			// search for first index with while loop in -ve  direction
			for start >= 0 && nums[start] == target {
				start--
			}
			ans[0] = start + 1
			// search for ending index in +ve direction
			for end < len(nums) && nums[end] == target {
				end++
			}
			ans[1] = end - 1
			return ans
		} else if target > nums[m] {
			start = m + 1
		} else {
			end = m - 1
		}
	}
	return ans
}
