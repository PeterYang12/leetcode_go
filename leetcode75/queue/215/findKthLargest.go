package main

import (
	"math/rand"
)

func findKthLargest(nums []int, k int) int {
	n := len(nums)
	return selectSort(nums, 0, n-1, k-1)
}

func selectSort(nums []int, left, right, k int) int {
	if left == right {
		return nums[left]
	}
	pivotIndex := rand.Intn(right-left+1) + left
	nums[right], nums[pivotIndex] = nums[pivotIndex], nums[right]
	pivot := nums[right]
	cnt := left
	for i := left; i < right; i++ {
		if nums[i] > pivot {
			nums[i], nums[cnt] = nums[cnt], nums[i]
			cnt++
		}
	}
	pivotIndex = cnt
	nums[right], nums[pivotIndex] = nums[pivotIndex], nums[right]
	if pivotIndex == k {
		return nums[pivotIndex]
	} else if pivotIndex > k {
		return selectSort(nums, left, pivotIndex-1, k)
	} else {
		return selectSort(nums, pivotIndex+1, right, k)
	}
}
