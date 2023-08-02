# Remove-duplicates-from-sorted-array

Example 1:

Input: nums = [1,1,2]
Output: 2, nums = [1,2,_]
Explanation: Your function should return k = 2, with the first two elements of nums being 1 and 2 respectively.
It does not matter what you leave beyond the returned k (hence they are underscores).
Example 2:

Input: nums = [0,0,1,1,1,2,2,3,3,4]
Output: 5, nums = [0,1,2,3,4,_,_,_,_,_]
Explanation: Your function should return k = 5, with the first five elements of nums being 0, 1, 2, 3, and 4 respectively.
It does not matter what you leave beyond the returned k (hence they are underscores).

Custom Judge:

The judge will test your solution with the following code:

int[] nums = [...]; // Input array
int[] expectedNums = [...]; // The expected answer with correct length

int k = removeDuplicates(nums); // Calls your implementation

assert k == expectedNums.length;
for (int i = 0; i < k; i++) {
    assert nums[i] == expectedNums[i];
}

If all assertions pass, then your solution will be accepted.


Given an integer array nums sorted in non-decreasing order, remove the duplicates in-place such that each unique element appears only once. The relative order of the elements should be kept the same. Then return the number of unique elements in nums.

Consider the number of unique elements of nums to be k, to get accepted, you need to do the following things:

Change the array nums such that the first k elements of nums contain the unique elements in the order they were present in nums initially. The remaining elements of nums are not important as well as the size of nums.
Return k.



map[]

1
[1 1 2] 0
map[1:true]
1
[1 1 2] 0
map[1:true]
2
[1 1 2] 1
[1]



[1,1,2] ==> [1,2,2]
index = 0 | val = 1
0 | 1 => cek_duplicate(false) => [1,...] index = 1
1 | 1 => cek_duplicate(true) => [1,...] index = 1
2 | 2 => cek_duplicate(false) => 



# Remove-duplicates-from-sorted-array II

Example 1:

Input: nums = [1,1,1,2,2,3]
Output: 5, nums = [1,1,2,2,3,_]
Explanation: Your function should return k = 5, with the first five elements of nums being 1, 1, 2, 2 and 3 respectively.
It does not matter what you leave beyond the returned k (hence they are underscores).
Example 2:

Input: nums = [0,0,1,1,1,1,2,3,3]
Output: 7, nums = [0,0,1,1,2,3,3,_,_]
Explanation: Your function should return k = 7, with the first seven elements of nums being 0, 0, 1, 1, 2, 3 and 3 respectively.
It does not matter what you leave beyond the returned k (hence they are underscores).

In this example, we define the majorityElement function that takes an array nums as input and returns the majority element. We initialize candidate to the first element of the array and count to 1. Then, we iterate over the remaining elements of the array and perform the following steps:

- If the count is 0, we update the candidate to the current element and reset the count to 1.
- If the candidate is equal to the current element, we increment the count.
- If the candidate is different from the current element, we decrement the count.

Given an integer array nums sorted in non-decreasing order, remove some duplicates in-place such that each unique element appears at most twice. The relative order of the elements should be kept the same.

Since it is impossible to change the length of the array in some languages, you must instead have the result be placed in the first part of the array nums. More formally, if there are k elements after removing the duplicates, then the first k elements of nums should hold the final result. It does not matter what you leave beyond the first k elements.

Return k after placing the final result in the first k slots of nums.

Do not allocate extra space for another array. You must do this by modifying the input array in-place with O(1) extra memory.

Custom Judge:

The judge will test your solution with the following code:

int[] nums = [...]; // Input array
int[] expectedNums = [...]; // The expected answer with correct length

int k = removeDuplicates(nums); // Calls your implementation

assert k == expectedNums.length;
for (int i = 0; i < k; i++) {
    assert nums[i] == expectedNums[i];
}
If all assertions pass, then your solution will be accepted.


nums [1 1 1 2 2 2 3]
loop on len(nums)
k : 0  |  updated nums: [1 1 1 2 2 2 3]  |  num 1  |  nums[j] = 1
nums [1 1 1 2 2 2 3]
1  |  [1 1 1 2 2 2 3]  |  1  |  1
[1 1 1 2 2 2 3]
2  |  false  |  [1 1 1 2 2 2 3]  |  1  |  1 false
[1 1 1 2 2 2 3]
2  |  true  |  [1 1 1 2 2 2 3]  |  2  |  1 true
[1 1 2 2 2 2 3]
3  |  true  |  [1 1 2 2 2 2 3]  |  2  |  1 true
[1 1 2 2 2 2 3]
4  |  false  |  [1 1 2 2 2 2 3]  |  2  |  2 false
[1 1 2 2 2 2 3]
4  |  true  |  [1 1 2 2 2 2 3]  |  3  |  2 true
[1 1 2 2 3 2 3]

++++++++++++++++++++++++++++++++++++++++++++++++
nums:  [1 1 1 2 2 2 3]  |  j > 2:  true
j:  0  |  nums:  [1 1 1 2 2 2 3]  |  num:  1  |  nums[j]:  1 (true) => j++ = 1
nums:  [1 1 1 2 2 2 3]  |  j > 2:  true
j:  1  |  nums:  [1 1 1 2 2 2 3]  |  num:  1  |  nums[j]:  1 (true) => j++ = 2

nums:  [1 1 1 2 2 2 3]  |  j > 2:  false
j:  2  |  nums:  [1 1 1 2 2 2 3]  |  num > nums[j-2]:  (false) => continue/do nothing  |   num:  1  |   |  nums[j-2]:  1  |  num > nums[j-2]:  false

nums:  [1 1 1 2 2 2 3]  |  j > 2:  false
j:  2  |  nums:  [1 1 1 2 2 2 3]  |  num > nums[j-2]:  (true) => nums[j] = num  |   num:  2  |   |  nums[j-2]:  1  |  num > nums[j-2]:  (true) j++ = 3

nums:  [1 1 2 2 2 2 3]  |  j > 2:  false
j:  3  |  nums:  [1 1 2 2 2 2 3]  |  num > nums[j-2]:  (true) => num[j] = num  |   num:  2  |   |  nums[j-2]:  1  |  num > nums[j-2]:  (true) j++ = 4

nums:  [1 1 2 2 2 2 3]  |  j > 2:  false
j:  4  |  nums:  [1 1 2 2 2 2 3]  |  num > nums[j-2]:  (false) => continue/do nothing   |   num:  2  |   |  nums[j-2]:  2  |  num > nums[j-2]:  false
nums:  [1 1 2 2 2 2 3]  |  j > 2:  false
j:  4  |  nums:  [1 1 2 2 2 2 3]  |  num > nums[j-2]:  (true) => num[j] = num  |   num:  3  |   |  nums[j-2]:  2  |  num > nums[j-2]:  (true) j++ = 5
[1 1 2 2 3 2 3]
--- PASS: TestRemoveDuplicate (0.00s)
    --- PASS: TestRemoveDuplicate/[1_1_1_2_2_2_3]_5 (0.00s)
PASS
ok      top-interview-150/internal/remove-duplicates-from-sorted-array  0.003s