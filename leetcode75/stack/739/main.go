package main

import "fmt"

func dailyTemperatures(temperatures []int) []int {
	stack := make([]int, 0)
	n := len(temperatures)
	ans := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		for len(stack) != 0 && temperatures[stack[len(stack)-1]] <= temperatures[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) != 0 {
			ans[i] = stack[len(stack)-1] - i
		}
		stack = append(stack, i)
	}
	return ans
}
func main() {
	nums := []int{89, 62, 70, 58, 47, 47, 46, 76, 100, 70}
	fmt.Println(dailyTemperatures(nums))
}
