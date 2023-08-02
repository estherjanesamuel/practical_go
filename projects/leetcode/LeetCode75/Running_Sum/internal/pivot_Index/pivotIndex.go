package pivotindex

import "fmt"


func pivotIndex(nums []int) int {
	var ls, rs int
	rs = sum(nums)

	for i, curr := range nums {
		if i != 0 {
			ls += nums[i-1]
		}
		rs -= curr

		if ls == rs {
			return i
		}
	}
	return -1
}

func sum(nums []int) int {
	sum := 0
	for _, valueInt := range nums {
		sum += valueInt
	}

	return sum
}

func PivotIndex(nums []int) int {
	return pivotIndex(nums) + pivotIndexII(nums)
}


func pivotIndexII(nums []int) int {
	var l, r int
	for _, v := range nums {
		r += v
	}
	fmt.Printf("r : %v | l : %v \n", r, l)
	for i, curr := range nums {
		// fmt.Printf("i : %v, curr: %v \n", i, curr)
		if i != 0{
			l += nums[i-1]
		}
		r -= curr
		fmt.Printf("i: %v l : %v, r: %v \n",i, l, r)
		if l == r {
			return i
		}
	}
	return -1
}