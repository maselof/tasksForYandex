package main

import "fmt"

func main() {
	fmt.Println(distinctAlphabet("adasabc", "abcd"))
}

func distinctAlphabet(s, p string) (int, int) {
	runeS := []rune(s)
	runeP := []rune(p)
	helperMap := make(map[rune]int)
	for _, char := range runeP {
		helperMap[char] = -1
	}
	leftIndex := 0
	x, y := 0, 0
	count := 0
	minDiff := len(runeS)
	for i := 0; i < len(runeS); i++ {
		_, ok := helperMap[runeS[i]]
		if ok {
			if helperMap[runeS[i]] == -1 {
				count++
			}
			helperMap[runeS[i]] = i
		}
		if count >= len(helperMap) {
			for helperMap[runeS[leftIndex]] != leftIndex {
				leftIndex++
			}
			if i-leftIndex+1 <= minDiff {
				minDiff = i - leftIndex + 1
				x = leftIndex
				y = i
			}
		}
	}
	return x, y
}
