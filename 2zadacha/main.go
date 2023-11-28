package main

import (
	"fmt"
)

func main() {
	fmt.Print(minSuffix("bbbbadasa", 3))
}

func minSuffix(s string, k int) (int, int) {
	runeS := []rune(s)
	leftIndex := 0
	helperMap := make(map[rune]int)
	minSuff := len(runeS)
	x, y := 0, len(runeS)
	for i := 0; i < len(runeS); i++ {
		helperMap[runeS[i]] = i
		if len(helperMap) >= k {
			if len(helperMap) > k {
				delete(helperMap, runeS[leftIndex])
				leftIndex++
				for helperMap[runeS[leftIndex]] != leftIndex {
					leftIndex++
				}
			}
			if i-leftIndex+1 < minSuff {
				minSuff = i - leftIndex + 1
				x = leftIndex
				y = i
			}
		}
	}
	return x, y
}
