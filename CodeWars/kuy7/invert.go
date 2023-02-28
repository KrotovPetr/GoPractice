package kuy7

func Invert(arr []int) (result []int) {
	for _, i := range arr {
		result = append(result, -i)
	}
	return
}
