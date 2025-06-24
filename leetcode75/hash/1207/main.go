package main

import "fmt"

func uniqueOccurrences(arr []int) bool {
	hash := make(map[int]int)
	for _, v := range arr {
		hash[v]++
	}
	set := make(map[int]struct{})
	for _, cnt := range hash {
		if _, ok := set[cnt]; ok {
			return false
		}
		set[cnt] = struct{}{}
	}
	return true
}
func main() {
	arr := []int{1, 2, 2, 1, 1, 3}
	fmt.Println(uniqueOccurrences(arr))
}
