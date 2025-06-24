package main

import (
	"fmt"
	"strconv"
)

func equalPairs(grid [][]int) int {
	ans := 0
	hash := make(map[string]int)
	n := len(grid)
	for i := 0; i < n; i++ {
		str := ""
		for j := 0; j < n; j++ {
			str += strconv.Itoa(grid[i][j]) + "-"
		}
		hash[str]++
	}
	for j := 0; j < n; j++ {
		str := ""
		for i := 0; i < n; i++ {
			str += strconv.Itoa(grid[i][j]) + "-"
		}
		if v, ok := hash[str]; ok {
			ans += v
		}
	}
	return ans
}
func main() {
	//grid := [][]int{{3, 1, 2, 2}, {1, 4, 4, 5}, {2, 4, 2, 2}, {2, 4, 2, 2}}
	grid := [][]int{{11, 1}, {1, 11}}
	ans := equalPairs(grid)
	fmt.Println(ans)
}
