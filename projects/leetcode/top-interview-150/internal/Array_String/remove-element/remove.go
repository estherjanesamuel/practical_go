package removeelement

import (
	"fmt"
	"sort"
)

func remove(nums []int, val int) int {
	result := 0
	new := []int{}
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			result += 1
		} else {
			new = append(new, nums[i])
		}
	}
	nums = new
	fmt.Println(new)
	fmt.Println(nums)
	return len(new)
}

func removeIndex(arr []int, pos int) []int {
	new := make([]int,len(arr)-1)
	k := 0
	for i := 0; i < len(arr) -1; {
		if i != pos {
			new[i] = arr[k]
			k++
		} else {
			k++
		}
		i++ 
	}
	sort.Ints(new)
	return new
}

func RemoveElement(nums []int, val int) int {
	index := 0

	for _, num := range nums {
		ok := num != val
		fmt.Println("index:", index)
		fmt.Println(num,val)
		fmt.Println(ok)
		if ok {
            nums[index] = num
            index += 1
		}
	}
	fmt.Println(nums)
	fmt.Println(nums[:index])
    return index
}

func Remove(nums []int, val int) int {
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[j] = nums[i]
			j++ 
		}
	}
	fmt.Println(nums[:j])
	return j
}