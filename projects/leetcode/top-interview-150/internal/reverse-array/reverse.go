package reversearray



func reverse(nums []int) []int {
	start  := 0 
	end := len(nums) -1
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start ++
		end --
	}
	return nums
}

func reverseII(nums []int, start ,end int) []int {
	if start >= len(nums) || end >= len(nums)  {
		return []int{}
	}
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start ++
		end --
	}
	return nums
}


func reverseIII(nums []int, start ,end int) []int {
	if start >= len(nums) || end >= len(nums)  {
		return []int{}
	}
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start ++
		end --
	}
	return nums
}

func rvrs(str string, start, end int) string {
	nums := []rune(str)
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start ++
		end --
	}
	return string(nums)
}

func rvrsII(s string) string {
	start, end := 0, len(s)-1
	runes := []rune(s)
	// missed the if instead of for
	if start < end {
		runes[start], runes[end] = runes[end], runes[start]
		start ++
		end --
	}
	return string(runes)
}

// func rvrsII(s string) string {
// 	start, end := 0, len(s)-1
// 	runes := []rune(s)
// 	for start < end {
// 		runes[start], runes[end] = runes[end], runes[start]
// 		start ++
// 		end --
// 	}
// 	return string(runes)
// }