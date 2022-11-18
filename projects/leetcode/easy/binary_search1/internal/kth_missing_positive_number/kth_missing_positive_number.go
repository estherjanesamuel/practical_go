package kthmissingpositivenumber

import "fmt"

/*
Given an array arr of positive integers sorted in a strictly increasing order, and an integer k.

Return the kth positive integer that is missing from this array.



Example 1:

Input: arr = [2,3,4,7,11], k = 5
Output: 9
Explanation: The missing positive integers are [1,5,6,8,9,10,12,13,...]. The 5th missing positive integer is 9.
Example 2:

Input: arr = [1,2,3,4], k = 2
Output: 6
Explanation: The missing positive integers are [5,6,7,...]. The 2nd missing positive integer is 6.
*/

func findKthPositive(arr []int, k int) int {
	left, right, mid := 0, len(arr), 0
	for left < right {
		mid = left + (right-left)/2
		if arr[mid]-mid-1 < k {
			left = mid + 1
		} else {
			right = mid
		}
		fmt.Println(left, mid, right, k)
	}
	return left + k
}

/*
 === Solutions ==
 https://leetcode.com/problems/kth-missing-positive-number/solutions/2283712/o-logn-cpp-ez-amazon-google-interview/
*/
