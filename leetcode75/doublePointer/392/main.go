package main

func isSubsequence(s string, t string) bool {
	n1 := len(s)
	n2 := len(t)
	i, j := 0, 0
	for i < n1 && j < n2 {
		if s[i] == t[j] {
			i++
		}
		j++
	}
	return i == n1
}
func main() {

}
