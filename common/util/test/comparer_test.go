package util_test

import (
	"testing"

	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/util"
)

func TestGetDifferenceBetweenStrArrOK(t *testing.T) {
	// given
	diffstr := "!"
	arr1 := []string{
		"hello", "world",
	}
	arr2 := arr1[:]
	arr2 = append(arr2, diffstr)

	// when
	result := util.GetDifferenceBetweenStrArr(arr1, arr2)

	// then
	if result[0] != diffstr {
		t.Fatalf("wrong result. given arr: %v", result)
	}
}
