package main

import (
	"fmt"
	"sort"
)

func containsDuplicate(nums []int) bool {
	sort.Ints(nums)

	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			return true
		}
	}

	return false
}

func main() {
	var n int

	// number of elements
	fmt.Scan(&n)

	nums := make([]int, n)

	// read array elements
	for i := 0; i < n; i++ {
		fmt.Scan(&nums[i])
	}

	fmt.Println(containsDuplicate(nums))
}
