package kuy7

import (
	"fmt"
	"strings"
)

func ReverseWords(str string) string {
	strArr := strings.Split(str, " ")
	var ansArr []string
	for i := 0; i < len(strArr); i++ {

		ansArr = append(ansArr, strArr[i])

	}

	for i := 0; i < len(ansArr); i++ {
		bufferArr := strings.Split(ansArr[i], "")
		bufferStr := ""
		for j := len(bufferArr) - 1; j >= 0; j-- {
			bufferStr += bufferArr[j]
		}
		ansArr[i] = bufferStr
	}
	fmt.Print(ansArr)
	ansStr := ""
	for i := 0; i < len(ansArr); i++ {
		if i == len(ansArr)-1 {
			ansStr += ansArr[i]
		} else {
			ansStr += ansArr[i] + " "
		}
	}
	return ansStr // reverse those words
}
