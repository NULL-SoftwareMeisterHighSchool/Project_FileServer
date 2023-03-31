package util

// compares two string arrays
// and returns an array of string which newArr doesn't have
func GetDifferenceBetweenStrArr(baseArr, newArr []string) []string {
	var diff []string

	newArrMap := make(map[string]bool)
	for _, str := range newArr {
		newArrMap[str] = true
	}

	for _, str := range baseArr {
		if _, contains := newArrMap[str]; !contains {
			diff = append(diff, str)
		}
	}

	return diff
}
