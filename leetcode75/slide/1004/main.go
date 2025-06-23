package main

import "fmt"

func longestOnes(nums []int, k int) int {
	left, right := 0, 0
	ans := 0
	cnt := 0
	for right < len(nums) {
		if nums[right] == 0 {
			cnt++
		}
		right++
		for cnt > k {
			if nums[left] == 0 {
				cnt--
			}
			left++
		}
		ans = max(ans, right-left)
	}
	return ans
}
func main() {
	nums := []int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}
	k := 2
	fmt.Println(longestOnes(nums, k))
}
