package twosum

/*
	params, array of integers nums and an integer target
	return indices of the two numbers such that they add up to target
	You may assume that each input would have exactly one solution,
	and you may not use the same element twice
*/
func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}

	}
	return []int{}
}

/*
	It only go through the list once. It's shorter and easier to understand. Hope this can help someone. Please tell me if you know how to make this better :)
	The key to the problem is that there is always only 1 pair of numbers that satisfy the condition of adding together to be target value.
	we can assume that for all numbers in the list(x1,x2, ....xn)
	that xa + xb = target to solve this with a single pass of the list we can changethe 
	equation above to xa = target - xb and since we know the target as long as we maintain a record
	of all previous values in the list we can compare the current value (xa) ti it's only pair, it it exist, in record of all previous values (xb)
*/
func sumTwo(nums []int, target int) []int {
	result := []int{}
	dict := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if val, exist := dict[target-nums[i]]; exist {
			result = append(result, val, i)
			return result
		}
		dict[nums[i]] = i
	}
	return result

}

/*
	targe 6
	[length 3   idx 0, 3 		idx 1, 2
		0. 3		3 + 2 = 5	2 + 4 = 6
		1. 2		3 + 4 = 7
		2. 4

	6,_

	[1_2]
    two_sum_test.go:26: got, []. wanted [1 2]

*/
