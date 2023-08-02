package majorityelement

import "fmt"

// refactor or better implementation for Majority functions
func MajorityElement(nums []int) int {
	counts := make(map[int]int)
	for _, num := range nums {
		counts[num]++
	}

	majority := nums[0]
	maxCount := counts[nums[0]]
	for num, count := range counts {
		if count > maxCount {
			majority = num
			maxCount = count
		}
	}
	return majority
}

func Majority(nums []int) int {
	exist := make(map[int]bool)
	majority := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if !exist[nums[i]] {
			exist[nums[i]] = true
			majority[nums[i]] = 1

		} else {
			majority[nums[i]] = majority[nums[i]] + 1
		}
	}
	return GetKeyWithMaxValue(majority)
}

func GetKeyWithMaxValue(nums map[int]int) int {
	max, value := 0,0
	for key, val := range nums {
		if val > value {
			max = key
			value = val
		}
	}
	return max
}

func majority(nums []int) int {
	fmt.Println(len(nums))
	major, cnt := nums[0], 1
	for i := 1; i < len(nums); i++ {
		fmt.Println(" i: ",i," | "," major: ",major," | "," count: ",cnt," | "," nums[i]: ",nums[i]," | "," nums: ", nums)
		if cnt == 0 {
			cnt ++
			major = nums[i]
		} else if major == nums[i] {
			cnt++
		} else {
			cnt--
		}
	}

	return major
}

func major(nums []int) int {
	// []int{2,2,1,1,1,2,2}
	// []int{2,1,1,1,2,1,2}
	/*
	[]int{
		0 : 2,
		1 : 1,
		2 : 1,
		3 : 1,
		4 : 2,
		5 : 1,
		6 : 2}
	*/
	maj := nums[0]
	count := 1
	for i := 1; i < len(nums); i++ {
		if count == 0 {
			count += 1
			maj = nums[i]
		} else if maj == nums[i]{
			count += 1
		} else {
			count -= 1
		}
	}
	return maj
}

func Maj(nums []int) int {
	major := nums[0] 
	count := 1
	for i :=1; i < len(nums); i++{
		if count== 0 {
			count++
			major = nums[i];
		} else if major==nums[i] {
			count++
		} else {
			count--;
		}
	}
	return major;
}