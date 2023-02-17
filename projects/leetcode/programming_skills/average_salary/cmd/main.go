package main

import (
	"fmt"
	"strconv"
)

func main()  {
	s := fmt.Sprintf("%b", 4294967293)
	// p := fmt.Sprintf("%d", 11111111111111111111111111111101)
	fmt.Println(s)

	p := 11
	fmt.Println(p)
	strs := []string{"00000000000000000000000000001011","00000000000000000000000010000000", "11111111111111111111111111111101"}
	for _, str := range strs {

		i, err := strconv.ParseInt(str,2,64)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Print(i," | ")
	}
	fmt.Println()
	fmt.Println(hammingWeight(p))
}

func hammingWeight(num int) int {
	result := 0
	for i := 0; i < 32; i++ {
		if (num & (1 << i)) >> i == 1 {
			result++
		}
	}
	return result
}
