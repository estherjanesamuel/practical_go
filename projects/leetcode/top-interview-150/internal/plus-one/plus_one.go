package math

import "fmt"

func AddOne(digits []int) []int {
	carry := 1
	for i := len(digits)-1; i >= 0; i-- {
		digitSum := digits[i] + carry
		fmt.Println(i,digitSum,digitSum < 10, digits[i],digitSum)
		if digitSum < 10 {
			digits[i] = digitSum
			carry = 0
		} else {
			digits[i] = digitSum % 10
			// fmt.Println(digits)
			carry = 1
		}
	}
	// fmt.Println(digits)
	if carry == 1 {
		digits = append([]int{1}, digits...)
	}
	fmt.Println(digits)

	return digits
}

func plusOne(digits []int) []int {
	if digits[len(digits)-1] == 9 {
		digits[len(digits)-1] = 1
		digits = append(digits, 0)
	} else {
		digits[len(digits)-1] += 1
	}


	// itoa := ""
	// for i := 0; i < len(digits); i++ {
	// 	itoa += strconv.Itoa(digits[i])
	// }

	// atoi, _ := strconv.(itoa)
	// atoi += 1
	// fmt.Println(itoa)
	// fmt.Println(atoi)
	// temp := atoi
	// digs := []int{}
	// for temp > 0 {
	// 	digit := temp % 10
	// 	// fmt.Println(temp,digit)
	// 	// fmt.Println(digit,temp)
	// 	digs = append(digs, digit)
	// 	temp /= 10
	// 	// fmt.Println(temp)
	// }
	// fmt.Println(digs)

	// start, end := 0, len(digs)-1
	// for start < end {
	// 	digs[start], digs[end] = digs[end], digs[start]
	// 	start++
	// 	end--
	// }
	// // fmt.Println(digs)
	return digits
}
