package main

import "fmt"

func main() {
	fmt.Print(minSuffix(""))
}

func minSuffix(s string) (int, int) {
	leftIndex := 0
	helperMap := make(map[rune]int)
	sz := len([]rune(s))
	minSuffix := sz
	for _, val := range s {
		helperMap[val] = -1
	}
	count := 0
	x, y := 0, sz
	for i := 0; i < sz; i++ {
		if helperMap[rune(s[i])] == -1 {
			count++
		}
		helperMap[rune(s[i])] = i
		if count == len(helperMap) {
			for helperMap[rune(s[leftIndex])] != leftIndex {
				leftIndex++
			}
			if (i - leftIndex + 1) < minSuffix {
				minSuffix = i - leftIndex + 1
				x = leftIndex
				y = i
			}
		}
	}
	return x, y
}
