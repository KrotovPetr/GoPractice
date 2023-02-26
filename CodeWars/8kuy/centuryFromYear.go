package kuy8

func Century(year int) int {
	if year%100 > 0 {
		return year/100 + 1
	} else {
		return year / 100
	}
}
