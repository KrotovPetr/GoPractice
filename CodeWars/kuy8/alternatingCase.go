package kuy8

import (
	"strings"
)

func ToAlternatingCase(str string) string {
	b := ""
	for _, symb := range str {
		if symb >= 'A' && symb <= 'Z' {
			b += strings.ToLower(string(symb))
		} else {
			b += strings.ToUpper(string(symb))
		}
	}
	return b
}
