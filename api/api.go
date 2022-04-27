package api

import (
	"strings"
)

//compare both the version
func Compare(version1, version2 string) int {
	if version1 == version2 {
		return 0
	}

	version1Arr := strings.Split(version1, ".")
	version2Arr := strings.Split(version2, ".")

	version1Arr, version2Arr = matchSize(version1Arr, version2Arr)

	for i, num := range version1Arr {
		if num < version2Arr[i] {
			return -1
		}
		if num > version2Arr[i] {
			return 1
		}

	}
	return 0
}

// add 0 to make both the versions of same length
func matchSize(versionarr1, versionarr2 []string) ([]string, []string) {
	if len(versionarr1) == len(versionarr2) {
		return versionarr1, versionarr2
	}
	if len(versionarr1) > len(versionarr2) {
		for i := range versionarr1 {
			if i+1 > len(versionarr2) {
				versionarr2 = append(versionarr2, "0")
			}
		}
	} else {
		for i := range versionarr2 {
			if i+1 > len(versionarr1) {
				versionarr1 = append(versionarr1, "0")
			}
		}
	}
	return versionarr1, versionarr2
}
