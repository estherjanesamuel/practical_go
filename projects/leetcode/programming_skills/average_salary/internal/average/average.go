package average

import (
	"sort"
)

func average(salary []float64) float64{
	// avg := 0.0
	var avg float64
	sort.Float64s(salary)

	// average without the higher and lower numer
	for i := 1; i < len(salary) - 1; i++ {
		avg += salary[i]
	}
	return avg / float64((len(salary) -2))
}

func averageInt(salary []int) float64{
	// avg := 0.0
	avg := 0
	sort.Ints(salary)

	// average without the higher and lower number
	for i := 1; i < len(salary) - 1; i++ {
		avg += salary[i]
	}
	return float64(avg) / float64((len(salary) -2))
}