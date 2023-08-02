package binarysearch

import "fmt"

/*
Given a sorted array of distinct integers and a target value, return the index if the target is found. If not, return the index where it would be if it were inserted in order.

You must write an algorithm with O(log n) runtime complexity.
*/
func searchInsert(nums []int, target int) int {

	// initialize the left variable to 0, which represent the leftmost index of the array
	// initialize the right variable to len(nums)-1 which represent the rightmost index of the array
	left, right := 0, len(nums)-1
	
	// enter the loop that continues as long as left is less than or equal to right
	for left <= right {
		// calculate the middle index 'middle' using formula `left + (right-left) / 2` to avoid potential overflow
		middle := left + (right - left ) / 2
		
		// check if the value at the middle index nums[middle] is equal to the target. if it is, return mId as the target is found
		if nums[middle] == target {
			return middle
		} else if nums[middle] < target{ // if the value at nums[middle] is less than the target, update right to mid - 1 to search in the left half of the array.
			left = middle + 1
		} else { // if the value at nums[middle] is greater than the target, update left to mid + 1 to search in the right half of the array.
			right = middle - 1
		}
	}

	// if the loop ends withou finding the target, return left as the index where the target would be inserted.
	return left

	/*
		The time complexity of this algorithm is O(log n) because in each iteration, we
		divide the search in half by updating either left or right.
		Therefore, the number of iterations required is proportional to the logarithm of the size of the input array.
	*/
}

func search_insert(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := left + (right - left) / 2
		fmt.Println(left, right,mid)
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}
