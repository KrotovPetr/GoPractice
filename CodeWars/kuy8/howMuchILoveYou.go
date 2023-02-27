package kuy8

func HowMuchILoveYou(i int) string {
	switch (i - 1) % 6 {
	case 0:
		return "I love you"
	case 1:
		return "a little"
	case 2:
		return "a lot"
	case 3:
		return "passionately"
	case 4:
		return "madly"
	default:
		return "not at all"
	}
}
