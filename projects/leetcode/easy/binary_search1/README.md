# Binary Search

In computer science, binary search, also known as half-interval search or logarithmic search, is a search algorithm that is commonly used to find the position of a target value within a sorted array. However, the binary search algorithm is not limited to arrays. In any scenario where one can determine if the target value is higher or lower than a given value, binary search can be used to repeatedly halve the search space for a target value.


## Introduction
In this card, we are going to help you understand the general concept of Binary Search.

What is Binary Search?
Binary Search is one of the most fundamental and useful algorithms in Computer Science. It describes the process of searching for a specific value in an ordered collection.

Terminology used in Binary Search:

Target - the value that you are searching for
Index - the current location that you are searching
Left, Right - the indicies from which we use to maintain our search Space
Mid - the index that we use to apply a condition to determine if we should search left or right
Other Binary Search Defintions:

[Wikipedia](https://en.wikipedia.org/wiki/Binary_search_algorithm)

## Background

How does it work?
In its simplest form, Binary Search operates on a contiguous sequence with a specified left and right index. This is called the Search Space. Binary Search maintains the left, right, and middle indicies of the search space and compares the search target or applies the search condition to the middle value of the collection; if the condition is unsatisfied or values unequal, the half in which the target cannot lie is eliminated and the search continues on the remaining half until it is successful. If the search ends with an empty half, the condition cannot be fulfilled and target is not found.

In the following chapters, we will review how to identify Binary Search problems, reasons why we use Binary Search, and the 3 different Binary Search templates that you might be previously unaware of. Since Binary Search is a common interview topic, we will also categorize practice problems to different templates so you can practice using each.

Note:

Binary Search can take many alternate forms and might not always be as straight forward as searching for a specific value. Sometimes you will have to apply a specific condition or rule to determine which side (left or right) to search next.
We will provide examples in the coming chapters. First, could you try write a binary search algorithm yourself?



### Do it yourself (template - 1) 
Description or explanation on Indonesia language

arr = {-1,0,3,5,9,12},
target = 9

left:  0
right:  6

mid:  3 ==> 5 (value)
mid < target ==> 5 < 9 ==> true 
setelah dibagi dua, atau dilihat dari sudut pandang tengah tengah atau midle,
untuk menentukan target diarah kanan (mid masih kerendahan atau kurang) atau target disebelah kiri, atau ketinggiian

Contoh:
mid = 3,
value pada array dengan index ke 3 adalah 5, 
kemudian dilakukan pengecekan / pembandingan apakah 5 lebih kecil(<) dari 9 (target)
maka, tambahkan mid dengan 1 (artinya index yang dicari masih kekecilan) ada disisi kanan 
left = mid + 1 

left sekarang bernilai 4
kemudian dicari kembali nilai mid dengan rumus
mid = left + (right - left) / 2
left:  4
mid = 4 + (6 - 4) / 2
mid = 5
value pada array dengan index ke 5 (mid) adalah 12,
kemudian dilakukan pengecekan / pembandingan apakah 12 lebih kecil(<) dari 9 (target) ==> false
masuk ke else blok, maka akan ada rumus perhitungan right = mid - 1, right = 5 - 1 = 4
maka, kurangkan mid dengan 1 (artinya index yang dicari masih kelebihan atau di seblah kiri) mundur dikit
right:  4
right sekarang bernilai 4
mid = left + (right - left) / 2
mid = 4 + (4-4) / 2
mid = 4

value pada array dengan index ke 4(mid) adalah 9,
kemudian dilakukan pengecekan / pembandingan apakah 9 sama dengan 9 (target) ==> true
maka return index tersebut (mid)
```go
func binarySeacrh(nums []int, target int) int {
	left, right := 0, len(nums)
	fmt.Println("left: ", left)
	fmt.Println("right: ", right)
	for i := 0; i < right; i++ {
		mid := left + (right-left)/2
		fmt.Println("mid: ", mid)
		
		if nums[mid] == target {
			return mid
		}
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
		fmt.Println("left: ", left)
		fmt.Println("right: ", right)
	}
	return -1
}

```