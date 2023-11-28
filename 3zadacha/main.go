package main

import "fmt"

func main() {
	fmt.Print(stringsEquality("babdc", "abc"))
}

func stringsEquality(s, p string) (int, int) {
	runeS := []rune(s)
	runeP := []rune(p)
	helperMap := make(map[rune]int)
	for _, val := range runeP {
		helperMap[val]++
	}
	leftIndex := 0
	x, y := 0, 0
	count := 0
	for i := 0; i < len(runeS); i++ {
		if helperMap[runeS[i]] != 0 {
			count++
			helperMap[runeS[i]]--
			if count == len(runeP) {
				x = leftIndex
				y = i
			}
		} else {
			_, ok := helperMap[runeS[i]]
			if ok {
				leftIndex++
			} else {
				for count != 0 {
					helperMap[runeS[leftIndex]]++
					leftIndex++
					count--
				}
				leftIndex++
			}
		}
	}
	return x, y
}
