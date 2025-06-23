package main

import (
	"fmt"
	"sort"
)

func maxOperations(nums []int, k int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	ans := 0
	l, r := 0, len(nums)-1
	for l < r {
		if nums[l]+nums[r] == k {
			ans++
			l++
			r--
		} else if nums[l]+nums[r] < k {
			l++
		} else {
			r--
		}
	}
	return ans
}

func main() {
	nums := []int{3, 1, 3, 4, 3}
	fmt.Println(maxOperations(nums, 6))

}
