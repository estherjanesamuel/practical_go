package findthedistancevalue

import "math"

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
