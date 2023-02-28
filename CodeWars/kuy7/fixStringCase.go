package kuy7

import (
	"strings"
	"unicode"
)

func solve(str string) string {
	upperSymbolsAmount := 0
	lowerSymbolsAmount := 0
	for _, symbol := range str {
		if unicode.IsUpper(symbol) {
			upperSymbolsAmount++
		} else {
			lowerSymbolsAmount++
		}

	}
	if upperSymbolsAmount <= lowerSymbolsAmount {
		str = strings.ToLower(str)
	} else {
		str = strings.ToUpper(str)
	}

	return str
}
