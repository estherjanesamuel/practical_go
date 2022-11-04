package peakindexinamountainarray

/*
	An array arr a mountain if the following properties hold:

	arr.length >= 3
	There exists some i with 0 < i < arr.length - 1 such that:
	arr[0] < arr[1] < ... < arr[i - 1] < arr[i]
	arr[i] > arr[i + 1] > ... > arr[arr.length - 1]

	Given a mountain array arr, return the index i such that arr[0] < arr[1] < ... < arr[i - 1] < arr[i] > arr[i + 1] > ... > arr[arr.length - 1].

	You must solve it in O(log(arr.length)) time complexity.
*/

// Intuition and Algorithm
// The comparison A[i] < A[i+1] in a mountain array looks like 
// [True, True, True, ..., True, False, False, ..., False]: 1 or more boolean Trues, 
// followed by 1 or more boolean False. For example, 
// in the mountain array [1, 2, 3, 4, 1], the comparisons A[i] < A[i+1] would be True, True, True, False.
// We can binary search over this array of comparisons, to find the largest index i such that A[i] < A[i+1]
func peakIndexInMountainArray(arr []int) int {
	low := 0
	high := len(arr)

	for low < high {
		mid := low + (high - low) / 2
		if arr[mid] > arr[mid + 1] {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return low
}

// The mountain increases until it doesn't. The point at which it stops increasing is the peak.
func peakIndexInMountainArray_1(arr []int) int {
    for i := 0; i < len(arr); i++ {
		if arr[i] > arr[i+1] {
			return i
		}
	}
	return 0
}
