package main

import (
	"fmt"
)

func main() {
	fmt.Print(stringCompression("aabbcccdaa"))
}

func stringCompression(s string) string {
	runeS := []rune(s)
	var res []rune
	count := 0
	leftIndex := 0
	for i := 0; i < len(s); i++ {
		if runeS[leftIndex] == runeS[i] {
			count++
		} else {
			res = append(res, runeS[leftIndex])
			if count > 1 {
				res = append(res, rune(count+48))
			}
			leftIndex = i
			count = 1
		}
	}
	res = append(res, runeS[leftIndex])
	if count > 1 {
		res = append(res, rune(count+48))
	}
	return string(res)
}
