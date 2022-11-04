package template

import "fmt"

func binarySeacrh(nums []int, target int) int {
	left, right := 0, len(nums)
	fmt.Println("len: ", right)
	for i := 0; i < right; i++ {
		mid := left + (right-left)/2
		fmt.Println(mid)
		if nums[mid] == target {
			return mid
		}
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}
