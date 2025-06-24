package main

import "fmt"

func findDifference(nums1 []int, nums2 []int) [][]int {
	hash1 := make(map[int]struct{})
	hash2 := make(map[int]struct{})
	for i := 0; i < len(nums1); i++ {
		hash1[nums1[i]] = struct{}{}
	}
	for i := 0; i < len(nums2); i++ {
		hash2[nums2[i]] = struct{}{}
	}
	ans := [][]int{{}, {}}
	for k, _ := range hash1 {
		if _, ok := hash2[k]; !ok {
			ans[0] = append(ans[0], k)
		}
	}
	for k, _ := range hash2 {
		if _, ok := hash1[k]; !ok {
			ans[1] = append(ans[1], k)
		}
	}
	return ans
}
func main() {
	nums1 := []int{1, 2, 3, 3}
	nums2 := []int{1, 1, 2, 2}
	fmt.Println(findDifference(nums1, nums2))
}
