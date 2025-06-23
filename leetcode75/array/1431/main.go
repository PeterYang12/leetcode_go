package main

import "fmt"

func kidsWithCandies(candies []int, extraCandies int) []bool {
	maxCandy := candies[0]
	for i := 0; i < len(candies); i++ {
		maxCandy = max(maxCandy, candies[i])
	}
	ans := make([]bool, len(candies))
	for i := 0; i < len(candies); i++ {
		if (candies[i] + extraCandies) >= maxCandy {
			ans[i] = true
		} else {
			ans[i] = false
		}
	}
	return ans
}
func main() {
	candies := []int{2, 3, 5, 1, 3}
	extraCandies := 3
	ans := kidsWithCandies(candies, extraCandies)
	fmt.Println(ans)
}
