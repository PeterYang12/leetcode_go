package main

import "fmt"

func pivotIndex(nums []int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	sum := 0
	for i := 0; i < len(nums); i++ {
		if sum*2+nums[i] == total {
			return i
		}
		sum += nums[i]
	}
	return -1
}
func main() {
	nums := []int{1, 7, 3, 6, 5, 6}
	fmt.Println(pivotIndex(nums))
}
