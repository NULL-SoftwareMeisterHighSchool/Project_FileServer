package util

// compares two string arrays
// and returns an array of string which arr1 doesn't have
func GetDifferenceBetweenStrArr(arr1, arr2 []string) []string {
	var diff []string

	arr1Map := make(map[string]bool)
	for _, str := range arr1 {
		arr1Map[str] = true
	}

	for _, str := range arr2 {
		if _, contains := arr1Map[str]; !contains {
			diff = append(diff, str)
		}
	}

	return diff
}
