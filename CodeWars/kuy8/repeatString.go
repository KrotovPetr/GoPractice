package kuy8

func RepeatStr(repetitions int, value string) string {
	str := ""
	i := 0
	for i < repetitions {
		str += value
		i++
	}
	return str
}
