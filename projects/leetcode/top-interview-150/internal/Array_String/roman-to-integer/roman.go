package romantointeger

import "fmt"

func romanToInt(s string) int {
	dict := make(map[string]int, 0)
	dict["I"] = 1
	dict["V"] = 5
	dict["X"] = 10
	dict["L"] = 50
	dict["C"] = 100
	dict["D"] = 500
	dict["M"] = 1000
	result := 0
	for i := 0; i < len(s); i++ {
		if i+1 < len(s) {
			if string(s[i]) == "I" {
				if string(s[i+1]) == "V" || string(s[i+1]) == "X" {
					result -= 2
				}
			} else if string(s[i]) == "X" {
				if string(s[i+1]) == "L" || string(s[i+1]) == "C" {
					result -= 20
				}
			} else if string(s[i]) == "C" {
				if string(s[i+1]) == "D" || string(s[i+1]) == "M" {
					result -= 200
				}
			}
		}
		result += dict[string(s[i])]
	}
	return result
}

func intToRoman(num int) string {
	s := [4][10]string{}
	s = [4][10]string{
		{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"},
		{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"},
		{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"},
		{"", "M", "MM", "MMM"},
	}
	fmt.Println(s[3][num/1000%10],num/1000%10,num/1000 , s[2][num/100%10],num/100%10,num/100 , s[1][num/10%10], num/10%10,num/10 , s[0][num%10],num%10)
	return s[3][num/1000%10] + s[2][num/100%10] + s[1][num/10%10] + s[0][num%10]
}

var (
	r0 = []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
	r1 = []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	r2 = []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
	r3 = []string{"", "M", "MM", "MMM"}
)

func intToRmn(num int) string {
	// this is efficient in Go, The 4 operands are evaluated, 
	// then a single allocated is made the exact size needed for the result
	return r3[num/1000%10] + r2[num/100%10] + r1[num/10%10] + r0[num%10]
}
