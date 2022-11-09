package main

import (
	binarysearch "binarysearch-1/internal/binary_search"
	filteritem "binarysearch-1/internal/filter_item"
	findsmallestlettergreaterthantarget "binarysearch-1/internal/find_smallest_letter_greater_than_target"
	"fmt"
)

func main()  {
	arr := [][]int{{5,2},{4,2},{7,7},{8,3},{4,5}}
	query := [][]int{{5,5,2,2,},{4,7,5,7},{4,8,2, 7}}
	fmt.Println(filteritem.FilterItem(5,arr,3,query))
	
	// fmt.Println(findsmallestlettergreaterthantarget.FindSmallestLetter([]string{"x","x","y","y"},"z"))
	fmt.Println(findsmallestlettergreaterthantarget.FindSmallestLetter([]string{"c","f","j"},"c"))

	fmt.Println(binarysearch.BinarySearch([]int{-1,0,3,5,9,12}, 9))
}