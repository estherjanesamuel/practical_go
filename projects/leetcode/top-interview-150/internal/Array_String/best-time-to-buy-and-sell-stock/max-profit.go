package array_string

import (
	"fmt"
)

func maxProfit(nums []int) int {
	/*
		7,
		1,
		5,
		3,
		6,
		4

	*/
	low := nums[0]
	max := nums[0]
	max_profit := 0
	for i := 1; i < len(nums); i++ {
		current_profit := 0
		fmt.Println(nums, nums[i], i, low, max, low > nums[i], max < nums[i],max_profit, current_profit)

		if low > nums[i] {
			if i+1 < len(nums) {
				low = nums[i]
				max = nums[i]
			}

		}
		if max < nums[i] {
			max = nums[i]
		}
		// fmt.Println(low, max)
		// fmt.Println(low, max, nums[i], i)
		// fmt.Println(low, max, i, count)
		current_profit = max - low
		if current_profit > max_profit {
			max_profit = current_profit
		}
	}
	return max_profit
}

func cuanMax(nums []int) int {
	/*
		7,
		1,
		5,
		3,
		6,
		4

	*/
	low := nums[0]
	max := nums[0]
	max_profit := 0
	for i := 1; i < len(nums); i++ {
		// fmt.Println(i,nums[i],len(nums))
		current_profit := 0
		if low > nums[i] {
			if i+1 < len(nums) {
				low = nums[i]
				max = nums[i]
			}

		}
		if max < nums[i] {
			max = nums[i]
		}
		current_profit = max - low
		if current_profit > max_profit {
			max_profit = current_profit
		}
	}
	return max_profit
}

func maxCuan(prices []int) int {
	low,max := prices[0],prices[0]
	max_profit := 0
	for i := 1; i < len(prices); i++ {
		fmt.Println(max, low,low < prices[i])
		if low > prices[i] {
			low = prices[i]
			max = prices[i]
		}
		if max < prices[i] {
			max = prices[i]
		}
		current_profit := max -low
		if current_profit > max_profit {
			max_profit = current_profit
		}

	}
	return max_profit
}

func maxCuanII(prices []int) int {
	low,max := prices[0],prices[0]
	max_profit := 0
	total_profit := 0
	profit := []int{}
	for i := 1; i < len(prices); i++ {
		fmt.Println(low, max, low > prices[i], low < prices[i])
		if low > prices[i] {
			low = prices[i]
			max = prices[i]
		}
		if max < prices[i] {
			max = prices[i]
		}
		current_profit := max -low
		if current_profit > max_profit {
			// profit[low] = current_profit
			profit = append(profit, current_profit)
			max_profit = 0
			low = prices[i]
		}
	}
	fmt.Println(profit)
	for _, prft := range profit {
		total_profit += prft
	}
	return total_profit
}