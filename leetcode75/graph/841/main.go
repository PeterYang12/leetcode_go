package main

import (
	"fmt"
)

func canVisitAllRooms1(rooms [][]int) bool {
	n := len(rooms)
	visited := make([]bool, n)
	visited[0] = true
	queue := []int{0}
	cnt := 0
	for len(queue) > 0 {
		m := len(queue)
		for i := 0; i < m; i++ {
			idx := queue[i]
			cnt++
			for _, num := range rooms[idx] {
				if !visited[num] {
					visited[num] = true
					queue = append(queue, num)
				}
			}
		}
		queue = queue[m:]
	}
	fmt.Println(cnt)
	return cnt == n
}
func canVisitAllRooms2(rooms [][]int) bool {
	n := len(rooms)
	visited := make([]bool, n)
	visited[0] = true
	queue := []int{0}
	cnt := 0
	for len(queue) > 0 {
		idx := queue[0]
		queue = queue[1:]
		cnt++
		for _, num := range rooms[idx] {
			if !visited[num] {
				visited[num] = true
				queue = append(queue, num)
			}
		}
	}
	return cnt == n
}

func canVisitAllRooms(rooms [][]int) bool {
	cnt := 0
	n := len(rooms)
	visited := make([]bool, n)
	var dfs func(rooms [][]int, i int)
	dfs = func(rooms [][]int, i int) {
		cnt++
		visited[i] = true
		for _, num := range rooms[i] {
			if !visited[num] {
				dfs(rooms, num)
			}
		}
		return
	}
	dfs(rooms, 0)
	return cnt == n
}
func main() {
	//rooms := [][]int{{1}, {2}, {3}, {}}
	//rooms := [][]int{{1}, {1}}
	rooms := [][]int{{1, 3}, {3, 0, 1}, {2, 0}, {0}}
	fmt.Println(canVisitAllRooms(rooms))

}
