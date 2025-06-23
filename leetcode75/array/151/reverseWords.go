package main

import "fmt"

// 双指针
func reverseWords(s string) string {
	ans := ""
	l := len(s) - 1
	// 去除尾部的空格
	for l >= 0 && s[l] == ' ' {
		l--
	}
	r := l
	for l >= 0 {
		for l >= 0 && s[l] != ' ' {
			l--
		}
		ans += s[l+1:r+1] + " "
		for l >= 0 && s[l] == ' ' {
			l--
		}
		r = l
	}

	return ans[:len(ans)-1]
}
func main() {
	fmt.Println(reverseWords("the sky is blue"))
}
