package main

import "fmt"

func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	// 2 <= cost.length <= 1000
	currCost := 0
	preCost := 0
	for i := 2; i <= n; i++ {
		tmp := min(currCost+cost[i-1], preCost+cost[i-2])
		preCost = currCost
		currCost = tmp
	}
	return currCost
}

func main() {
	//cost := []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}
	cost := []int{10, 15, 20}
	fmt.Println(minCostClimbingStairs(cost))
}
