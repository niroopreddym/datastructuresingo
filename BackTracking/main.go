package main

import "fmt"

func main() {
	input := "abc"
	runes := strToCharArray(input)
	calculatePermuations(runes, 0, len(runes)-1)
}

func calculatePermuations(inputSlice []rune, startIndex int, endIndex int) {
	if startIndex == endIndex {
		fmt.Println(string(inputSlice))
	} else {
		for i := startIndex; i <= endIndex; i++ {
			str := swap(inputSlice, startIndex, i)
			calculatePermuations(strToCharArray(str), startIndex+1, endIndex)
		}
	}
}

func swap(inputSlice []rune, index1 int, index2 int) string {
	temp := inputSlice[index1]
	inputSlice[index1] = inputSlice[index2]
	inputSlice[index2] = temp
	return string(inputSlice)
}

func strToCharArray(str string) []rune {
	runes := []rune{}
	for _, rune1 := range str {
		runes = append(runes, rune1)
	}

	return runes
}
