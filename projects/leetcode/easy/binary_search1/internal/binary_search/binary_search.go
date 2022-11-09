package binarysearch

import "fmt"

func BinarySearch(arr []int, target int) int {
	return binarySearchDbg(arr, target) + binarySrch(arr, target)
}

/*
Given an array of integers nums which is sorted in ascending order,
and an integer target, write a function to search target in nums.
If target exists, then return its index. Otherwise, return -1.

You must write an algorithm with O(log n) runtime complexity.
e.g..([]int{-1,0,3,5,9,12}, 9))
*/
func binarySearch(arr []int, target int) int {
	left, right := 0, len(arr)
	for left <= right {
		mid := left + (right-left)/2
		fmt.Println("mid: ", mid)
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}


func binarySearchDbg(arr []int, target int) int {
	left, right := 0, len(arr)
	fmt.Println("left: ", left)
	fmt.Println("right: ", right)
	fmt.Println()
	for left <= right {
		mid := left + (right-left)/2
		fmt.Println("mid: ", mid)
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
		fmt.Println("left: ", left)
		fmt.Println("right: ", right)
		fmt.Println()
	}
	return -1
}


// e.g..([]int{-1,0,3,5,9,12}, 9))
func binarySrch(arr []int, target int) int {
	left, right := 0, len(arr)
	for left <= right {
		mid := left + (right - left) / 2
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}
