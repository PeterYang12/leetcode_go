package main

import "fmt"

func orangesRotting(grid [][]int) int {
	var queue [][]int

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 2 {
				queue = append(queue, []int{i, j})
			}
		}
	}
	ans := 0
	dir := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for len(queue) > 0 {
		var new_queue [][]int
		for i := 0; i < len(queue); i++ {
			x, y := queue[i][0], queue[i][1]
			for _, d := range dir {
				new_x := x + d[0]
				new_y := y + d[1]
				if new_x >= 0 && new_x < len(grid) && new_y >= 0 && new_y < len(grid[0]) && grid[new_x][new_y] == 1 {
					grid[new_x][new_y] = 2
					new_queue = append(new_queue, []int{new_x, new_y})
				}
			}
		}
		if len(new_queue) == 0 {
			break
		}
		ans++
		queue = new_queue
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				return -1
			}
		}
	}
	return ans
}
func main() {
	grid := [][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}}
	fmt.Println(orangesRotting(grid))
}
