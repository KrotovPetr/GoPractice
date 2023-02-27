package kuy8

func Between(a, b int) []int {
	var arr []int
	for a <= b {
		arr = append(arr, a)
		a++
	}
	return arr
}
