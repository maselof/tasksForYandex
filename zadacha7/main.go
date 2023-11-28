package main

import (
	"log"
	"strings"
)

func main() {
	log.Println(supersvop("etot arbuz spelii"))
}

func supersvop(s string) string {
	mass := strings.Fields(s)
	leftIndex := 0
	rightIndex := len(mass) - 1
	for leftIndex < rightIndex {
		mass[leftIndex], mass[rightIndex] = mass[rightIndex], mass[leftIndex]
		leftIndex++
		rightIndex--
	}
	return strings.Join(mass, " ")
}
