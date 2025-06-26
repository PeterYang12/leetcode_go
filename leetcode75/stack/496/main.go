package main

import "fmt"

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	var stack []int
	hashmap := make(map[int]int)
	for i := len(nums2) - 1; i >= 0; i-- {
		for len(stack) > 0 && stack[len(stack)-1] < nums2[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			hashmap[nums2[i]] = stack[len(stack)-1]
		} else {
			hashmap[nums2[i]] = -1
		}
		stack = append(stack, nums2[i])
	}
	ans := make([]int, len(nums1))
	for idx, num := range nums1 {
		if _, ok := hashmap[num]; ok {
			ans[idx] = hashmap[num]
		} else {
			ans[idx] = -1
		}
	}
	return ans
}
func main() {
	nums1 := []int{4, 1, 2}
	nums2 := []int{1, 3, 4, 2}
	fmt.Println(nextGreaterElement(nums1, nums2))
}
