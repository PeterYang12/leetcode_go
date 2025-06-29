package main

import "fmt"

func letterCombinations(digits string) []string {
	ans := []string{}
	if len(digits) == 0 {
		return ans
	}
	hashmap := map[byte]string{'1': "", '2': "abc", '3': "def", '4': "ghi", '5': "jkl", '6': "mno", '7': "pqrs", '8': "tuv", '9': "wxyz"}

	var dfs func(index int, path string)
	dfs = func(index int, path string) {
		if index == len(digits) {
			ans = append(ans, path)
			return
		}
		for _, ch := range hashmap[digits[index]] {
			dfs(index+1, path+string(ch))
		}
	}
	dfs(0, "")
	return ans
}
func main() {
	digits := "23"
	fmt.Println(letterCombinations(digits))
}
