package main

import "fmt"

func predictPartyVictory(senate string) string {
	var r, d []int
	for i, ch := range senate {
		if ch == 'R' {
			r = append(r, i)
		} else if ch == 'D' {
			d = append(d, i)
		}
	}
	for len(r) > 0 && len(d) > 0 {
		if r[0] < d[0] {
			r = append(r, r[0]+len(senate))
		} else {
			d = append(d, d[0]+len(senate))
		}
		r = r[1:]
		d = d[1:]
	}
	if len(r) > 0 {
		return "Radiant"
	}
	return "Dire"

}
func main() {
	s := "RDD"
	fmt.Println(predictPartyVictory(s))

}
