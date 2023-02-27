package kuy8

func FakeBin(x string) (result string) {

	for _, char := range x {
		if char < '5' {
			result += "0"
		} else {
			result += "1"
		}
	}
	return
}
