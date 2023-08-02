package intervals

import (
	"fmt"
	"strconv"
)

func summaryRanges(nums []int) []string {
	result := []string{}
	var head int

	for i := range nums {
		if i < len(nums)-1 {
			fmt.Println(head, i, len(nums)-1," | ", nums[i], nums[i]+1, nums[i+1], nums[i]+1 == nums[i+1],head == i ,result)
			if nums[i]+1 == nums[i+1] {
				continue
			}
		}
		if head == i {
			result = append(result, strconv.Itoa(nums[i]))
		} else {
			tmp := strconv.Itoa(nums[head]) + "->" + strconv.Itoa(nums[i])
			result = append(result, tmp)
		}

		head = i + 1
	}

	return result
}

/*
	return the smallest sorted list of ranges that cover all the numbers in the array exactly
	that is, each element of nums is covered by exactly one of the ranges, and there is no integer x such 
	that x is in one of the ranges but not in nums

	Each range [a,b] in the list should be output as:

	"a->b" if a != b
	"a" if a == b
*/
func sumRanges(nums []int) (result []string) {
	// input 		- 	ouput
	// [0_1_2_4_5_7]_[0->2_4->5_7]
	var head int
	for i := range nums{
		if i < len(nums) -1 {
			fmt.Println(i,head,nums[i],nums[i+1],nums[i] + 1 == nums[i+1],head == nums[i],result)
			if nums[i] + 1 == nums[i+1] {
				continue
			}
		} 

		if nums[head] == nums[i] {
			result = append(result, strconv.Itoa(nums[head]))
			fmt.Println(nums[head],i)
		} else {
			tmp := strconv.Itoa(nums[head]) + "->" + strconv.Itoa(nums[i])
			result = append(result, tmp)
		}
		head = i + 1
	}
	return result
}


func summary_ranges(nums []int) (result []string) {
	var head int

	for i := range nums {
		if i < len(nums) - 1 {
			if nums[i] + 1 == nums[i + 1] {
				continue
			}
		}
		if nums[head] == nums[i] {
			result = append(result, strconv.Itoa(nums[head]))
		} else {
			tmp := strconv.Itoa(nums[head]) + "->" + strconv.Itoa(nums[i])
			result = append(result, tmp)
		}
		head = i + 1
	}
	return result
}