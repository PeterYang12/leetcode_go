package main

import "fmt"

func productExceptSelf(nums []int) []int {
	n := len(nums)
	left := make([]int, n)
	right := make([]int, n)
	left[0] = 1
	right[n-1] = 1
	for i := 1; i < n; i++ {
		left[i] = left[i-1] * nums[i-1]
		right[n-1-i] = right[n-i] * nums[n-i]
	}
	ans := make([]int, n)
	for i := 0; i < n; i++ {
		ans[i] = left[i] * right[i]
	}
	return ans
}
func main() {
	nums := []int{1, 2, 3, 4}
	ans := productExceptSelf(nums)
	fmt.Println(ans)
}
