package findthedistancevalue

import "math"

/*
Given two integer arrays arr1 and arr2, and the integer d, return the distance value between the two arrays.

The distance value is defined as the number of elements arr1[i] such that there is not any element arr2[j] where |arr1[i]-arr2[j]| <= d.

*/
func findTheDistanceValue(num, arr []int, distance int) int {
	dist := 0
	l := len(arr)
	for i := 0; i < len(num); i++ {
		d := 0
		for j := 0; j < len(arr); j++ {
			a := int(math.Abs(float64(num[i] - arr[j])))
			if a <= distance {
				break
			}
			d += 1
		}

		if d == l {
			dist += 1
		}
	}
	return dist
}

func findDistVal(nums, arr []int, d int) int {
	count := 0
	var j int
	for i := 0; i < len(nums); i++ {
		for j = 0; j < len(arr); j++ {
			if abs(nums[i]-arr[j]) <= d {
				break
			}
		}
		if j == len(arr) {
			count++
		}
	}
	return count
	
}

func abs(n int) int {
	if n < 0 {
		return n * -1
	} else {
		return n
	}
}
