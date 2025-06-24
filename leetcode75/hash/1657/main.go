package main

import "sort"

func closeStrings(word1 string, word2 string) bool {
	cnt1 := make([]int, 26)
	cnt2 := make([]int, 26)
	for _, ch := range word1 {
		cnt1[ch-'a']++
	}
	for _, ch := range word2 {
		cnt2[ch-'a']++
	}
	for i := 0; i < len(cnt1); i++ {
		if cnt1[i]+cnt2[i] == 0 {
			continue
		} else if cnt1[i] == 0 || cnt2[i] == 0 {
			return false
		}
	}
	sort.Ints(cnt1)
	sort.Ints(cnt2)
	for i := 0; i < len(cnt1); i++ {
		if cnt1[i] != cnt2[i] {
			return false
		}
	}
	return true
}
func main() {

}
