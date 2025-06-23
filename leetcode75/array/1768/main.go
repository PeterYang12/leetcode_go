package main

import "fmt"

func mergeAlternately(word1 string, word2 string) string {
	n1 := len(word1)
	n2 := len(word2)
	ans := make([]byte, n1+n2)
	l1, l2 := 0, 0
	for i := 0; i < n1+n2; {
		if l1 < n1 {
			ans[i] = word1[l1]
			l1 += 1
			i++
		}
		if l2 < n2 {
			ans[i] = word2[l2]
			l2++
			i++
		}
	}
	return string(ans)
}

func main() {
	a := "abc"
	b := "pqr"
	ans := mergeAlternately(a, b)
	fmt.Printf(ans)
}
