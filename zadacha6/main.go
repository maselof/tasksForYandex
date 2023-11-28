package main

import "fmt"

func main() {
	fmt.Print(pallidrom("abcbbca"))
}

func pallidrom(s string) bool {
	runeS := []rune(s)
	leftIndex := 0
	rightIndex := len(runeS) - 1
	res := true
	count := 0
	for leftIndex < rightIndex && count <= 1 {
		if runeS[leftIndex] != runeS[rightIndex] {
			if runeS[leftIndex+1] == runeS[rightIndex] {
				leftIndex++
				count++
			} else if runeS[leftIndex] == runeS[rightIndex-1] {
				rightIndex--
				count++
			} else {
				res = false
				break
			}
		}
		leftIndex++
		rightIndex--
	}
	if count > 1 {
		res = false
	}
	return res
}
