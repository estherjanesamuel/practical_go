package Hashmap

import (
	"fmt"
)

/*
	Given an integer array nums and an integer k,
	return true if there are two distinct indices i and j
	in the array such that nums[i] == nums[j] and abs(i - j) <= k.
*/
func containsNearbyDuplicate(nums []int, k int) bool {
	for i := 0; i < len(nums)-1; i++ {
		// fmt.Println(i, nums[i])
		for j := i + 1; j <= len(nums)-1; j++ {
			// fmt.Println(j, nums[j])
			fmt.Println( i, j ,nums[i] == nums[j],  nums[i],  nums[j],  i-j, k, " | ", (i - j) <= k,k, i- j, "|",  j - i,(j - i) <= k, j,i,k,  len(nums)-1)
			if nums[i] == nums[j] && (j - i) <= k {
				return true
			}
		}
	}
	return false
}

func contains_duplicate(nums []int, k int) bool {
	size := len(nums)-1
	for i := 0; i < size; i++ {
		for j := i+1; j <= size; j++ {
			if nums[i] == nums[j] && j - i <= k {
				return true
			}
		}
	}
	return false
}

// return true if there are two distinct indices i and j
// 	in the array such that nums[i] == nums[j] and abs(i - j) <= k.
func nearby_distinc(nums []int, k int) bool {
	// iterate one by one using two for loop
	for i := 0; i < len(nums); i++ {
		// fmt.Println(i, nums[i])
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				fmt.Println(i,j, k)
				if j - i <= k {
					return true
				}
			}
		}
	}
	return false
}