package kuy8

func SmallestIntegerFinder(numbers []int) int {
	min := numbers[0]
	for index, number := range numbers {
		if index == 0 {
			continue
		} else {
			if number < min {
				min = number
			}
		}
	}
	return min // your code here
}
