package binarysearch


func findPeakElement(nums []int) int {
	
	// initialize the left variable to 0, which represent the leftmost index of the array
	// initialize the right variable to len(nums)-1 which represent the rightmost index of the array
	left, right := 0, len(nums)-1


	// enter a loop that continues as long as left is less than right.
	// if left becomes equal to right, we found a pick element
	for left < right {
		mid := left + (right - left) / 2

		// compare the element at index mid with its right neighbor nums[mid + 1] 
		// if nums[mid] is less than nums[mid + 1], it's mean there must be a peak element on the right side of mid.
		// so, update left to mid + 1 to search in the right half of the array
		if nums[mid] < nums[mid + 1] {
			left = mid + 1
		} else { // otherwise, if nums[mid] is greater than or equal to nums[mid + 1], its mean there could be a peak element on the left of mid, including the element at md itself, so update right to mid to search in the half of other array.
			right = mid
		} 

		// if left == right {
		// 	fmt.Println(left, right)
		// }
	}

	// if the loop ends, it means left is pointing to a peak element, so returnn left as the index of the peak elemetn
	return left

	/*
		The time complexity of this algorithm is O(log n) because in each iteration, we divide the search space in half by updating either left or right. Therefore, the number of iterations required is proportional to the logarithm of the size of the input array.
	*/
}
