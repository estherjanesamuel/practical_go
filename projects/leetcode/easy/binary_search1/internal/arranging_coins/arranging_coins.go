package arrangingcoins

import "fmt"

/*
You have n coins and you want to build a staircase with these coins. The staircase consists of k rows where the ith row has exactly i coins. The last row of the staircase may be incomplete.

Given the integer n, return the number of complete rows of the staircase you will build.
n = 5
$
$$
$$
output = 2
Explanation: Because the 3rd row is incomplete, we return 2.
*/
func arrangeCoins(n int) int {
	left, right := 1, n
	for left < right {
		mid := left + (right - left) / 2
		fmt.Println(left,mid, right)
		if n < mid {
			left = mid + 1
		} 
		right = mid - 1
	}
	return left
}

func arrCoins(n int) int {
	counter, sum := 1, 0
	for {
		sum += counter
		// fmt.Println("sum: ", sum)
		// fmt.Println(n,"_",  sum, "_" ,counter)
		if n < sum {
			return counter - 1
		} else if n == sum {
			return counter
		}
		counter += 1
	}
}