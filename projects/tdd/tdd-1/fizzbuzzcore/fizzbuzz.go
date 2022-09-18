package fizzbuzzcore

import "strconv"

func GetValue(input int) string {
	var s string
	
	if input % 3 == 0 {
		s += "fizz"
	}
	if input % 5 == 0 {
		s += "buzz"
	}

	if s == "" {
		s += strconv.Itoa(input)
	}
	return s
}