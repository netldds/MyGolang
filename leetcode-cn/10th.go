package main

import "fmt"

const (
	J      = "aA"
	S      = "aAAbbbb"
	LOVELY = "LOVELY"
)

func numJewelsInStones(J string, S string) int {
	count := 0
	if len(S) > 1 {
		count += numJewelsInStones(J, S[1:])
	}
	for i, _ := range J {
		if J[i] == S[0] {
			count++
		}
	}
	return count
}
func toLowerCase(str string) string {
	var chars []rune
	if len(str) > 1 {
		chars = append(chars, []rune(toLowerCase(str[:len(str)-1]))...)
	}
	if str[len(str)-1] <= 90 && str[len(str)-1] >= 65 {
		chars = append(chars, rune(str[len(str)-1]+32))
	} else {
		chars = append(chars, rune(str[len(str)-1]))
	}
	return string(chars)
}
func main() {
	//fmt.Println(numJewelsInStones(J, S))
	fmt.Println(toLowerCase(LOVELY))
}
