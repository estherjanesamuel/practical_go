package math

import (
	"fmt"
	"testing"
)


func TestAddBinary(t *testing.T)  {
	fmt.Println("")
	tests := []struct{
		a,b,want string
	}{
		{"1010","1011","10101"},
		{"11","1","100"},
		{"0","0","0"},
	}
	for _, ts := range tests {
		testAddBinary := fmt.Sprintf("%v %v %v",ts.a, ts.b, ts.want)
		t.Run(testAddBinary, func(t *testing.T) {
			got := AddBin3(ts.a, ts.b)
			if got != ts.want {
				t.Errorf("got %v, wanted %v", got, ts.want)
			}
		})
	}
}







