package main

func removeStars(s string) string {
	stack := make([]rune, 0)
	for _, ch := range s {
		if ch != '*' {
			stack = append(stack, ch)
		} else {
			stack = stack[:len(stack)-1]
		}
	}
	return string(stack)
}
func main() {

}
