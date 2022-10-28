package main

import (
	"fmt"
	"sort"
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

	 fmt.Println(Two_sum([]int{2,7,11,15}, 9))
	 fmt.Println(Two_sum([]int{3,3}, 6))
	 fmt.Println(Two_sum([]int{3,2,4}, 6))

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

// Using a Hash Map — O(N) Time, O(N) Space
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

//A Brute Force Solution — O(N²) Time, O(1) Space
//https://levelup.gitconnected.com/two-number-sum-in-golang-355627d6c861
func Two_Sum(array []int, target int) []int {
    for i := 0; i < len(array); i++ {
        var start = array[i]
        for j := i+1; j < len(array); j++ {
            var end = array[j]
            if (start + end) == target {
                return []int{i, j}
            }
        }
    }

    return []int{}
}


//Sorting & Crunching— O(NlogN) Time, O(1) Space
func Two_sum(nums []int, target int) []int {
    sort.Ints(nums)
    start, end := 0, len(nums)-1
    for start < end {
        currentSum := nums[start] + nums[end]
        if currentSum == target {
            return []int{start, end}
        } else if currentSum < target {
            start += 1
        } else {
            end -= 1
        }
    }
    return []int{}
}