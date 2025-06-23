package main

import (
	"fmt"
	"strconv"
)

func compress(chars []byte) int {
	l, r := 0, 0
	write := 0
	for r < len(chars) {
		if r == len(chars)-1 || chars[l] != chars[r+1] {
			chars[write] = chars[l]
			write++
			n := r - l + 1
			if n > 1 {
				str := strconv.Itoa(n)
				for i := 0; i < len(str); i++ {
					chars[write] = str[i]
					write++
				}
			}
			l = r + 1
		}
		r++
	}
	return write
}
func main() {
	chars := []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}
	//chars := []byte{'a', 'b', 'c'}
	ans := compress(chars)
	fmt.Println(ans)
	for i := 0; i < len(chars); i++ {
		fmt.Printf("%c", chars[i])
	}
}
