package kuy7

import "strings"

func Capitalize(st string, arr []int) string {
	stCopy := ""
	leftBorder := 0
	for index, sym := range st {
		if leftBorder == len(arr) {
			continue
		}
		if index == arr[leftBorder] {
			stCopy += strings.ToUpper(string(sym))
			leftBorder++
		}

	}
	return stCopy
}
