package main

func increasingTriplet(nums []int) bool {
	n := len(nums)
	if n < 3 {
		return false
	}
	leftMin := make([]int, n)
	rightMax := make([]int, n)
	leftMin[0] = nums[0]
	for i := 1; i < n; i++ {
		leftMin[i] = min(leftMin[i-1], nums[i])
	}
	rightMax[n-1] = nums[n-1]
	for i := n - 2; i >= 0; i-- {
		rightMax[i] = max(rightMax[i+1], nums[i])
	}
	for i := 1; i < n-1; i++ {
		if nums[i] > leftMin[i] && nums[i] < rightMax[i] {
			return true
		}
	}
	return false
}
func main() {

}
