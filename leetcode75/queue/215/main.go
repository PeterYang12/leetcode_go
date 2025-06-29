package main

import (
	"fmt"
	"math/rand"
)

func findKthLargest1(nums []int, k int) int {
	return quickSelect(nums, 0, len(nums)-1, len(nums)-k)
}

func quickSelect(nums []int, left, right, k int) int {
	if left == right {
		return nums[left]
	}

	// 随机pivot
	pivotIndex := left + rand.Intn(right-left+1)
	nums[pivotIndex], nums[right] = nums[right], nums[pivotIndex]
	pivot := nums[right]
	i := left
	for j := left; j < right; j++ {
		if nums[j] < pivot {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}

	nums[i], nums[right] = nums[right], nums[i]
	pivotIndex = i

	if pivotIndex == k {
		return nums[pivotIndex]
	} else if pivotIndex < k {
		return quickSelect(nums, pivotIndex+1, right, k)
	} else {
		return quickSelect(nums, left, pivotIndex-1, k)
	}
}
func main() {
	//nums := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	nums := []int{-1, 2, 0}
	fmt.Println(findKthLargest(nums, 3))
}
