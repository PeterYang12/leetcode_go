package main

import "fmt"

func combinationSum2(k int, n int) [][]int {
	ans := make([][]int, 0)
	var dfs func(id int, currentSum int, path []int)
	dfs = func(id int, currentSum int, path []int) {
		if len(path) == k {
			if currentSum == n {
				// copy result to avoid shared slice
				ans = append(ans, append([]int(nil), path...))
			}
			return
		}
		for i := id; i <= 9; i++ {
			s := currentSum + i
			if s > n {
				break
			}
			// create new slice
			newPath := append([]int{}, path...)
			newPath = append(newPath, i)
			dfs(i+1, s, newPath)
		}
	}
	dfs(1, 0, []int{})
	return ans
}
func combinationSum3(k int, n int) [][]int {
	var ans [][]int
	var path []int
	var dfs func(start, sum int)
	dfs = func(start, sum int) {
		if len(path) == k {
			if sum == n {
				ans = append(ans, append([]int(nil), path...))
			}
			return
		}

		for i := start; i <= 9; i++ {
			if sum+i > n {
				break
			}
			path = append(path, i)
			dfs(i+1, sum+i)
			path = path[:len(path)-1]
		}
	}
	dfs(1, 0)
	return ans
}
func main() {
	fmt.Println(combinationSum3(3, 7))

}
