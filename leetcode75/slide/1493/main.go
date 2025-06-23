package main

import "fmt"

// 只有一个0组成的子数组中1的数量最多
func longestSubarray(nums []int) int {
	ans := 0
	left, right := 0, 0
	cnt0 := 0
	for right < len(nums) {
		cnt0 += 1 - nums[right]
		for cnt0 > 1 {
			cnt0 -= 1 - nums[left]
			left++
		}
		ans = max(ans, right-left)
		right++
	}
	return ans
}
func main() {
	nums := []int{1, 1, 1}
	fmt.Println(longestSubarray(nums))
}
