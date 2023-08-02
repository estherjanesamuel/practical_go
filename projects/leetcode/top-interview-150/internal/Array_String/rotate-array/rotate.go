package rotatearray

import (
	"fmt"
)

func rotate(nums []int, k int) []int {
	// [1_2_3_4_5_6_7]
	c := make([]int, len(nums[:k+1]))
	copy(c, nums[:k+1])
	fmt.Println(nums[k:])
	fmt.Println(nums[k+1:])
	fmt.Println(nums[:k+1])
	fmt.Println(nums[:k])
	copy(nums[:k+1],c)
	fmt.Println(c)
	copy(nums[:k], nums[k+1:])
	fmt.Println(nums)
	fmt.Println(nums[k:])
	copy(nums[k:], c)
	fmt.Println(nums)

	
	// fmt.Println(nums[k:])
	// fmt.Println(nums[:k+1])


	return nums
}
/*
 void rotate(vector<int>& nums, int k) {
        vector<int> numcopy = nums;
        for(int i=0;i<nums.size();i++) 
            nums[(i+k)%nums.size()]=numcopy[i];
    }
*/

func Rotate(nums []int, k int) []int {
	numcopy := nums
	for i := 0; i < len(nums); i++ {
		nums[(i+k)%len(nums)]=numcopy[i]
	}
	fmt.Println(numcopy, nums)
	return numcopy
}

func rtt(nums []int, k int) []int {
    n := len(nums)
    fmt.Println("before : ", k, len(nums))

    k = k % len(nums)
    fmt.Println("after : ",k, len(nums))
    // k = k % n  // handle cases where k is greater than the array length
    
    // Reverse the entire array
    nums = reverse(nums, 0, n-1)
    fmt.Println("Reverse the entire array", nums, " | ", n-1)

    // Reverse the first k elements
    nums = reverse(nums, 0, k-1)
    fmt.Println("Reverse the first k elements", nums, " | ", k-1)
    
    // Reverse the remaining n-k elements
    nums = reverse(nums, k, n-1)
    fmt.Println("Reverse the remaining n-k elements", nums, " | ", k, n-1)

	return nums
}
    
func reverse(nums []int, start int, end int) []int {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
        start += 1
        end -= 1
	}
	return nums
}

/*
To rotate an array to the right by k steps, you can use several approaches. One simple method is to reverse the entire array, then reverse the first k elements, and finally reverse the remaining n-k elements. Here's the implementation in Python:
def rotate(nums, k):
    n = len(nums)
    k = k % n  # handle cases where k is greater than the array length
    
    count = 0  # tracks the number of elements that have been moved
    start = 0  # starting index for each cycle
    
    while count < n:
        curr = start
        prev = nums[start]
        
        # Move each element to its correct position in the cycle
        while True:
            next_idx = (curr + k) % n
            nums[next_idx], prev = prev, nums[next_idx]
            curr = next_idx
            count += 1
            
            # Break the cycle when we return to the starting position
            if curr == start:
                break
                
        start += 1  # Move to the next cycle
*/

func rotte(nums []int, k int) []int	 {
	n := len(nums)
    // k = k % n  # handle cases where k is greater than the array length
    
    count := 0  // tracks the number of elements that have been moved
    start := 0  // starting index for each cycle
    
    for count < n {
        curr := start
        prev := nums[start]
        
        // Move each element to its correct position in the cycle
        for {

			next_idx := (curr + k) % n
            nums[next_idx], prev = prev, nums[next_idx]
            curr = next_idx
            count += 1
            
            //Break the cycle when we return to the starting position
            if curr == start {
				break
			}
		}
			
        start += 1  // Move to the next cycle
	}

	return nums
}