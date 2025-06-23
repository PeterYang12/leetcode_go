package main

import "fmt"

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	ans := 0
	for left < right {
		if height[left] < height[right] {
			ans = max(ans, (right-left)*height[left])
			left++
		} else {
			ans = max(ans, (right-left)*height[right])
			right--
		}
	}
	return ans
}
func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	ans := maxArea(height)
	fmt.Println(ans)
}
