package utils

type V int

func Max(a, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}
