package largestperimetertriangle

import "sort"

/*
Given an integer array nums,
return the largest perimeter of a triangle with a non-zero area,
formed from three of these lengths.
f it is impossible to form any triangle of a non-zero area,
return 0.

e.g.
Input: nums = [2,1,2]
Output: 5
Explanation: You can form a triangle with three side lengths: 1, 2, and 2.

Input: nums = [1,2,1,10]
Output: 0
Explanation:
You cannot use the side lengths 1, 1, and 2 to form a triangle.
You cannot use the side lengths 1, 1, and 10 to form a triangle.
You cannot use the side lengths 1, 2, and 10 to form a triangle.
As we cannot use any three side lengths to form a triangle of non-zero area, we return 0.

*/

/* the goal is to find the largest number n1,
	where two others numbers, n2 and n3, exist, 
	and n1 > n2 && n1 > n3, and n1 < n2 + n3

	if we sort our sizes, we just need to find the largest nums[i] 
	such as nums[i] < nums[i-1] + nums[i-2]
	so, we analuze triplets largest to smallest, and retun the perimeter 
	for the first triplet that matches our criterion

*/
func largestPerimeter(nums []int) int {
	sort.Ints(nums)

	for i := len(nums) - 1; i >= 2; i-- {
		if nums[i] < nums[i - 1] + nums[i -2] {
			return nums[i] + nums[i-1] + nums[i - 2]
		}
	}
	return 0
}