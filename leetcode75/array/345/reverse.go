package main

import (
	"fmt"
	"strings"
)

func reverseVowels(s string) string {
	t := []byte(s)
	n := len(t)
	i, j := 0, len(s)-1
	for i < j {
		for i < n && !strings.Contains("aeiouAEIOU", string(t[i])) {
			i++
		}
		for j >= 0 && !strings.Contains("aeiouAEIOU", string(t[j])) {
			j--
		}
		if i < j {
			t[i], t[j] = t[j], t[i]
			i++
			j--
		}
	}
	return string(t)
}
func main() {
	s := "IceCreAm"
	fmt.Println(reverseVowels(s))
}
