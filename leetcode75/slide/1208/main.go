package main

import "fmt"

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
func equalSubstring(s string, t string, maxCost int) int {
	left, right := 0, 0
	ans := 0
	cost := 0
	for ; right < len(s); right++ {
		cost += abs(int(s[right]) - int(t[right]))
		for cost > maxCost {
			cost -= abs(int(s[left]) - int(t[left]))
			left++
		}
		ans = max(ans, right-left+1)
	}
	return ans
}
func main() {
	s1 := "abcd"
	s2 := "cdef"
	fmt.Println(equalSubstring(s1, s2, 3))

}
