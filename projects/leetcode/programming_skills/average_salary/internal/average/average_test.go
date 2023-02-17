package average

import (
	"fmt"
	"testing"
)

func TestAverageSalary(t *testing.T) {
	tests := []struct {
		salary []float64
		want float64
	}{
		{[]float64{4000,3000,1000,2000}, 2500},
		{[]float64{1000,2000,3000}, 2000},
	}

	for _, ts := range tests {
		testAverageSalary := fmt.Sprintf("%v, %v", ts.salary, ts.want)
		t.Run(testAverageSalary, func(t *testing.T) {
			got := average(ts.salary)
			if got != ts.want {
				t.Errorf("got %v, wanted %v", got, ts.want)
			}
		})
	}


}

func TestAverageIntSalary(t *testing.T) {
	tests := []struct {
		salary []int
		want float64
	}{
		{[]int{4000,3000,1000,2000}, 2500},
		{[]int{1000,2000,3000}, 2000},
		{[]int{1000,2000,3000,4000,5000,6000,7000,8000,9000,10000}, 5500},
	}

	for _, ts := range tests {
		testAverageSalary := fmt.Sprintf("%v, %v", ts.salary, ts.want)
		t.Run(testAverageSalary, func(t *testing.T) {
			got := averageInt(ts.salary)
			if got != ts.want {
				t.Errorf("got %v, wanted %v", got, ts.want)
			}
		})
	}


}

