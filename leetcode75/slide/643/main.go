package main

import "fmt"

func findMaxAverage(nums []int, k int) float64 {
	sum := 0
	for i := 0; i < k; i++ {
		sum += nums[i]
	}
	ans := sum
	for i := k; i < len(nums); i++ {
		sum += nums[i]
		sum -= nums[i-k]
		ans = max(ans, sum)
	}
	return float64(ans) / float64(k)
}
func main() {
	nums := []int{1, 12, -5, -6, 50, 3}
	k := 4
	fmt.Println(findMaxAverage(nums, k))

}
