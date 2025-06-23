package main

import (
	"fmt"
	"strings"
)

func maxVowels(s string, k int) int {
	str := "aeiou"
	ans := 0
	cnt := 0
	for i := 0; i < k; i++ {
		if strings.Contains(str, string(s[i])) {
			cnt++
		}
	}
	ans = max(ans, cnt)
	for i := k; i < len(s); i++ {
		if strings.Contains(str, string(s[i])) {
			cnt++
		}
		if strings.Contains(str, string(s[i-k])) {
			cnt--
		}
		ans = max(ans, cnt)
	}
	return ans
}
func main() {
	str := "abciiidef"
	fmt.Println(maxVowels(str, 3))

}
