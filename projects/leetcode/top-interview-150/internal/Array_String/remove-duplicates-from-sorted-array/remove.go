package removeduplicatesfromsortedarray

import "fmt"

func Remove(nums []int) int {
	j := 0
	exist := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		if !exist[nums[i]] {
			exist[nums[i]] = true
			nums[j] = nums[i]
			j++
		}
	}
	fmt.Println(nums[:j])
	return j
}

func remove(nums []int) int {
	i := 0
	for _, num := range nums {
		if i < 2 || num > nums[i-2] {
			nums[i] = num
			i++
		}
	}
	fmt.Println(nums[:i])
	fmt.Println(nums)
	return i
}

func removeDuplicate(nums []int) int {
	j := 0
	for _, num := range nums {
		fmt.Println("nums: ", nums, " | ", "j > 2: ", j < 2)
		if j < 2 {
			fmt.Println("j: ", j, " | ", "nums: ", nums, " | ", "num: ", num, " | ", "nums[j]: ", nums[j])
			nums[j] = num
			j++

		} else {
			fmt.Println("j: ", j, " | ", "nums: ", nums, " | ", "num > nums[j-2]: ", num > nums[j-2], " | ", " num: ", num, " | ", " | ", "nums[j-2]: ", nums[j-2], " | ", "num > nums[j-2]: ", num > nums[j-2])
			if num > nums[j-2] {
				nums[j] = num
				j++
			}
		}
	}
	fmt.Println(nums)
	return j
}
func RemoveDuplicate(nums []int) int {
	j := len(nums)
	counts := make(map[int]int)
	for i, num := range nums {
		counts[num]++
		if counts[num] > 2 {
			fmt.Println(nums[i:])
			fmt.Println(nums[:i])
			j--
			fmt.Println(counts[num], " | ", i)
		}
		fmt.Println("i: ", i, "|", "count: ", counts, " | ", "nums: ", nums, " | ", "num: ", num, " | ", "j: ", j)
	}
	fmt.Println(nums[:j])
	return j
}

func remove_duplicate(nums []int) int {
	// {[]int{1,1,1,2,2,2,3},5},
	expected_length := 0
	for i := 0; i < len(nums); i++ {
		if expected_length < 2 {
			nums[expected_length] = nums[i]
			expected_length++
		} else {
			if nums[i] == nums[expected_length-2] {
				continue
			} else { // else continue
				nums[expected_length] = nums[i]
				expected_length++
			}
		}
	}
	fmt.Println(nums[:expected_length])
	return expected_length
}

/*
Given an integer array nums sorted in non-decreasing order,
remove some duplicates in-place such that each unique element appears at most twice.
The relative order of the elements should be kept the same.

Since it is impossible to change the length of the array in some languages,
you must instead have the result be placed in the first part of the array nums.
More formally, if there are k elements after removing the duplicates,
then the first k elements of nums should hold the final result.
It does not matter what you leave beyond the first k elements.

Return k after placing the final result in the first k slots of nums.

Do not allocate extra space for another array.
You must do this by modifying the input array in-place with O(1) extra memory.
*/
func Remove_Duplicate(nums []int) int {
	// {[]int{1,1,1,2,2,2,3},5},
	k := 0
	for i := 0; i < len(nums); i++ {
		if k < 2 {
			k++ // store the the value tha < 2
		} else {
			if nums[i] == nums[k-2] {
				continue
			} else {
				nums[k] = nums[i]
				k++
				// {[]int{1,1,2,2,2,2,3},5},
				// {[]int{1,1,2,2,3,2,3},5},

			}
		}
	}
	return k
}
