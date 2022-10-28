package main

import (
	"fmt"
	"strings"
)

func main()  {
	// fmt.Println(twosum([]int{2,7,11,15}, 9))
	// fmt.Println(twosum([]int{3,3}, 6))
	// fmt.Println(twosum([]int{3,2,4}, 6))

	// fmt.Println(two_sum([]int{2,7,11,15}, 9))
	// fmt.Println(two_sum([]int{3,3}, 6))
	// fmt.Println(two_sum([]int{3,2,4}, 6))

	// fmt.Println(twoSum([]int{2,7,11,15}, 9))
	// fmt.Println(twoSum([]int{3,3}, 6))
	// fmt.Println(twoSum([]int{3,2,4}, 6))

	fmt.Println(TwoSum([]int{3,2,4}, 6))

}

func twoSum(nums []int, target int) []int {
	fmt.Print(strings.Trim(strings.Replace(fmt.Sprint(nums), " ", ",", -1), "[]"))
	indices := []int{}
	seen := map[int]int{}
	for i, num:= range nums {
		match := target - num
		if j, found := seen[match]; found {
			indices = append(indices, i)
			indices = append(indices, j)
			return indices
		}

		seen[num] = i
	}
	return indices
}

func twosum(nums []int, target int) []int{
	fmt.Print(strings.Trim(strings.Replace(fmt.Sprint(nums), " ", ",", -1), "[]"))
	
	seen := make(map[int]int)
	for i, num := range nums{
		 match := target - num
		if j, found := seen[match]; found {
			return []int{j,i}
		}
		seen[num] = i
	}
	return []int{}
}

func two_sum(nums []int, target int) []int {
	idx := []int{}
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if target-nums[i] == nums[j] {
				idx = append(idx, i)
				idx = append(idx, j)
			}
		}
	}
	return idx
}

func TwoSum(nums []int, target int) []int {
	fmt.Println(strings.Trim(strings.Replace(fmt.Sprint(nums), " ", ",", -1), "[]"))
	//create map or hasmap to store value 
	seen := map[int]int{}
	for i, num := range nums {
		fmt.Println("num:",num)
		match := target - num
		fmt.Println("match ",match)
		fmt.Println(seen)
		if j, found := seen[match]; found {
			fmt.Println(i,j)
			return []int{j,i}
		}
		seen[num] = i
	}
	return []int{}
}