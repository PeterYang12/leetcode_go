package main

import "fmt"

func canPlaceFlowers(flowerbed []int, n int) bool {
	if n == 0 {
		return true
	}
	length := len(flowerbed)
	for i := 0; i < length; i++ {
		if flowerbed[i] == 0 && (i == 0 || flowerbed[i-1] == 0) && (i == length-1 || flowerbed[i+1] == 0) {
			flowerbed[i] = 1
			n--
		}
		if n == 0 {
			return true
		}
	}
	return false
}
func main() {
	flowers := []int{0, 0, 0, 0, 0, 1, 0, 0}
	ans := canPlaceFlowers(flowers, 0)
	fmt.Println(ans)
}
