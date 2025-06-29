package main

import (
	"fmt"
	"math/rand"
)

func quickSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	pivotIndex := rand.Intn(len(nums))
	nums[pivotIndex], nums[len(nums)-1] = nums[len(nums)-1], nums[pivotIndex]
	pivot := nums[len(nums)-1]
	left := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] < pivot {
			nums[left], nums[i] = nums[i], nums[left]
			left++
		}
	}
	nums[left], nums[len(nums)-1] = nums[len(nums)-1], nums[left]

	quickSort(nums[:left])
	quickSort(nums[left+1:])
	return nums
}
func main() {
	// 测试快速排序
	nums := []int{10, 7, 8, 9, 1, 5, 3, 6, 2, 4}
	fmt.Println("排序前:", nums)

	quickSort(nums)
	fmt.Println("排序后:", nums)

}
