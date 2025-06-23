package main

import (
	"fmt"
	"strings"
)

func gcdOfStrings(str1 string, str2 string) string {
	l := min(len(str1), len(str2))
	for i := l; i > 0; i-- {
		if len(str1)%i == 0 && len(str2)%i == 0 {
			substr1 := str1[:i]
			if strings.Repeat(substr1, len(str1)/i) == str1 && strings.Repeat(substr1, len(str2)/i) == str2 {
				return substr1
			}
		}
	}
	return ""
}
func main() {
	str1 := "ABCABC"
	str2 := "ABC"
	fmt.Println(gcdOfStrings(str1, str2))
}
