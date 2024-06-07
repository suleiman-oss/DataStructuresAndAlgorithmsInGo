package main

import "fmt"

func CountPairs(s string, a rune, g rune) int {
	ac := 0
	ans := 0
	for _, v := range s {
		if v == a {
			ac++
		}
		if v == g {
			ans += ac
		}
	}
	return ans
}

func main() {
	a := 'a'
	g := 'g'
	s := "bcaggaag"
	fmt.Println(CountPairs(s, a, g))
}
