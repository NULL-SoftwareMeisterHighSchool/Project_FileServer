package util

func GetMin(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func GetMax(x, y int) int {
	if x > y {
		return x
	}
	return y
}
