package kuy7

import "math"

func SumCubes(n int) (result int) {
	result = 0
	for i := 1; i <= n; i++ {
		result += int(math.Pow(float64(i), 3))
	}
	return
}
